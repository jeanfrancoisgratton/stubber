# {{ SOFTWARE NAME }} — Windows (.msi) packaging
#
# Recipes:
#   make build     -> cross-compile {{ SOFTWARE NAME }}.exe and package it into
#                      a real .msi with msitools' wixl
#   make release   -> upload the freshly built .msi to the Nexus winLocal repo
#   make clean     -> remove the staging tree and the built .msi
#   make info      -> print the values extracted from {{ SOFTWARE NAME }}.json
#
# Windows has no distro package manager, so there's no APKBUILD/control/spec/
# PKGBUILD equivalent here. Metadata (name / version / binary / description)
# is read straight from ../{{ SOFTWARE NAME }}.json instead -- the manifest
# stubber itself writes at `create` time -- so this stays a single source of
# truth, same principle as every other __*/Makefile this tool generates.
#
# "Packaging" here means: cross-compile with GOOS=windows GOARCH=amd64
# CGO_ENABLED=0 (built inside the winbuilder container), then build a proper
# .msi installer around it with `wixl` (msitools -- a from-scratch, C,
# Linux-native reimplementation of the WiX toolchain; no Windows/.NET/mono
# needed to produce a real Windows Installer file). {{ SOFTWARE NAME }}.wxs
# is the WiX source; read the comment at its top before touching the
# UpgradeCode or the Component Guid inside it -- both are permanent
# identifiers, not build artifacts. Values wrapped in $(var.X) in that file
# are filled in below via `wixl -D X=value`; wixl substitutes those as plain
# string values, so passing $(DESCRIPTION) unescaped here is correct --
# XML-escaping only matters for text written directly into the .wxs source,
# not for -D values.

SHELL       := /bin/bash
.SHELLFLAGS := -o pipefail -c

MANIFEST     := ../{{ SOFTWARE NAME }}.json
MANUFACTURER := famillegratton.net

NAME        := $(shell jq -r '.softwarename'  $(MANIFEST))
VERSION     := $(shell jq -r '.versionnumber' $(MANIFEST))
BINARY      := $(shell jq -r '.binaryname'    $(MANIFEST))
DESCRIPTION := $(shell jq -r '.description'   $(MANIFEST))

NEXUSREPO := winLocal
SRCDIR    := ../src
ARCH      := amd64
STAGING   := $(NAME)-$(VERSION)_windows_$(ARCH)
MSI       := $(CURDIR)/$(STAGING).msi

.PHONY: build upload release clean info

build:
	rm -rf "$(STAGING)"
	install -d "$(STAGING)"
	cd $(SRCDIR) && CGO_ENABLED=0 GOOS=windows GOARCH=$(ARCH) go build -trimpath \
	  -ldflags="-s -w -buildid=" -o "$(CURDIR)/$(STAGING)/$(BINARY).exe" .
	wixl -D ProductName=$(NAME) -D ProductVersion=$(VERSION) \
	  -D Manufacturer=$(MANUFACTURER) -D "Description=$(DESCRIPTION)" \
	  -D Staging=$(STAGING) -D BinaryName=$(BINARY) \
	  -o "$(MSI)" {{ SOFTWARE NAME }}.wxs

upload:
	/opt/bin/nxtools upload $(NEXUSREPO) $(MSI)

release:
	${MAKE} build
	${MAKE} upload
	${MAKE} clean

clean:
	rm -rf "$(STAGING)" "$(MSI)"

info:
	@echo "NAME        = $(NAME)"
	@echo "VERSION     = $(VERSION)"
	@echo "BINARY      = $(BINARY)"
	@echo "DESCRIPTION = $(DESCRIPTION)"
	@echo "ARCH        = $(ARCH)"
	@echo "STAGING     = $(STAGING)"
	@echo "MSI         = $(MSI)"
