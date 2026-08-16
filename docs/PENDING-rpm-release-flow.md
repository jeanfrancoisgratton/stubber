# Pending: port the 2026-08-16 RPM release-flow changes into the templates

Work order for a future session. **Nothing in stubber has been changed yet.**
Two changes were rolled out across the `devops/` fleet on 2026-08-16 and now
need to reach `src/assets/rpm/` so newly scaffolded projects are born with
them. Written while the fleet rollout was fresh; the reasoning is the part
worth keeping, because both changes look arbitrary without it.

Reference implementation, authoritative over anything quoted here:
`/localrepos/devops/vmman4/__redhat/{Makefile,updateChangelog.sh}`.

---

## Change 1 — record the RPM changelog on develop, not on the built branch

### What the old flow did

`make release` ran `rpmcl → commitcl → git push → upload → tag`. The builder
checks out `origin/main` (`git reset --hard`, see
`docker_artifacts/rpmbuilder/files/sbin/buildpackage.sh`), so `commitcl`
committed the regenerated `%changelog` onto **main**, and `git push` sent it
there. Every release therefore left main one commit ahead of develop and
required a manual main→develop merge afterwards.

### Why it moved to develop rather than being merged back

Forgejo's webhook filters on `main`
(`forgejo/files/configure_webhooks.sh`, `JK_BRANCH_FILTER`), and the pipeline
filters again on `^refs/heads/main$`
(`jenkins/files/jenkins/pipelines/package-build.Jenkinsfile`). So the old
flow's push to main **re-fired the hook and triggered a redundant build of all
four distributions on every release**. The second round was mostly idempotent
(the duplicate guard suppressed a second changelog entry, the tag guard
skipped) but it still re-ran `nxtools upload` for a version already in Nexus.

A push to develop is dropped at the hook's branch filter. That asymmetry is
the whole argument: recording on develop costs one commit and nothing else,
while merging back costs a merge, two pushes, and a spurious four-distro
build.

The tradeoff accepted: main's specfile history lags one release behind, until
the next develop→main merge carries the entry forward. The built RPM is
unaffected — `rpmcl` runs `changelog` before `rpmbuild`, so the entry is in
the package either way.

### Target state

- New `CL_BRANCH ?= develop` and
  `SPECREL := $(shell git ls-files --full-name -- $(SPEC))` after the `TAG`
  assignment.
- `commitcl` rewritten: fetch `origin/$(CL_BRANCH)`, extract *that branch's*
  specfile, insert the entry, commit with `commit-tree` against the fetched
  tip, `git push origin <commit>:refs/heads/$(CL_BRANCH)`.
- `git push` removed from `release` — nothing commits to the built branch any
  more, so it is dead.
- `updateChangelog.sh` gains `--apply FILE`; notices move to stderr.

### Traps — each of these cost real time

1. **`git push HEAD:develop` does not work.** Develop is normally ahead of
   main, so it is a non-fast-forward and gets rejected. The commit must be
   built with the fetched `origin/develop` as its parent, which is why this
   uses plumbing instead of a checkout. (The workspace also has to stay on
   main: it is mid-build.)

2. **Never copy the built branch's specfile wholesale onto develop.** It
   silently reverts whatever develop has moved on to — a `_rel` bump, most
   obviously. This was verified against a live case: main built `_rel 3` while
   develop sat at `_rel 2`, and only the `%changelog` block must cross over.

3. **The duplicate guard has to follow the target file.** Once the changelog
   lives on another branch, the built branch's copy never receives the entry
   and so cannot answer "was this already recorded?". Guarding the wrong copy
   records the entry twice.

4. **`git ls-tree <rev> -- <path>` resolves the pathspec against the current
   directory.** The Makefile always runs from `__redhat/`, so the lookup found
   nothing and reported `does not exist on develop`. Needs `--full-tree`. This
   is the one bug that actually surfaced during testing.

5. **`awk` rewrites the file unchanged and reports success when there is no
   `%changelog` line.** With `--apply` the script is handed a file it did not
   write, so it checks explicitly. Same failure family as the provisioner
   notes in `docker_artifacts/CLAUDE.md`: an error made indistinguishable from
   an empty result.

6. **Log to stderr only.** `--apply` exists so a caller can capture output;
   the same rule the fleet's `plog`/`pwarn` follow.

### Race behavior, deliberately left as-is

The push's parent is the tip just fetched, so it is always a fast-forward. If
someone pushes to develop inside that window it is rejected and the build goes
red *before* `upload`. Every step is idempotent, so re-running is the fix. No
retry loop was added — this was a conscious choice, not an oversight.

---

## Change 2 — tag name must not contain `~`

