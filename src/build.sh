#!/usr/bin/env sh

set -eu

BRANCH=`git rev-parse --abbrev-ref HEAD`
BRANCH=$(echo "$BRANCH" | tr '/' '_')
BINARY=stubber
OUTPUT=/opt/bin
CHECK_PERMS=0

# Parse arguments
while [ "$#" -gt 0 ]; do
    case "$1" in
        -c|--checkperms)
            CHECK_PERMS=1
            ;;
        *)
            OUTPUT="$1"
            ;;
    esac
    shift
done

if [ "$BRANCH" = "master" ] || [ "$BRANCH" = "main" ] || [ "$BRANCH" = "develop" ]; then
    FULLNAME="$BINARY"
else
    FULLNAME="$BINARY-$BRANCH"
fi

echo
go vet ./...
echo
go test ./...

VERSION=$(sed -n 's/.*"versionnumber": *"\([^"]*\)".*/\1/p' ../stubber.json)
MODPATH=$(go list -m)
BUILDDATE=$(date -u +%Y-%m-%d)

echo
echo "Building ${OUTPUT}/${FULLNAME}"
echo
go build -ldflags="-X ${MODPATH}/cmd.buildVersion=${VERSION} -X ${MODPATH}/cmd.buildDate=${BUILDDATE}" -o ${OUTPUT}/${FULLNAME} .
