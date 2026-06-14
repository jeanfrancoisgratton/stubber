<img src="./images/stubber_banner.png" alt="stubber logo" height="384" width="768" />

# stubber
___
A CLI tool that generates the directory structure, build scripts, and packaging
metadata for a new Go project, following a personal CI/CD convention (Alpine
APK, Debian/Ubuntu DEB, Archlinux .pkg.tar.zst and RHEL/Fedora RPM packaging, plus a ready-to-build
Go skeleton).

Everything is generated from templates embedded in the binary
(`src/templates/assets.go`, built from `src/assets/`), with placeholders
substituted at run time from the flags you pass.

---

## Contents

- [How it works](#how-it-works)
- [Installation](#installation)
- [Quick start](#quick-start)
- [Usage](#usage)
  - [Global flags](#global-flags)
  - [`create` flags](#create-flags)
  - [Examples](#examples)
- [Generated layout](#generated-layout)
  - [Skeleton (`-k`)](#skeleton--k)
  - [Alpine (`-a`)](#alpine--a)
  - [Debian (`-d`)](#debian--d)
  - [RedHat (`-r`)](#redhat--r)
- [Template placeholders](#template-placeholders)
- [Shell completion](#shell-completion)
- [Building from source](#building-from-source)
- [Building packages](#building-packages)
- [Known issues / caveats](#known-issues--caveats)
- [License](#license)

---

## How it works

`stubber create` takes a software name and one or more "stub" flags
(`-a`, `-d`, `-r`, `-A`, `-k`). Each stub flag tells stubber which set of templated
files to render into the project directory. Other flags (`-V`, `-D`,
`-M`, `-u`, etc.) supply the values used to fill in the placeholders in those
templates.

stubber only creates files — it does **not** initialize a git repository or
run `go mod init` for you. See [Caveats](#known-issues--caveats).

## Installation

Pre-built packages (`.apk`, `.deb`, `.rpm`, `.pkg.tar.zst` ) are published on the
[Releases](https://github.com/jeanfrancoisgratton/stubber/releases) page.

Otherwise, build it yourself — see [Building from source](#building-from-source).

## Quick start

```sh
# Full project: Go skeleton + all three packaging stubs
stubber create -a -d -r -k -A \
  -V 1.0.0 -R 1 \
  -D "My awesome CLI tool" \
  -u "https://git.famillegratton.net:3000/jfgratton/mytool" \
  mytool
```

This creates a `mytool/` directory in the current working directory containing
a buildable Go skeleton plus `__alpine/`, `__debian/`,`__archlinux/`, `__redhat/` packaging stubs.

> **Always review the generated files.** stubber is a *generic* stub
> generator — it doesn't know anything about your project beyond what you
> pass on the command line.

## Usage

```sh
stubber [global flags] <command> [command flags] <args>
```

Available commands:

| Command      | Description                                                |
|--------------|-------------------------------------------------------------|
| `create`     | Generates the directory structure (skeleton/stubs) for a new software project |
| `completion` | Generates shell completion scripts (bash, zsh)             |

Run `stubber -h`, `stubber create -h`, or `stubber completion -h` for the
built-in help at any time.

### Global flags

These apply to `stubber` itself and are inherited by all subcommands:

| Flag | Shorthand | Default | Description |
|------|-----------|---------|-------------|
| `--quiet` | `-q` | `false` | Silence non-essential output. |
| `--projectrootdir` | `-p` | `.` | Directory in which to create the project. If left at the default, stubber creates and uses `./<SOFTWARENAME>`. If set to an existing path, stubber generates files directly into it. |
| `--binaryname` | `-b` | *(same as `<SOFTWARENAME>`)* | Name of the compiled binary. Used to name Alpine packaging scripts (`<binaryname>.post-install`, etc.) and as the `{{ BINARY NAME }}` placeholder. |
| `--gover` | `-g` | `1.25.7` | Go version to embed in generated files (`go.mod`, `go.version`, packaging scripts). *Despite the built-in help text, this sets the Go version — not an output path.* |
| `--version` | | | Print stubber's own version and Go toolchain version. |
| `--help` | `-h` | | Show help for any command. |

### `create` flags

`stubber create` requires **exactly one** positional argument (the software
name) and **at least one** of `-a`, `-d`, `-r`, `-k`:

```sh
stubber create [-a] [-d] [-r] [-k] [-A] [other flags] <SOFTWARENAME>
```

| Stub flag     | Shorthand | What it generates                                                                                                            |
|---------------|-----------|------------------------------------------------------------------------------------------------------------------------------|
| `--alpine`    | `-a`      | `__alpine/` — APKBUILD + install/upgrade/deinstall scripts                                                                   |
| `--archlinux` | `-A`      | `__archlinux/` — `PKGBUILD`, `1.install-build-deps.sh`, `2.build-package.sh`, `<SOFTWARENAME>.install`                       |
| `--debian`    | `-d`      | `__debian/` — control file, build scripts, maintainer scripts                                                                |
| `--redhat`    | `-r`      | `__redhat/` — `<SOFTWARENAME>.spec`, `Makefile`, updateChangelog.sh, etc                                                     |
| `--skeleton`  | `-k`      | Go project skeleton: `src/`, `go.mod`, `main.go`, `cmd/root.go`, build scripts, `README.md`, `CHANGELOG.md`, `LICENSE`, etc. |

Value flags used to fill in the templates:

| Flag | Shorthand | Default | Description |
|------|-----------|---------|-------------|
| `--packagever` | `-V` | *(empty)* | Package version number — `{{ PACKAGE VERSION }}`. |
| `--packagerel` | `-R` | *(empty)* | Package release number — `{{ PACKAGE RELEASE }}`. |
| `--desc` | `-D` | *(empty)* | Package description — `{{ DESCRIPTION }}`. |
| `--maintainer` | `-M` | `Jean-Francois Gratton <jean-francois@famillegratton.net>` | Maintainer name/email — `{{ MAINTAINER }}`. **Override this if you're not me.** |
| `--packager` | `-P` | `APK Builder <builder@famillegratton.net>` | Packager name/email (Alpine) — `{{ PACKAGER }}`. **Override this too.** |
| `--section` | `-s` | `Packaging tool` | Debian package section — `{{ PACKAGE SECTION }}`. |
| `--depends` | `-e` | *(empty)* | Package dependencies, passed straight into the Debian `control` file — `{{ DEPENDENCIES }}`. |
| `--url` | `-u` | `https://git.famillegratton.net:3000/ADD_URL_HERE` | Upstream/git repo URL — `{{ URL }}`. **Override this.** |

### Examples

Generate just a Go skeleton, targeting Go 1.26.2, in a custom directory:

```sh
stubber create -k -g 1.26.2 -p ./projects/mytool mytool
```

Add Debian packaging to an existing project, with dependencies:

```sh
stubber create -d \
  -D "Backup helper for Docker volumes" \
  -M "Your Name <you@example.com>" \
  -e "docker.io" \
  dvol
```

Generate RPM packaging only, with version/release set explicitly:

```sh
stubber create -r -V 2.3.0 -R 1 -u "https://git.example.com/me/myproj" myproj
```

Generate everything quietly (e.g. from a script), with a binary name that
differs from the project name:

```sh
stubber -q create -a -d -r -k -A -b mytoolctl -V 0.1.0 -R 1 mytool
```

## Generated layout

The exact set of files generated depends on which of `-a` / `-d` / `-r` / `-k` / `-A`
you pass. Combined, a full run (`-a -d -r -k -A`) produces:

```
.
├── __alpine/
│   ├── APKBUILD
│   ├── <binary>.post-install
│   ├── <binary>.pre-install
│   ├── <binary>.pre-upgrade
│   ├── <binary>.post-upgrade
│   ├── <binary>.pre-deinstall
│   └── <binary>.post-deinstall
├── __debian/
│   ├── 1.install-build-deps.sh
│   ├── 2.build_binary.sh
│   ├── 3.restore_repo.sh
│   ├── control
│   ├── preinst
│   ├── postinst
│   ├── prerm
│   └── postrm
├── <SOFTWARENAME>.spec
├── rpmbuild-deps.sh
├── .gitignore
├── CHANGELOG.md
├── ISSUES.md
├── LICENSE
├── README.md
├── go.version
└── src
    ├── build.sh
    ├── go.mod
    ├── main.go
    ├── updateBuildDeps.sh
    └── cmd
        └── root.go
```

### Skeleton (`-k`)

A minimal but buildable Go project: `src/main.go`, `src/cmd/root.go` (a Cobra
root command), `src/go.mod`, plus `build.sh` / `updateBuildDeps.sh` helper
scripts and the usual `README.md` / `CHANGELOG.md` / `ISSUES.md` / `LICENSE` /
`.gitignore` / `go.version` files at the project root.

### Alpine (`-a`)

`__alpine/APKBUILD` plus one maintainer script per package lifecycle hook
(`post-install`, `pre-install`, `pre-upgrade`, `post-upgrade`, `pre-deinstall`,
`post-deinstall`), each named `<binaryname>.<hook>` and made executable.

### Debian (`-d`)

`__debian/control` plus build helper scripts (`1.install-build-deps.sh`,
`2.build_binary.sh`, `3.restore_repo.sh`) and maintainer scripts (`preinst`,
`postinst`, `prerm`, `postrm`), all made executable.

### RedHat (`-r`)

`<SOFTWARENAME>.spec` and `rpmbuild-deps.sh` at the project root.

## Template placeholders

If you customize the templates under `src/assets/` (see
[Building from source](#building-from-source)), these are the placeholders
stubber substitutes, and which flag/value feeds each one:

| Placeholder | Source |
|-------------|--------|
| `{{ SOFTWARE NAME }}` | The `<SOFTWARENAME>` positional argument |
| `{{ BINARY NAME }}` | `-b` / `--binaryname` (defaults to software name) |
| `{{ PACKAGE VERSION }}` | `-V` / `--packagever` |
| `{{ PACKAGE RELEASE }}` | `-R` / `--packagerel` |
| `{{ DESCRIPTION }}` | `-D` / `--desc` |
| `{{ MAINTAINER }}` | `-M` / `--maintainer` |
| `{{ PACKAGER }}` | `-P` / `--packager` (Alpine only) |
| `{{ SECTION }}` / `{{ PACKAGE SECTION }}` | `-s` / `--section` |
| `{{ DEPENDENCIES }}` | `-e` / `--depends` (Debian only) |
| `{{ URL }}` | `-u` / `--url` |
| `{{ GO VERSION }}` | `-g` / `--gover` |
| `{{ GO MAJOR MINOR }}` | Derived from `{{ GO VERSION }}` (e.g. `1.26.2` → `1.26`) |
| `{{ ARCHITECTURE }}` | Hardcoded `amd64` (Debian); Alpine maps `amd64` → `x86_64` internally |
| `{{ RELEASE DATE }}` | Today's date (`YYYY.MM.DD`), generated automatically |

## Shell completion

stubber can generate Bash and Zsh completion scripts via Cobra:

```sh
# Bash, current session
source <(stubber completion bash)

# Bash, persistent
stubber completion bash | sudo tee /etc/bash_completion.d/stubber > /dev/null

# Zsh, current session
source <(stubber completion zsh)

# Zsh, persistent
stubber completion zsh > ~/.zsh/stubber
echo 'fpath=($HOME/.zsh $fpath)' >> ~/.zshrc
echo 'autoload -Uz compinit && compinit' >> ~/.zshrc
```

Fish/PowerShell completions are not provided.

## Building from source

```sh
git clone https://github.com/jeanfrancoisgratton/stubber.git
cd stubber/src
```

To customize the generated output, edit the templates under `src/assets/`
(organized into `apk/`, `deb/`, `rpm/`, and `skeleton/`), then regenerate the
embedded asset bundle and build:

```sh
# regenerates src/templates/assets.go from src/assets/
./build.sh
```

The build script compiles with `CGO_ENABLED=0` and strips debug info
(`-ldflags="-s -w"`). The required Go version is tracked in `go.version`
at the repo root (currently 1.26.2).

## Building packages

The `__alpine/` and `__debian/` directories, plus `stubber.spec` and
`rpmbuild-deps.sh` at the repo root, are stubber's *own* packaging stubs
(generated by stubber itself, for itself). They're written for the author's
home-lab build containers, which aren't published.

If you'd rather not build from source, grab a pre-built package from the
[Releases](https://github.com/jeanfrancoisgratton/stubber/releases) page.

## Known issues / caveats

- stubber creates a directory structure, but does **not** initialize a git
  repository (or any other VCS) for you.
- After running `stubber create -k`, `go.mod` / `go.sum` in the generated
  skeleton may need a `go mod init` / `go mod tidy` pass before `src/build.sh`
  will succeed.
- Mainly tested on x86_64/amd64. Lightly tested on Apple Silicon (arm64);
  generated files may need extra tweaking on other architectures.
- The `-g`/`--gover` flag's built-in `-h` description is misleading (it says
  "Where to put the skeleton dir") — it actually sets the Go version used in
  generated files.
- **Always review the generated files.** This is a generic stub generator; it
  doesn't validate the values you pass, and downstream packaging tools (`apk`,
  `dpkg-buildpackage`, `rpmbuild`) are unforgiving of malformed metadata.

See [ISSUES.md](ISSUES.md) for the current open-issue tracker.

## License

[GPL-3.0](LICENSE)
