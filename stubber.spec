%ifarch aarch64
%global _arch aarch64
%global BuildArchitectures aarch64
%endif

%ifarch x86_64
%global _arch x86_64
%global BuildArchitectures x86_64
%endif

%define debug_package   %{nil}
%define _build_id_links none
%define _name   stubber
%define _prefix /opt
%define _version 1.80.02
%define _rel 0
#%define _arch x86_64
%define _binaryname stubber

Name:       stubber
Version:    %{_version}
Release:    %{_rel}
Summary:    stubber

Group:      Utils
License:    GPL2.0
URL:        https://github.com/jeanfrancoisgratton/stubber

Source0:    %{name}-%{_version}.tar.gz
#BuildArchitectures: x86_64
BuildRequires: gcc
#Requires: sudo
#Obsoletes: vmman1 > 1.140

%description
Creates a GO software skeleton

%prep
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src/templates
rm -f assets.go
sudo GOBIN=/opt/go/bin /opt/go/bin/go install -a github.com/go-bindata/go-bindata/...@latest
sudo /opt/go/bin/go generate
cd ..
CGO_ENABLED=0 /opt/go/bin/go build -o %{_sourcedir}/%{_binaryname} .
strip %{_sourcedir}/%{_binaryname}


%clean
rm -rf $RPM_BUILD_ROOT

%pre
if getent group devops > /dev/null; then
  exit 0
else
  if getent group 2500; then
    groupadd devops
  else
    groupadd -g 2500 devops
  fi
fi
exit 0

%install
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post
cd /opt/bin
sudo chgrp -R devops .
sudo chmod 775 /opt/bin/%{_binaryname}

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{_binaryname}


%changelog
* Wed Sep 24 2025 Binary package builder <builder@famillegratton.net> 1.80.02-0
- Fix attempt #3 on checkImports (jean-francois@famillegratton.net)

* Wed Sep 24 2025 Binary package builder <builder@famillegratton.net> 1.80.01-0
- Package version bump (builder@famillegratton.net)
- renamed the import checker script (jean-francois@famillegratton.net)
- gitignore fix (jean-francois@famillegratton.net)


* Wed Jul 09 2025 Binary package builder <builder@famillegratton.net> 1.78.01-1
- GO version bump (jean-francois@famillegratton.net)

* Tue Jul 01 2025 Binary package builder <builder@famillegratton.net> 1.78.00-1
- disabled CGO wherever I had forgotten (jean-francois@famillegratton.net)

* Tue Jul 01 2025 Binary package builder <builder@famillegratton.net> 1.78.00-0
- disable CGO in assets generation (jean-francois@famillegratton.net)
- Disabled CGO in package building (jean-francois@famillegratton.net)

* Sat Jun 14 2025 APK Builder <builder@famillegratton.net> 1.77.00-0
- fixes to default version value (jean-francois@famillegratton.net)
- added a new asset (template) variable to deal with version and changelog
  (jean-francois@famillegratton.net)

* Sat Jun 14 2025 APK Builder <builder@famillegratton.net> 1.76.00-0
- Software version bump (jean-francois@famillegratton.net)
- Fixed build script, GO version bump (jean-francois@famillegratton.net)
- GO version bump (builder@famillegratton.net)
- GO version bump (builder@famillegratton.net)
- updated builddeps (builder@famillegratton.net)
- Fixed rpm builddeps script (jean-francois@famillegratton.net)
- Stubbed new branch for git management (jean-francois@famillegratton.net)

* Tue Apr 08 2025 APK Builder <builder@famillegratton.net> 1.75.01-0
- Version bump (builder@famillegratton.net)
- updated rpm deps script (builder@famillegratton.net)

* Sat Mar 29 2025 APK Builder <builder@famillegratton.net> 1.75.00-0
- updated build.sh script (jean-francois@famillegratton.net)

* Fri Mar 14 2025 APK Builder <builder@famillegratton.net> 1.74.00-0
- rpmbuilddeps fixes (builder@famillegratton.net)
- updated builddeps (jean-francois@famillegratton.net)
- APK url var fix, go version bump (jean-francois@famillegratton.net)

