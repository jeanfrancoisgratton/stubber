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
- [ ] rpmbuild-deps.sh is being ignored when upgrading GO version. Mode reverted to 0644
  - [x] GO version half fixed in `__redhat/` (read `../go.version`, not `go.version` — the script runs from `__redhat/`)
  - [ ] still to do: same fix in `src/assets/rpm/rpmbuild-deps.sh`, and the 0644 mode on the generated asset
- [ ] RPM : port the 2026-08-16 release-flow changes into `src/assets/rpm/` (changelog commit goes to develop, `~` in tag names) — see [PENDING-rpm-release-flow.md](PENDING-rpm-release-flow.md)

<br><br><br>