%define debug_package   %{nil}
%define _build_id_links none
%define _name   stubber
%define _prefix              /opt
%define _bash_completionsdir /usr/share/bash-completion/completions
%define _zsh_completionsdir  /usr/share/zsh/site-functions
%define _version 2.5.0
%define _rel 0
%define _arch x86_64
%define _binaryname stubber

Name:       stubber
Version:    %{_version}
Release:    %{_rel}
Summary:    stubber

Group:      Utils
License:    GPL2.0
URL:        https://git.famillegratton.net:3000/mainline/stubber.git

Source0:    %{name}-%{_version}.tar.gz
#BuildArchitectures: x86_64
BuildRequires: gcc
Requires: bash-completion

%description
Creates a GO software skeleton

%prep
%autosetup

%build
cd src
CGO_ENABLED=0 /opt/go/bin/go build -trimpath -ldflags="-s -w -buildid=" -o %{_builddir}/%{name}-%{version}/%{_binaryname} .


%clean
rm -rf $RPM_BUILD_ROOT

%pre

%install
rm -rf %{buildroot}
install -Dpm 0755 %{_builddir}/%{name}-%{version}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post
# Bash completion — always install
mkdir -p /etc/usr/share/bash-completion/completions
/opt/bin/stubber completion bash > %{_bash_completionsdir}/stubber

# Zsh completion — only if zsh is present
if command -v zsh > /dev/null 2>&1; then
    mkdir -p %{_zsh_completionsdir}/zsh/site-functions
    /opt/bin/stubber completion zsh > %{_zsh_completionsdir}/_stubber
fi

%preun

%postun
if [ $1 -eq 0 ]; then
    # $1 == 0 means this is a full uninstall, not an upgrade
    rm -f %{_bash_completionsdir}/stubber
    rm -f %{_zsh_completionsdir}/_stubber
fi

%files
%defattr(-,root,root,-)
%attr(0775,root,root) %{_bindir}/%{_binaryname}

%changelog
* Thu Jul 16 2026 Binary package builder <builder@famillegratton.net> 2.5.0-0
- Fixed rpmbuild scripts, and others
- builddeps updates
- Added GO test files
- completed the refresh command
- rpmbuild enhancements to guard against pushing an existing tag

* Wed Jul 08 2026 Binary package builder <builder@famillegratton.net> 2.4.0-0
- fixed issue where some embedded files were missing; new SemVer-aligned version number

* Sun Jul 05 2026 Binary package builder <builder@famillegratton.net> 2.03.00-0
- fixed specfile name
- updated the rpmbuilder section
- re-updated README.md
- updated doc
- chore: update changelog for 2.02.00-0
- upgraded to GO 1.26.4, doc update, packaging script fixes
- chore: update changelog for 2.01.00-3
- Merge remote-tracking branch 'refs/remotes/origin/develop' into develop
- Fixed variable path, again
- chore: update changelog for 2.01.00-2
- moved _datadir to new macros
- Fixed typo in ARCH packaging scripts
- chore: update changelog for 2.01.00-1
- Fixed tool invocation in command completion as the tool might not be yet in the PATH var
- Fixed wrong filename
- chore: update changelog for 2.01.00-0
- Added install-time command completion support
- Merge branch 'develop' of ssh://git.famillegratton.net:9722/mainline/stubber into develop
- removed unnecessary go dependency
- Assets templates fixes
- cosmetic output fix
- Fixed path in DEB build scripts
- chore: update changelog for 2.00.00-0
- more build fixes, doc updates
- fixed rpmbuild in both templated assets and actual specfile
- Merge remote-tracking branch 'refs/remotes/origin/develop' into develop
- replaced sed with awk
- Fixed badly copied script
- cleaned the Makefile
- Completed addition of __redhat and __archlinux to the stub
- interim commmit
- interim commit

* Sun Jun 14 2026 Binary package builder <builder@famillegratton.net> 2.02.00-0
- upgraded to GO 1.26.4, doc update, packaging script fixes
- chore: update changelog for 2.01.00-3
- Merge remote-tracking branch 'refs/remotes/origin/develop' into develop
- Fixed variable path, again
- chore: update changelog for 2.01.00-2
- moved _datadir to new macros
- Fixed typo in ARCH packaging scripts
- chore: update changelog for 2.01.00-1
- Fixed tool invocation in command completion as the tool might not be yet in the PATH var
- Fixed wrong filename
- chore: update changelog for 2.01.00-0
- Added install-time command completion support
- Merge branch 'develop' of ssh://git.famillegratton.net:9722/mainline/stubber into develop
- removed unnecessary go dependency
- Assets templates fixes
- cosmetic output fix
- Fixed path in DEB build scripts
- chore: update changelog for 2.00.00-0
- more build fixes, doc updates
- fixed rpmbuild in both templated assets and actual specfile
- Merge remote-tracking branch 'refs/remotes/origin/develop' into develop
- replaced sed with awk
- Fixed badly copied script
- cleaned the Makefile
- Completed addition of __redhat and __archlinux to the stub
- interim commmit
- interim commit

