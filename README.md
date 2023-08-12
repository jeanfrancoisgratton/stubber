<H1>stubber</H1>

Tool to create the directory structure to build and package a GO software using my current CI-CD infra.
<br><br>
____

<H2>How does it work</H2>
You use this tool to create the directory structure of a GO software according to my own CI-CD pipeline.
All files come from templated files and are generated in order to replace placeholders in the templates with values provided at run-time.

Basically, you provide values (using the various built-in flags) and all the necessary files will be created

Below is a directory tree of the files that would be generated by the tool (when all flags are enabled)
```
.
├── __alpine
│   ├── APKBUILD
│   └── Makefile
├── __debian
│   ├── 1.install-build-deps.sh
│   ├── 2.build_binary.sh
│   ├── 3.restore_repo.sh
│   ├── control
│   ├── current_pkg_release
│   └── preinst
├── FIXME.md
├── go.version
├── IN THIS BRANCH.md
├── LICENSE
├── PACKAGING.md
├── README.md
├── ROADMAP.md
├── rpm-install-build-deps.sh
├── src
│   ├── build.sh
│   ├── cmd
│   │   └── root.go
│   ├── go.mod
│   ├── helpers
│   │   ├── misc.go
│   ├── main.go
│   └── upgrade_pkgs.sh
├── stubber.spec
└── TODO.md
```

All sections (also called stubs) are generated this way:

| Stub          | Flag | File or directory                                   |
|---------------|------|-----------------------------------------------------|
| apk packaging | -a   | __alpine/                                           |
| deb packaging | -d   | __debian/                                           |
| rpm packaging | -r   | src/rpm-install-build.sh<br/>src/stubber.spec       |
| skeleton      | -k   | All other files under src/, src/helpers and src/cmd |

Of course, regarding the `-r` flag, you need to replace `stubber.spec` with the correct filename. All other flags (see `stubber -h` and `stubber create -h`) provide values to pass to the generated files.<br>
All templated files are embedded into `src/templates/assets.go` . The files were generated from `src/assets/` using the build script in `src/build.sh`

<H2>Building from source</H2>
Simple steps:<br>
- Clone this repo<br>
- Modify all the templated files in src/assets/ according to your tastes<br>
- Run: `./build.sh` from the src/ directory<br>

<H2>Building packages</H2>
Instructions are provided in [PACKAGING.md](PACKAGING.md).<br>
The instructions provided in that file are for use with my own build containers, which are not yet published as they rely too heavily on my own infra at home.<br>
In the meantime, you can use the packages under the Releases link herein.

