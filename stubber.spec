%define debug_package   %{nil}
%define _build_id_links none
%define _name   stubber
%define _prefix /opt
%define _version 1.505
%define _rel 2
%define _arch x86_64
%define _binaryname stubber

Name:       stubber
Version:    %{_version}
Release:    %{_rel}
Summary:    stubber

Group:      Utils
License:    GPL2.0
URL:        https://github.com/jeanfrancoisgratton/stubber

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: gcc
#Requires: sudo
#Obsoletes: vmman1 > 1.140

%description
Push binary package to NxRM

%prep
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src/templates
rm -f assets.go
sudo GOBIN=/opt/go/bin /opt/go/bin/go install -a github.com/go-bindata/go-bindata/...@latest
sudo /opt/go/bin/go generate
cd ..
/opt/go/bin/go build -o %{_sourcedir}/%{_binaryname} .
strip %{_sourcedir}/%{_binaryname}


%clean
rm -rf $RPM_BUILD_ROOT

%pre
exit 0

%install
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{_binaryname}


%changelog
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