`TAG := v$(VERSION)` → `TAG := v$(subst ~,-,$(VERSION))`, plus a
`git check-ref-format` guard as step 0 of `check-tags`. Both lifted from
vmman4 commit `3a6c60b`.

RPM uses `~` as its pre-release separator (`0.30.00~DEBUG` sorts *before*
`0.30.00`); git forbids `~` in a ref name because it is revision syntax.
`-` is legal in both, and is also a valid SemVer pre-release separator, so it
keeps the "sorts before the final release" meaning. Only the tag is rewritten
— the tarball keeps the raw version, since `Source0` expects
`$(NAME)-$(VERSION).tar.gz` verbatim.

**Why the guard matters as much as the substitution:** `check-tags` cannot
otherwise catch it. Its existence check is
`git rev-parse -q --verify "refs/tags/$(TAG)"`, which exits 1 on a *malformed*
ref exactly as it does on a *missing* one. So check-tags sails past and the
failure surfaces in the `tag` target — which is last in `release`, i.e. after
the RPM has been built, the changelog committed, and the artifact uploaded to
Nexus. And the missing tag then widens the next release's changelog range,
because `updateChangelog.sh` derives it from `git describe --tags`. The damage
compounds.

Latent for plain SemVer versions: `$(subst ~,-,2.8.0)` is `2.8.0`. That is the
argument for shipping it in the template — zero behavioural cost now, versus
finding it the day someone scaffolds a project that ships `1.0.0~rc1`.

---

## Notes specific to stubber

- The templates to change are `src/assets/rpm/Makefile` and
  `src/assets/rpm/updateChangelog.sh`. `src/createAssets/redhat.go` already
  writes both (`paths` slice); no new asset file is needed, so
  `src/createAssets/coverage_test.go` should not need a new mapping — verify
  rather than assume.

- **The templates are a third variant, not a copy of the fleet's.** A verbatim
  patch from the devops repos will not apply. Known divergences as of
  2026-08-16: the template `commitcl` prints `No changelog changes staged;
  nothing to commit.` where the fleet prints `NOTE: no staged changes to
  $(SPEC), ...`; the template `tag:` target has different messages and
  slightly different logic; the template uses `$(SPEC)` where some fleet
  copies use `${SPEC}`. Re-diff before editing.

- The template `updateChangelog.sh` already carries the "HEAD is already
  tagged" fallback, so only the `--apply` mechanics are missing.

- `{{ SOFTWARE NAME }}` is the placeholder in both files; the `--apply`
  rewrite must not disturb it.

- **Do not add `_dontexec` to `src/assets/skeleton/gitignore`.** It is
  correctly absent today. The build veto is evaluated inside the builder's
  checkout of `origin/main`, so an ignored flag never reaches it — a
  gitignored `_dontexec` looks like a veto and does nothing. One fleet repo
  had this and it silently dropped the flags from a commit.

- Consider whether scaffolding should emit `dontexec.sh` (the helper that
  creates/removes the marker). Only vmman4 has one today; the rest of the
  fleet does not. Open question, not a decision.

---

## How to verify without touching a live repo

The bug in trap 4 was found this way, and it is cheap:

1. `git clone --bare <repo> /tmp/origin.git`, then clone that into a workspace.
2. In the workspace, reproduce what the builder does:
   `git checkout -B main --track origin/main && git reset --hard origin/main`.
3. Copy the candidate templates in, commit, push to both `main` and `develop`.
4. `make commitcl` with a stubbed `rpmspec` on `PATH` if the host lacks it — a
   dozen-line shim that greps `%define _name/_version/_rel` out of the spec is
   enough.
5. Assert: the commit lands on develop with the fetched tip as its parent, the
   diff is confined to `%changelog`, `main` is unchanged, and a second run
   reports "already carries ... nothing to commit" and exits 0.
6. Then make develop's `_rel` differ from main's and confirm develop's value
   survives.

---

## Fleet status as of 2026-08-16 (context, not work)

Live on main: `vmman4`, `dtools2`, `dvol`, `nxtools`, `vaultreader`, `vclt`.
Edited but left uncommitted because their workspaces were dirty:
`certificatemanager` (on `feature/v2`), `monctl`. `pkitools` was deleted; its
Forgejo repo was nuked upstream.

`dvol` carries a deliberate `__redhat/_dontexec`, so its RPM build is vetoed
and the new flow is inert there — that is intentional, not a bug to fix.

Known tag/version drift, never investigated: `monctl` spec 1.6.2 vs newest tag
`v1.6.1`; `vclt` has `v2.4.0`/`v2.4.1`/`v2.4.3` but no `v2.4.2`; `vaultreader`
spec 2.3.1 with no `v*` tags at all.