* Mon May 18 2026 Binary package builder <builder@famillegratton.net> 2.01.00-3
- Merge remote-tracking branch 'refs/remotes/origin/develop' into develop
- Fixed variable path, again
- chore: update changelog for 2.01.00-2
- moved _datadir to new macros
- Fixed typo in ARCH packaging scripts
- chore: update changelog for 2.01.00-1
- Fixed tool invocation in command completion as the tool might not be yet in the PATH var
- Fixed wrong filename
- chore: update changelog for 2.01.00-0
- Added install-time command completion support
- Merge branch 'develop' of ssh://git.famillegratton.net:9722/mainline/stubber into develop
- removed unnecessary go dependency
- Assets templates fixes
- cosmetic output fix
- Fixed path in DEB build scripts
- chore: update changelog for 2.00.00-0
- more build fixes, doc updates
- fixed rpmbuild in both templated assets and actual specfile
- Merge remote-tracking branch 'refs/remotes/origin/develop' into develop
- replaced sed with awk
- Fixed badly copied script
- cleaned the Makefile
- Completed addition of __redhat and __archlinux to the stub
- interim commmit
- interim commit

* Mon May 18 2026 Binary package builder <builder@famillegratton.net> 2.01.00-2
- moved _datadir to new macros
- Fixed typo in ARCH packaging scripts
- chore: update changelog for 2.01.00-1
- Fixed tool invocation in command completion as the tool might not be yet in the PATH var
- Fixed wrong filename
- chore: update changelog for 2.01.00-0
- Added install-time command completion support
- Merge branch 'develop' of ssh://git.famillegratton.net:9722/mainline/stubber into develop
- removed unnecessary go dependency
- Assets templates fixes
- cosmetic output fix
- Fixed path in DEB build scripts
- chore: update changelog for 2.00.00-0
- more build fixes, doc updates
- fixed rpmbuild in both templated assets and actual specfile
- Merge remote-tracking branch 'refs/remotes/origin/develop' into develop
- replaced sed with awk
- Fixed badly copied script
- cleaned the Makefile
- Completed addition of __redhat and __archlinux to the stub
- interim commmit
- interim commit

* Mon May 18 2026 Binary package builder <builder@famillegratton.net> 2.01.00-1
- Fixed tool invocation in command completion as the tool might not be yet in the PATH var
- Fixed wrong filename
- chore: update changelog for 2.01.00-0
- Added install-time command completion support
- Merge branch 'develop' of ssh://git.famillegratton.net:9722/mainline/stubber into develop
- removed unnecessary go dependency
- Assets templates fixes
- cosmetic output fix
- Fixed path in DEB build scripts
- chore: update changelog for 2.00.00-0
- more build fixes, doc updates
- fixed rpmbuild in both templated assets and actual specfile
- Merge remote-tracking branch 'refs/remotes/origin/develop' into develop
- replaced sed with awk
- Fixed badly copied script
- cleaned the Makefile
- Completed addition of __redhat and __archlinux to the stub
- interim commmit
- interim commit

* Mon May 18 2026 Binary package builder <builder@famillegratton.net> 2.01.00-0
- Added install-time command completion support
- Merge branch 'develop' of ssh://git.famillegratton.net:9722/mainline/stubber into develop
- removed unnecessary go dependency
- Assets templates fixes
- cosmetic output fix
- Fixed path in DEB build scripts
- chore: update changelog for 2.00.00-0
- more build fixes, doc updates
- fixed rpmbuild in both templated assets and actual specfile
- Merge remote-tracking branch 'refs/remotes/origin/develop' into develop
- replaced sed with awk
- Fixed badly copied script
- cleaned the Makefile
- Completed addition of __redhat and __archlinux to the stub
- interim commmit
- interim commit

* Mon May 04 2026 Binary package builder <builder@famillegratton.net> 2.00.00-0
- more build fixes, doc updates
- fixed rpmbuild in both templated assets and actual specfile
- Merge remote-tracking branch 'refs/remotes/origin/develop' into develop
- replaced sed with awk
- Fixed badly copied script
- cleaned the Makefile
- Completed addition of __redhat and __archlinux to the stub
- interim commmit
- interim commit

