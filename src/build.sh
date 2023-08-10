#!/usr/bin/env sh

OUTPUT=/opt/bin

if [ "$#" -gt 0 ]; then
    OUTPUT=$1
fi

echo "Embedding resources..."
cd templates && rm -f assets.go
go generate
cd ..
echo "Building binary..."
go build -o ${OUTPUT}/stubber .

