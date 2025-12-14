#!/usr/bin/env ash

echo "Installing the Go binaries"
sudo /opt/bin/install_golang.sh `cat ../go.version` amd64

