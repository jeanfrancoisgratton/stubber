#!/usr/bin/env sh

OUTPUT=/opt/bin

if [ "$#" -gt 0 ]; then
    OUTPUT=$1
fi
rm -f templates/assets.go
cd templates
go generate
cd ..
go build -o ${OUTPUT}/stubber .