* Thu Apr 23 2026 Binary package builder <builder@famillegratton.net> 1.95.00-0
- fixed missing files in assets (jean-francois@famillegratton.net)

* Thu Feb 05 2026 Binary package builder <builder@famillegratton.net> 1.94.01-0
- GO upgrade, build script logic fix (jean-francois@famillegratton.net)

* Thu Jan 29 2026 Binary package builder <builder@famillegratton.net> 1.94.00-0
- Added new files in assets (jean-francois@famillegratton.net)
- Added GO version value (jean-francois@famillegratton.net)
- build deps update (builder@famillegratton.net)

* Thu Jan 29 2026 Binary package builder <builder@famillegratton.net> 1.93.00-1
- Build script fix (jean-francois@famillegratton.net)

* Tue Jan 06 2026 Binary package builder <builder@famillegratton.net> 1.93.00-0
- updated the build script (jean-francois@famillegratton.net)

* Sun Dec 21 2025 Binary package builder <builder@famillegratton.net> 1.92.00-0
- Version bump, updated builddeps update check script (jean-
  francois@famillegratton.net)

* Sun Dec 14 2025 Binary package builder <builder@famillegratton.net> 1.91.00-0
- Fixed wrong release number (builder@famillegratton.net)
- Updated build dependencies (builder@famillegratton.net)
- Fixed go dep installation script (builder@famillegratton.net)
- GO version bump, removed environment vars from build() step in APKBUILD
  (jean-francois@famillegratton.net)
- added assets to repo (builder@famillegratton.net)
- Added a dependency script (jean-francois@famillegratton.net)
- tab indentation fix (jean-francois@famillegratton.net)
- Re-instated removed asset generation snippet (jean-
  francois@famillegratton.net)

* Mon Nov 10 2025 Binary package builder <builder@famillegratton.net> 1.90.01-0
- Re-instated sudo for go-bindata and go generate (builder@famillegratton.net)
- Automatic commit of package [stubber] release [1.90.01-0].
  (builder@famillegratton.net)
- Cleaned build scripts up (jean-francois@famillegratton.net)
- Fix missing dependency in build script (builder@famillegratton.net)
- builddeps update (builder@famillegratton.net)

* Mon Nov 10 2025 Binary package builder <builder@famillegratton.net> 1.90.00-1
- dependencies cleanup (builder@famillegratton.net)

* Mon Nov 10 2025 Binary package builder <builder@famillegratton.net> 1.90.00-0
- reverted tito tag (builder@famillegratton.net)
- more specfile fixes (jean-francois@famillegratton.net)
- specfile changelog cleanup (builder@famillegratton.net)
- Automatic commit of package [stubber] release [1.90.00-0].
  (builder@famillegratton.net)
- Automatic commit of package [stubber] release [1.90.00-0].
  (builder@famillegratton.net)
- Completed build script cleanup both in core and assets (jean-
  francois@famillegratton.net)
- asset specfile cleanup (jean-francois@famillegratton.net)
- version bump completed (jean-francois@famillegratton.net)

* Mon Nov 03 2025 Binary package builder <builder@famillegratton.net> 1.84.00-0
- resynched APKBUILD (builder@famillegratton.net)
- Completed removal of PIE-linking on Alpine (jean-francois@famillegratton.net)

* Mon Nov 03 2025 Binary package builder <builder@famillegratton.net> 1.83.00-0
- Removed gcompat dependency (builder@famillegratton.net)
- Removed gcompat from link chain (jean-francois@famillegratton.net)

* Sun Nov 02 2025 Binary package builder <builder@famillegratton.net> 1.82.00-0
- removed mentions of strip, tuned the build process in Alpine (jean-
  francois@famillegratton.net)
- stripping binaries at build-time (jean-francois@famillegratton.net)

* Thu Oct 09 2025 Binary package builder <builder@famillegratton.net> 1.81.02-0
- Fixed wrong program name in completion subcommand (jean-
  francois@famillegratton.net)

* Thu Oct 09 2025 Binary package builder <builder@famillegratton.net> 1.81.01-0
- Fixed missing completion file in assets/ (jean-francois@famillegratton.net)

* Thu Oct 09 2025 Binary package builder <builder@famillegratton.net> 1.81.00-0
- added shell completion feature (jean-francois@famillegratton.net)

* Thu Oct 09 2025 Binary package builder <builder@famillegratton.net> 1.80.03-0
- Package version number bump (builder@famillegratton.net)
- updated GO and builddeps (jean-francois@famillegratton.net)

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

