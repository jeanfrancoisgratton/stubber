#### CREATE
- [ ] RPM : Quotes in specfile ?<br> 
- [ ] RPM : Group cannot be empty in specfile<br>
- [ ] APK : pkgdesc needs to be enclosed in Quotes<br>
- [ ] APK : Makefile is fetching GOLANG x86_64 instead of amd64<br>
- [x] All __debian/ files need chmod 755

#### UPDATE

**When a templated field is ignored at command line, this field becomes blank in the target asset**

- [ ] __alpine/APKBUILD: arch: reverted to amd64 instead of X86_64
- [ ] __debian/ multiple issues: typo in control, missing fields, file needlessly updated (2.*, etc)
- [ ] dtools.spec is being borked
- [x] rpmbuild-deps.sh is being ignored when upgrading GO version. Mode reverted to 0644
  - Fixed in `__redhat/` and in `src/assets/rpm/`: read `../go.version`, not `go.version` — the script runs from `__redhat/`, so `VER` came out empty and wget fetched a nonexistent tarball
  - Mode: no longer reproducible. `createAssets/redhat.go` chmods it 0755 after rendering, and a scaffold check confirms 0755 on the generated file
- [x] RPM : port the 2026-08-16 release-flow changes into `src/assets/rpm/` (changelog commit goes to develop, `~` in tag names) — see [PENDING-rpm-release-flow.md](PENDING-rpm-release-flow.md), now fully consumed and safe to delete
- [x] `__archlinux/PKGBUILD` renders `url="{{ URL }}"` literally: `createAssets/archlinux.go` did not pass `{{ URL }}` in its placeholders map. Fixed
- [x] ArchLinux silently dropped `-e` / `--depends`: `PKGBUILD` hardcoded `depends=()`. Now rendered via `{{ DEPENDENCIES }}`, rewritten from the Debian-shaped flag into a quoted pacman array
- [x] License metadata contradicted the shipped `LICENSE` (GPLv3): `PKGBUILD` said `GPL-2.0-only`, `APKBUILD` said `GPL2`, the specfile said `GPL2.0`. All now declare `GPL-3.0-or-later`, in the templates **and** in stubber's own `__alpine` / `__archlinux` / `__redhat`. Both `LICENSE` files are the complete GPLv3 text and were already correct — the "or later" grant is in their boilerplate
- [x] README swept against a real scaffold: `__alpine`/`__archlinux`/`__debian` Makefiles were missing from the layout tree; `__debian` claimed `1.install-build-deps.sh` / `2.build_binary.sh` / `3.restore_repo.sh` (the numbered scripts are gone, folded into the Makefile); the RedHat "typical workflow" block still ended in `git push`; `[ISSUES.md](ISSUES.md)` and `[GPL-3.0](LICENSE)` were both broken links (the files live in `docs/`); the Contents anchor `#archlinux--a-1` no longer resolved
- [x] `src/assets/skeleton/ROADMAP.md` and `TODO.md` were embedded but never rendered — not in `skeleton.go`'s `paths`. Now emitted. Added `TestEveryEmbeddedSkeletonAssetIsRendered`, the skeleton counterpart of the packaging drift-guard, which promptly found a third: `src/_importCheck.sh`. That one was a leftover from an earlier epoch, so it was deleted rather than emitted
- [x] No `debian/copyright` in the Debian stub: `control` carries no license field, so a generated `.deb` declared no license at all. Added a DEP-5 `copyright` to the template **and** to stubber's own `__debian/`; the Makefile installs it to `/usr/share/doc/<package>/copyright` at 0644 (policy location — not `DEBIAN/`). Needed two new placeholders, `{{ URL }}` and `{{ COPYRIGHT YEAR }}`
- [x] `LICENSE` moved from `docs/` to the repo root, matching where stubber already puts it in scaffolded projects, and fixing the README link
- [x] Embedded asset directories renamed to match their `__*` output dirs and stub functions: `apk`→`alpine`, `deb`→`debian`, `arch`→`archlinux`, `rpm`→`redhat`
- [x] `PACKAGE_RPM.md` dropped from the RedHat stub; the RPM workflow is documented in README instead

<br><br><br>