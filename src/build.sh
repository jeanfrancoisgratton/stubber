#!/usr/bin/env sh


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


echo "Building ${OUTPUT}/${FULLNAME}"
go build -o ${OUTPUT}/${FULLNAME} .