* Fri Nov 15 2024 APK Builder <builder@famillegratton.net> 1.73.00-0
- removed unneeded files, GO version bump (jean-francois@famillegratton.net)
- Updated rpm builddeps (builder@famillegratton.net)

* Sat Oct 19 2024 RPM Builder <builder@famillegratton.net> 1.72.01-0
- Attempt at modifying go-bindata issue (builder@famillegratton.net)

* Sat Oct 19 2024 RPM Builder <builder@famillegratton.net> 1.72.00-0
- re-initialized tito

* Sun Aug 11 2024 RPM Builder <builder@famillegratton.net> 1.70.00-0
- new package built with tito

* Sun Aug 11 2024 RPM Builder <builder@famillegratton.net> 1.70.00-0
- Go version bump, reverted tito tag naming scheme change

* Mon Aug 05 2024 RPM Builder <builder@famillegratton.net> 1.65.01-0
- Fixed wrong flag for GHA enabling (jean-francois@famillegratton.net)
- mode change (jean-francois@famillegratton.net)
- Various updates in assets (jean-francois@famillegratton.net)

* Sat Aug 03 2024 RPM Builder <builder@famillegratton.net> 1.65.00-0
- Added GHA to template (jean-francois@famillegratton.net)
- Asset update (jean-francois@famillegratton.net)
- updated GO version in pre-flight script (builder@famillegratton.net)

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- 

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- retagging (jean-francois@famillegratton.net)

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- 

* Sat May 25 2024 RPM Builder <builder@famillegratton.net> 1.61.01-0
- Added missing asset file (jean-francois@famillegratton.net)

* Sat May 25 2024 RPM Builder <builder@famillegratton.net> 1.61.00-0
- Version bump and deps maintenance scripts update (jean-
  francois@famillegratton.net)
