#!/usr/bin/env ash

echo "Installing the Go binaries"
sudo rm -rf /opt/go
sudo /opt/bin/install_golang.sh `cat ../go.version` amd64

