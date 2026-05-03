#!/usr/bin/env sh

cd src/templates
rm -f assets.go

go generate