- Rewrote build.sh asset (jean-francois@famillegratton.net)
- GO version bump, rewrite of build.sh (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- Fixed perms on deps script (builder@famillegratton.net)
- Fixed issue with go mod tidy (builder@famillegratton.net)
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Feb 16 2024 RPM Builder <builder@famillegratton.net>
- Assets update (jean-francois@famillegratton.net)
- Packaging fixes (jean-francois@famillegratton.net)

* Thu Feb 15 2024 RPM Builder <builder@famillegratton.net>
- Forgot bumping release in deb packaging (builder@famillegratton.net)
- Ensuring that all binary packages have the same version/release number (jean-
  francois@famillegratton.net)

* Thu Feb 15 2024 RPM Builder <builder@famillegratton.net>
- Fixes in RPM and APK packaging scripts (jean-francois@famillegratton.net)
- Removed arch variable as we no longer support arm64 (jean-
  francois@famillegratton.net)

* Wed Feb 14 2024 RPM Builder <builder@famillegratton.net>
- Go version bump, arm64 arch removal, more binary package scripts (jean-
  francois@famillegratton.net)
- Fix to upgradeBuildDeps (jean-francois@famillegratton.net)
- Added FIXME issues, renamed upgrade_pkgs.sh (jean-
  francois@famillegratton.net)
- Version bump : forgotten files (jean-francois@famillegratton.net)
- Go and software version bump (jean-francois@famillegratton.net)

* Tue Jan 09 2024 RPM Builder <builder@famillegratton.net> 1.53.00-0
- Assets fixes (jean-francois@famillegratton.net)
- Minor version fix, will not re-release for that (jean-
  francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.02-0
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net>
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.02-0
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.01-1
- Release number bump (jean-francois@famillegratton.net)
- Fixed default GO version to 1.21.5 (jean-francois@famillegratton.net)
- Update NEED_FIXES.txt (jean-francois@famillegratton.net)
- Update NEED_FIXES.txt (jean-francois@famillegratton.net)
- Fixed assets path (jean-francois@famillegratton.net)
- Asset fixes (jean-francois@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net> 1.52.00-0
- GO and package versions update (jean-francois@famillegratton.net)
- Automatic commit of package [stubber] release [1.52.00-0].
  (builder@famillegratton.net)
- Syntax-typo fixes (jean-francois@famillegratton.net)
- Finalized synching (jean-francois@famillegratton.net)
- sync zenika -> (jean-francois@famillegratton.net)
- Sync zenika-> (jean-francois@famillegratton.net)
- Fixed version number on Debian package (jean-francois@famillegratton.net)
- Removed unused line (jean-francois@famillegratton.net)
- Sync Zenika-> (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Permission fix on build script (builder@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net>
- GO and package versions update (jean-francois@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net> 1.52.00-0
- Syntax-typo fixes (jean-francois@famillegratton.net)
- Finalized synching (jean-francois@famillegratton.net)
- sync zenika -> (jean-francois@famillegratton.net)
- Sync zenika-> (jean-francois@famillegratton.net)
- Fixed version number on Debian package (jean-francois@famillegratton.net)
- Removed unused line (jean-francois@famillegratton.net)
- Sync Zenika-> (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Permission fix on build script (builder@famillegratton.net)

* Sat Aug 19 2023 RPM Builder <builder@famillegratton.net> 1.505-2
- Added extra cleanup task to DEB package (builder@famillegratton.net)
- Typo fix, version bump in RPM stub (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Fixed issue of a unresolved function name in cmd/root.go, version bump (jean-
  francois@famillegratton.net)
- Bug fix: undefined command in cmd/root.go (jean-francois@famillegratton.net)

* Thu Aug 17 2023 RPM Builder <builder@famillegratton.net> 1.500-0
- Completed debugging (jean-francois@famillegratton.net)
- Screwup (jean-francois@famillegratton.net)
- Sync between branches (jean-francois@famillegratton.net)
- Fixed various filepaths (jean-francois@famillegratton.net)
- Software version bump (jean-francois@famillegratton.net)
- Refactoring before creating the updateAssets package (jean-
  francois@famillegratton.net)
- Removed helpers.Changelog() from assets (jean-francois@famillegratton.net)
- minor DEB stub fix (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.206-0
- Typo fix (jean-francois@famillegratton.net)
- Reverted version bump (jean-francois@famillegratton.net)
- Added gitignore in assets (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.205-0
- Fixed missing placeholder in rpm stub (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-2
- Forgotten version bump (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-1
- Minor fix: changelog update (cosmetic issue) (jean-
  francois@famillegratton.net)


* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-0
- Fixed flags duplication (jean-francois@famillegratton.net)
- Doc update (builder@famillegratton.net)
- Added GO GEN commands (jean-francois@famillegratton.net)
- Updated fixme (jean-francois@famillegratton.net)
- Ready to test on Debian (jean-francois@famillegratton.net)

* Sat Aug 12 2023 builder <builder@famillegratton.net> 1.100-0
- Yet another permission issue (my_email@internet.net)
- Fixed missing placeholder and various ARCH issues (jean-
  francois@famillegratton.net)

* Sat Aug 12 2023 builder <builder@famillegratton.net> 1.010-0
- Debian packaging fixes (my_email@internet.net)
- Fixed missing flags, removed some assets (jean-francois@famillegratton.net)
- Gave up on MD formating (jean-francois@famillegratton.net)
- Minor doc update (jean-francois@famillegratton.net)
- rpm packaging perms fix (builder@famillegratton.net)

* Fri Aug 11 2023 builder <builder@famillegratton.net> 1.000-0
- new package built with tito


* Sun Aug 11 2024 RPM Builder <builder@famillegratton.net> 1.70.00-0
- 

* Mon Aug 05 2024 RPM Builder <builder@famillegratton.net> 1.65.01-0
- Fixed wrong flag for GHA enabling (jean-francois@famillegratton.net)
- mode change (jean-francois@famillegratton.net)
- Various updates in assets (jean-francois@famillegratton.net)

* Sat Aug 03 2024 RPM Builder <builder@famillegratton.net> 1.65.00-0
- Added GHA to template (jean-francois@famillegratton.net)
- Asset update (jean-francois@famillegratton.net)
- updated GO version in pre-flight script (builder@famillegratton.net)

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- 

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- retagging (jean-francois@famillegratton.net)

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- 

* Sat May 25 2024 RPM Builder <builder@famillegratton.net> 1.61.01-0
- Added missing asset file (jean-francois@famillegratton.net)

* Sat May 25 2024 RPM Builder <builder@famillegratton.net> 1.61.00-0
- Version bump and deps maintenance scripts update (jean-
  francois@famillegratton.net)
- Rewrote build.sh asset (jean-francois@famillegratton.net)
- GO version bump, rewrite of build.sh (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- Fixed perms on deps script (builder@famillegratton.net)
- Fixed issue with go mod tidy (builder@famillegratton.net)
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Feb 16 2024 RPM Builder <builder@famillegratton.net>
- Assets update (jean-francois@famillegratton.net)
- Packaging fixes (jean-francois@famillegratton.net)

* Thu Feb 15 2024 RPM Builder <builder@famillegratton.net>
- Forgot bumping release in deb packaging (builder@famillegratton.net)
- Ensuring that all binary packages have the same version/release number (jean-
  francois@famillegratton.net)

* Thu Feb 15 2024 RPM Builder <builder@famillegratton.net>
- Fixes in RPM and APK packaging scripts (jean-francois@famillegratton.net)
- Removed arch variable as we no longer support arm64 (jean-
  francois@famillegratton.net)

* Wed Feb 14 2024 RPM Builder <builder@famillegratton.net>
- Go version bump, arm64 arch removal, more binary package scripts (jean-
  francois@famillegratton.net)
- Fix to upgradeBuildDeps (jean-francois@famillegratton.net)
- Added FIXME issues, renamed upgrade_pkgs.sh (jean-
  francois@famillegratton.net)
- Version bump : forgotten files (jean-francois@famillegratton.net)
- Go and software version bump (jean-francois@famillegratton.net)

* Tue Jan 09 2024 RPM Builder <builder@famillegratton.net> 1.53.00-0
- Assets fixes (jean-francois@famillegratton.net)
- Minor version fix, will not re-release for that (jean-
  francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.02-0
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net>
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.02-0
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.01-1
- Release number bump (jean-francois@famillegratton.net)
- Fixed default GO version to 1.21.5 (jean-francois@famillegratton.net)
- Update NEED_FIXES.txt (jean-francois@famillegratton.net)
- Update NEED_FIXES.txt (jean-francois@famillegratton.net)
- Fixed assets path (jean-francois@famillegratton.net)
- Asset fixes (jean-francois@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net> 1.52.00-0
- GO and package versions update (jean-francois@famillegratton.net)
- Automatic commit of package [stubber] release [1.52.00-0].
  (builder@famillegratton.net)
- Syntax-typo fixes (jean-francois@famillegratton.net)
- Finalized synching (jean-francois@famillegratton.net)
- sync zenika -> (jean-francois@famillegratton.net)
- Sync zenika-> (jean-francois@famillegratton.net)
- Fixed version number on Debian package (jean-francois@famillegratton.net)
- Removed unused line (jean-francois@famillegratton.net)
- Sync Zenika-> (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Permission fix on build script (builder@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net>
- GO and package versions update (jean-francois@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net> 1.52.00-0
- Syntax-typo fixes (jean-francois@famillegratton.net)
- Finalized synching (jean-francois@famillegratton.net)
- sync zenika -> (jean-francois@famillegratton.net)
- Sync zenika-> (jean-francois@famillegratton.net)
- Fixed version number on Debian package (jean-francois@famillegratton.net)
- Removed unused line (jean-francois@famillegratton.net)
- Sync Zenika-> (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Permission fix on build script (builder@famillegratton.net)

* Sat Aug 19 2023 RPM Builder <builder@famillegratton.net> 1.505-2
- Added extra cleanup task to DEB package (builder@famillegratton.net)
- Typo fix, version bump in RPM stub (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Fixed issue of a unresolved function name in cmd/root.go, version bump (jean-
  francois@famillegratton.net)
- Bug fix: undefined command in cmd/root.go (jean-francois@famillegratton.net)

* Thu Aug 17 2023 RPM Builder <builder@famillegratton.net> 1.500-0
- Completed debugging (jean-francois@famillegratton.net)
- Screwup (jean-francois@famillegratton.net)
- Sync between branches (jean-francois@famillegratton.net)
- Fixed various filepaths (jean-francois@famillegratton.net)
- Software version bump (jean-francois@famillegratton.net)
- Refactoring before creating the updateAssets package (jean-
  francois@famillegratton.net)
- Removed helpers.Changelog() from assets (jean-francois@famillegratton.net)
- minor DEB stub fix (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.206-0
- Typo fix (jean-francois@famillegratton.net)
- Reverted version bump (jean-francois@famillegratton.net)
- Added gitignore in assets (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.205-0
- Fixed missing placeholder in rpm stub (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-2
- Forgotten version bump (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-1
- Minor fix: changelog update (cosmetic issue) (jean-
  francois@famillegratton.net)


* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-0
- Fixed flags duplication (jean-francois@famillegratton.net)
- Doc update (builder@famillegratton.net)
- Added GO GEN commands (jean-francois@famillegratton.net)
- Updated fixme (jean-francois@famillegratton.net)
- Ready to test on Debian (jean-francois@famillegratton.net)

* Sat Aug 12 2023 builder <builder@famillegratton.net> 1.100-0
- Yet another permission issue (my_email@internet.net)
- Fixed missing placeholder and various ARCH issues (jean-
  francois@famillegratton.net)

* Sat Aug 12 2023 builder <builder@famillegratton.net> 1.010-0
- Debian packaging fixes (my_email@internet.net)
- Fixed missing flags, removed some assets (jean-francois@famillegratton.net)
- Gave up on MD formating (jean-francois@famillegratton.net)
- Minor doc update (jean-francois@famillegratton.net)
- rpm packaging perms fix (builder@famillegratton.net)

* Fri Aug 11 2023 builder <builder@famillegratton.net> 1.000-0
- new package built with tito



* Sun Aug 11 2024 RPM Builder <builder@famillegratton.net> 1.70.00-0
- Changed tag naming scheme

* Mon Aug 05 2024 RPM Builder <builder@famillegratton.net> 1.65.01-0
- Fixed wrong flag for GHA enabling (jean-francois@famillegratton.net)
- mode change (jean-francois@famillegratton.net)
- Various updates in assets (jean-francois@famillegratton.net)

* Sat Aug 03 2024 RPM Builder <builder@famillegratton.net> 1.65.00-0
- Added GHA to template (jean-francois@famillegratton.net)
- Asset update (jean-francois@famillegratton.net)
- updated GO version in pre-flight script (builder@famillegratton.net)

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- 

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- retagging (jean-francois@famillegratton.net)

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- 

* Sat May 25 2024 RPM Builder <builder@famillegratton.net> 1.61.01-0
- Added missing asset file (jean-francois@famillegratton.net)

* Sat May 25 2024 RPM Builder <builder@famillegratton.net> 1.61.00-0
- Version bump and deps maintenance scripts update (jean-
  francois@famillegratton.net)
- Rewrote build.sh asset (jean-francois@famillegratton.net)
- GO version bump, rewrite of build.sh (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- Fixed perms on deps script (builder@famillegratton.net)
- Fixed issue with go mod tidy (builder@famillegratton.net)
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Feb 16 2024 RPM Builder <builder@famillegratton.net>
- Assets update (jean-francois@famillegratton.net)
- Packaging fixes (jean-francois@famillegratton.net)

* Thu Feb 15 2024 RPM Builder <builder@famillegratton.net>
- Forgot bumping release in deb packaging (builder@famillegratton.net)
- Ensuring that all binary packages have the same version/release number (jean-
  francois@famillegratton.net)

* Thu Feb 15 2024 RPM Builder <builder@famillegratton.net>
- Fixes in RPM and APK packaging scripts (jean-francois@famillegratton.net)
- Removed arch variable as we no longer support arm64 (jean-
  francois@famillegratton.net)

* Wed Feb 14 2024 RPM Builder <builder@famillegratton.net>
- Go version bump, arm64 arch removal, more binary package scripts (jean-
  francois@famillegratton.net)
- Fix to upgradeBuildDeps (jean-francois@famillegratton.net)
- Added FIXME issues, renamed upgrade_pkgs.sh (jean-
  francois@famillegratton.net)
- Version bump : forgotten files (jean-francois@famillegratton.net)
- Go and software version bump (jean-francois@famillegratton.net)

* Tue Jan 09 2024 RPM Builder <builder@famillegratton.net> 1.53.00-0
- Assets fixes (jean-francois@famillegratton.net)
- Minor version fix, will not re-release for that (jean-
  francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.02-0
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net>
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.02-0
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.01-1
- Release number bump (jean-francois@famillegratton.net)
- Fixed default GO version to 1.21.5 (jean-francois@famillegratton.net)
- Update NEED_FIXES.txt (jean-francois@famillegratton.net)
- Update NEED_FIXES.txt (jean-francois@famillegratton.net)
- Fixed assets path (jean-francois@famillegratton.net)
- Asset fixes (jean-francois@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net> 1.52.00-0
- GO and package versions update (jean-francois@famillegratton.net)
- Automatic commit of package [stubber] release [1.52.00-0].
  (builder@famillegratton.net)
- Syntax-typo fixes (jean-francois@famillegratton.net)
- Finalized synching (jean-francois@famillegratton.net)
- sync zenika -> (jean-francois@famillegratton.net)
- Sync zenika-> (jean-francois@famillegratton.net)
- Fixed version number on Debian package (jean-francois@famillegratton.net)
- Removed unused line (jean-francois@famillegratton.net)
- Sync Zenika-> (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Permission fix on build script (builder@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net>
- GO and package versions update (jean-francois@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net> 1.52.00-0
- Syntax-typo fixes (jean-francois@famillegratton.net)
- Finalized synching (jean-francois@famillegratton.net)
- sync zenika -> (jean-francois@famillegratton.net)
- Sync zenika-> (jean-francois@famillegratton.net)
- Fixed version number on Debian package (jean-francois@famillegratton.net)
- Removed unused line (jean-francois@famillegratton.net)
- Sync Zenika-> (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Permission fix on build script (builder@famillegratton.net)

* Sat Aug 19 2023 RPM Builder <builder@famillegratton.net> 1.505-2
- Added extra cleanup task to DEB package (builder@famillegratton.net)
- Typo fix, version bump in RPM stub (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Fixed issue of a unresolved function name in cmd/root.go, version bump (jean-
  francois@famillegratton.net)
- Bug fix: undefined command in cmd/root.go (jean-francois@famillegratton.net)

* Thu Aug 17 2023 RPM Builder <builder@famillegratton.net> 1.500-0
- Completed debugging (jean-francois@famillegratton.net)
- Screwup (jean-francois@famillegratton.net)
- Sync between branches (jean-francois@famillegratton.net)
- Fixed various filepaths (jean-francois@famillegratton.net)
- Software version bump (jean-francois@famillegratton.net)
- Refactoring before creating the updateAssets package (jean-
  francois@famillegratton.net)
- Removed helpers.Changelog() from assets (jean-francois@famillegratton.net)
- minor DEB stub fix (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.206-0
- Typo fix (jean-francois@famillegratton.net)
- Reverted version bump (jean-francois@famillegratton.net)
- Added gitignore in assets (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.205-0
- Fixed missing placeholder in rpm stub (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-2
- Forgotten version bump (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-1
- Minor fix: changelog update (cosmetic issue) (jean-
  francois@famillegratton.net)


* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-0
- Fixed flags duplication (jean-francois@famillegratton.net)
- Doc update (builder@famillegratton.net)
- Added GO GEN commands (jean-francois@famillegratton.net)
- Updated fixme (jean-francois@famillegratton.net)
- Ready to test on Debian (jean-francois@famillegratton.net)

* Sat Aug 12 2023 builder <builder@famillegratton.net> 1.100-0
- Yet another permission issue (my_email@internet.net)
- Fixed missing placeholder and various ARCH issues (jean-
  francois@famillegratton.net)

* Sat Aug 12 2023 builder <builder@famillegratton.net> 1.010-0
- Debian packaging fixes (my_email@internet.net)
- Fixed missing flags, removed some assets (jean-francois@famillegratton.net)
- Gave up on MD formating (jean-francois@famillegratton.net)
- Minor doc update (jean-francois@famillegratton.net)
- rpm packaging perms fix (builder@famillegratton.net)

* Fri Aug 11 2023 builder <builder@famillegratton.net> 1.000-0
- new package built with tito


* Sun Aug 11 2024 RPM Builder <builder@famillegratton.net> 1.70.00-0
- 

* Mon Aug 05 2024 RPM Builder <builder@famillegratton.net> 1.65.01-0
- Fixed wrong flag for GHA enabling (jean-francois@famillegratton.net)
- mode change (jean-francois@famillegratton.net)
- Various updates in assets (jean-francois@famillegratton.net)

* Sat Aug 03 2024 RPM Builder <builder@famillegratton.net> 1.65.00-0
- Added GHA to template (jean-francois@famillegratton.net)
- Asset update (jean-francois@famillegratton.net)
- updated GO version in pre-flight script (builder@famillegratton.net)

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- 

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- retagging (jean-francois@famillegratton.net)

* Sun Jul 28 2024 RPM Builder <builder@famillegratton.net> 1.62.00-0
- 

* Sat May 25 2024 RPM Builder <builder@famillegratton.net> 1.61.01-0
- Added missing asset file (jean-francois@famillegratton.net)

* Sat May 25 2024 RPM Builder <builder@famillegratton.net> 1.61.00-0
- Version bump and deps maintenance scripts update (jean-
  francois@famillegratton.net)
- Rewrote build.sh asset (jean-francois@famillegratton.net)
- GO version bump, rewrite of build.sh (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- Fixed perms on deps script (builder@famillegratton.net)
- Fixed issue with go mod tidy (builder@famillegratton.net)
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Mar 15 2024 RPM Builder <builder@famillegratton.net>
- APKBUILD now respects the -u flag (jean-francois@famillegratton.net)

* Fri Feb 16 2024 RPM Builder <builder@famillegratton.net>
- Assets update (jean-francois@famillegratton.net)
- Packaging fixes (jean-francois@famillegratton.net)

* Thu Feb 15 2024 RPM Builder <builder@famillegratton.net>
- Forgot bumping release in deb packaging (builder@famillegratton.net)
- Ensuring that all binary packages have the same version/release number (jean-
  francois@famillegratton.net)

* Thu Feb 15 2024 RPM Builder <builder@famillegratton.net>
- Fixes in RPM and APK packaging scripts (jean-francois@famillegratton.net)
- Removed arch variable as we no longer support arm64 (jean-
  francois@famillegratton.net)

* Wed Feb 14 2024 RPM Builder <builder@famillegratton.net>
- Go version bump, arm64 arch removal, more binary package scripts (jean-
  francois@famillegratton.net)
- Fix to upgradeBuildDeps (jean-francois@famillegratton.net)
- Added FIXME issues, renamed upgrade_pkgs.sh (jean-
  francois@famillegratton.net)
- Version bump : forgotten files (jean-francois@famillegratton.net)
- Go and software version bump (jean-francois@famillegratton.net)

* Tue Jan 09 2024 RPM Builder <builder@famillegratton.net> 1.53.00-0
- Assets fixes (jean-francois@famillegratton.net)
- Minor version fix, will not re-release for that (jean-
  francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.02-0
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net>
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.02-0
- Misc asset fixes (jean-francois@famillegratton.net)

* Sun Dec 31 2023 RPM Builder <builder@famillegratton.net> 1.52.01-1
- Release number bump (jean-francois@famillegratton.net)
- Fixed default GO version to 1.21.5 (jean-francois@famillegratton.net)
- Update NEED_FIXES.txt (jean-francois@famillegratton.net)
- Update NEED_FIXES.txt (jean-francois@famillegratton.net)
- Fixed assets path (jean-francois@famillegratton.net)
- Asset fixes (jean-francois@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net> 1.52.00-0
- GO and package versions update (jean-francois@famillegratton.net)
- Automatic commit of package [stubber] release [1.52.00-0].
  (builder@famillegratton.net)
- Syntax-typo fixes (jean-francois@famillegratton.net)
- Finalized synching (jean-francois@famillegratton.net)
- sync zenika -> (jean-francois@famillegratton.net)
- Sync zenika-> (jean-francois@famillegratton.net)
- Fixed version number on Debian package (jean-francois@famillegratton.net)
- Removed unused line (jean-francois@famillegratton.net)
- Sync Zenika-> (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Permission fix on build script (builder@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net>
- GO and package versions update (jean-francois@famillegratton.net)

* Fri Dec 29 2023 RPM Builder <builder@famillegratton.net> 1.52.00-0
- Syntax-typo fixes (jean-francois@famillegratton.net)
- Finalized synching (jean-francois@famillegratton.net)
- sync zenika -> (jean-francois@famillegratton.net)
- Sync zenika-> (jean-francois@famillegratton.net)
- Fixed version number on Debian package (jean-francois@famillegratton.net)
- Removed unused line (jean-francois@famillegratton.net)
- Sync Zenika-> (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Permission fix on build script (builder@famillegratton.net)

* Sat Aug 19 2023 RPM Builder <builder@famillegratton.net> 1.505-2
- Added extra cleanup task to DEB package (builder@famillegratton.net)
- Typo fix, version bump in RPM stub (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Fixed issue of a unresolved function name in cmd/root.go, version bump (jean-
  francois@famillegratton.net)
- Bug fix: undefined command in cmd/root.go (jean-francois@famillegratton.net)

* Thu Aug 17 2023 RPM Builder <builder@famillegratton.net> 1.500-0
- Completed debugging (jean-francois@famillegratton.net)
- Screwup (jean-francois@famillegratton.net)
- Sync between branches (jean-francois@famillegratton.net)
- Fixed various filepaths (jean-francois@famillegratton.net)
- Software version bump (jean-francois@famillegratton.net)
- Refactoring before creating the updateAssets package (jean-
  francois@famillegratton.net)
- Removed helpers.Changelog() from assets (jean-francois@famillegratton.net)
- minor DEB stub fix (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.206-0
- Typo fix (jean-francois@famillegratton.net)
- Reverted version bump (jean-francois@famillegratton.net)
- Added gitignore in assets (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.205-0
- Fixed missing placeholder in rpm stub (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-2
- Forgotten version bump (jean-francois@famillegratton.net)

* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-1
- Minor fix: changelog update (cosmetic issue) (jean-
  francois@famillegratton.net)


* Sun Aug 13 2023 RPM Builder <builder@famillegratton.net> 1.201-0
- Fixed flags duplication (jean-francois@famillegratton.net)
- Doc update (builder@famillegratton.net)
- Added GO GEN commands (jean-francois@famillegratton.net)
- Updated fixme (jean-francois@famillegratton.net)
- Ready to test on Debian (jean-francois@famillegratton.net)

* Sat Aug 12 2023 builder <builder@famillegratton.net> 1.100-0
- Yet another permission issue (my_email@internet.net)
- Fixed missing placeholder and various ARCH issues (jean-
  francois@famillegratton.net)

* Sat Aug 12 2023 builder <builder@famillegratton.net> 1.010-0
- Debian packaging fixes (my_email@internet.net)
- Fixed missing flags, removed some assets (jean-francois@famillegratton.net)
- Gave up on MD formating (jean-francois@famillegratton.net)
- Minor doc update (jean-francois@famillegratton.net)
- rpm packaging perms fix (builder@famillegratton.net)

* Fri Aug 11 2023 builder <builder@famillegratton.net> 1.000-0
- new package built with tito

