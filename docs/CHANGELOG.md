# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html)
since 2.4.0. Version headings are plain text (no compare links), as entries may be
written before anything is published.

## Unreleased
n/a

## 3.0.0 - 2026-10-22

### Added
- Added git repository creation support (`-G` | `--git`) flags so now you create the scaffolding and have it commited/pushed to git

  (have a look at the [README.md](../README.md) file for more information

### Changed

- Documentation, dependencies update


## 2.11.0 - 2026-10-22

### Changed

- Updated Default GO version to 1.27.3 in both the template generator and main code
- New directories (`images/`, `docs/`)
- Moved {ISSUES,ROADMAP,CHANGELOG,TODO}.md in `docs/`

## 2.10.0 - 2026-09-27

### Changed

- Changed destination Nexus repos in assets and source

## 2.9.1 - 2026-09-15

### Added

- `version-bump.sh` (repo root and scaffolded into every new project) to sync `versionnumber` and reset every present distro's release/pkgrel counter in one step

### Changed

- Binary version/date are now injected via `-ldflags -X` at build time instead of baked in at generation time
- `go vet`/`go test` now run inline before every packaging build (Alpine, ArchLinux, Debian, RedHat, Windows, `build.sh`)
- Manifest's `releasenumber` renamed to `defaultreleasenumber`

### Removed

- ArchLinux's dead `check()` phase

## 2.8.2 - 2026-08-30

### Changed

- Windows installs now default to a flat `C:\utils` (no per-product subdir, no Program Files), overridable via the generated `__windows/Makefile`'s `INSTALLDIR` variable instead of a create-time flag

### Removed

- The `-t/--target` flag and the manifest's `target` field it seeded

## 2.8.1 - 2026-08-26

### Added

- New Windows `targetdir` flag to control where a package installs

### Changed

- Some assets moved in new directories

## 2.8.0 - 2026-08-26

### Added

- Windows packaging support (`-w`): `__windows/` stub generates a `.wxs` (built into a real `.msi` via `wixl`/`msitools`, no Windows/.NET/mono needed) and a `Makefile`. Unlike every other stub, `refresh -w` never re-renders an existing `.wxs` -- its `UpgradeCode`/Component `Guid` are permanent per-project identifiers, generated once at `create` time

## 2.7.2 - 2026-08-21

### Added

- **ARCHBUILDER**: Safeguards when cleaning the build environment

## 2.7.1 - 2026-08-16

### Changed

- The `-e` flag is no longer mandatory

## 2.7.0 - 2026-08-16

### Changed

- Resynched asset generation with current-days workflow
- Various assets enhancements

### Removed

- Group management from generated assets (debian)
- Shell completion code

## 2.6.0 - 2026-07-20

### Changed

- Stub generation now takes into account the new CICD pipelines in $BUILDER packaging dir
- `refresh` command now re-generates `{{ SOFTWARE NAME }}.json`

## 2.5.1 - 2026-07-17

### Added

- `stubber.json`

### Changed

- All initial package releases will start at 1, now

## 2.5.0 - 2026-07-16

### Added

- A `refresh` command to change all packaging metadata

## 2.4.0 - 2026-07-08

### Changed

- New version numbering scheme, aligned with SemVer

### Fixed

- Missing embedded file issue

## 2.03.00 - 2026-07-03

_Unpublished._

## 2.02.00 - 2026-06-13

### Changed

- Updated packaging files in assets

## 2.01.00 - 2026-05-18

### Added

- Command completion at package install-time

## 2.00.01 - 2026-05-11

### Fixed

- Assets templates fixes

## 2.00.00 - 2026-05-02

### Added

- Archlinux distro to the stub as `__archlinux/`

### Changed

- Moved RPM building into `__redhat/`

## 1.95.00 - 2026-04-23

### Changed

- Upgraded to GO 1.26.2
- Improved build and builddeps scripts

## 1.94.01 - 2026-02-05

### Changed

- GO upgrade: 1.25.7

### Fixed

- Build script change to honour the -b flag

## 1.94.00 - 2026-01-29

### Added

- {TODO,ROADMAP}.md in assets

### Changed

- Modified the -v output

## 1.93.00 - 2026-01-06

### Fixed

- Fixes to build script

## 1.91.00 - 2025-12-14

### Changed

- GO version bump

### Removed

- GO environment variables from APKBUILD build() steps as this is handled by apkbuilder

## 1.90.01 - 2025-11-10

### Changed

- GO version bump
- Build scripts update (in core and in assets)
- Build scripts cleanup

## 1.84.00 - 2025-11-03

### Changed

- Completed PIE-linking disabling

### Removed

- Un-needed `__alpine/Makefile`

## 1.83.00 - 2025-11-02

### Changed

- Disabling gcompat and various link options
- Binaries are stripped at build-time now

## 1.82.00 - unknown

_Change doc is lost._

## 1.81.02 - 2025-10-09

### Added

- Enabled shell completion

### Changed

- GO version bump
- Upgraded customError and helperFunctions

## 1.80.02 - 2025-09-23

### Changed

- Renamed the import checker script

## 1.80.00 - 2025-09-12

### Added

- A new script in src/ to check for cyclic imports

### Changed

- GO version bump

## 1.78.00 - 2025-07-01

### Changed

- Disable CGO_ENABLED when building the tool (also updated assets)

## 1.77.00 - 2025-06-14

### Added

- A new template variable for version and changelog dates

## 1.76.00 - 2025-06-14

### Fixed

- Broken asset generation in build.sh

## 1.75.01 - 2025-05-09

### Fixed

- rpm-builddeps.sh, broken for a long time

## 1.75.00 - 2025-03-29

### Changed

- GO version bump
- build.sh update

## 1.74.00 - 2025-03-14

### Changed

- GO version bump

### Fixed

- APK fixes

## 1.72.00 - 2024-09-15

### Changed

- Disable GHA for now

## 1.71.00 - 2024-08-27

Last version before bringing git support.

### Changed

- Go version bump in binary, assets, templates

## 1.70.00 - 2024-08-11

### Changed

- Specfile is now arch-independent
- GO version bump

## 1.65.01 - 2024-08-05

### Added

- Github action code to publish releases

## 1.62.00 - 2024-07-28

### Changed

- Assets updates

## 1.61.00 - 2024-05-25

### Changed

- GO version bump
- Updated updateBuildDeps.sh

## 1.60.00 - 2024-04-05

### Changed

- GO version bump
- Rewrite of build.sh

## 1.55.01 - 2024-03-15

### Fixed

- APKBUILD now respects the -u parameter

## 1.55.00 - 2024-02-16

### Fixed

- More packaging fixes

## 1.54.00 - 2024-02-14

### Changed

- Go version bump

### Removed

- arm64 support

## 1.53.01 - 2024-01-11

### Changed

- Go version bump
- Build deps version bump
- Software version bump

## 1.53.00 - 2024-01-09

### Changed

- Removed/renamed some files

### Fixed

- Misc fixes in assets generation

## 1.52.01 - 2023-12-29

### Changed

- New version numbering scheme

### Fixed

- Minor fixes

## 1.505 - 2023-08-18

### Removed

- An undefined command from cmd/root.go

## 1.500 - 2023-08-16

### Added

- Update feature

### Fixed

- Issue where whatever executed command fails to return to CWD
- Filepath creation
- Arch name in Alpine's Makefile
- Missing GO VERSION placeholder in Alpine's APKBUILD asset
- Issue on where to create the structure when -p is omitted

## 1.206 - 2023-08-13

### Added

- gitignore

## 1.205 - 2023-08-13

### Fixed

- Missing placeholder in RPM stub

## 1.201 - 2023-08-13

### Fixed

- RPM fixes on assets generation

## 1.200 - 2023-08-12

### Added

- First iteration of go generate implementation in stubs

## 1.100-1 - 2023-08-12

### Fixed

- Missing control script in a debian control script

## 1.100 - 2023-08-12

### Added

- Extra debian packaging scripts

### Fixed

- Misnamed Archs

## 1.010 - 2023-08-12

### Added

- CHANGELOG.md to assets

### Fixed

- Version and Release flags had been accidentally removed. Restored

### Removed

- "IN THIS BRANCH.md" and src/helpers

## 1.000 - 2023-08-11

First "prod" version.

## 0.100 - 2023-06

Initial work.
