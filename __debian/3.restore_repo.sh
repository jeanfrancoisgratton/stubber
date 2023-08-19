#!/usr/bin/env bash

git restore control preinst prerm postinst postrm
git checkout ../src/templates/assets.go
rm -rf stubber*
