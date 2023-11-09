#!/usr/bin/env sh



if [ "$#" -gt 0 ]; then
    BINARY={{ BINARY NAME }}
fi

go build -o /opt/bin/${BINARY} .
