#!/usr/bin/env bash

PKGDIR=stubber-1.81.01-0_amd64

mkdir -p ${PKGDIR}/opt/bin ${PKGDIR}/DEBIAN
for i in control preinst prerm postinst postrm;do
  mv $i ${PKGDIR}/DEBIAN/
done

echo "Installing assets generator (go-bindata)"
cd ../src
GOBIN=$HOME/bin go install -a github.com/go-bindata/go-bindata/...@latest
cd templates
rm -f assets.go
echo "Generating assets"
go generate
echo "Building binary"
cd ..
CGO_ENABLED=0 go build -o ../__debian/${PKGDIR}/opt/bin/stubber .
strip ../__debian/${PKGDIR}/opt/bin/stubber
sudo chown 0:0 ../__debian/${PKGDIR}/opt/bin/stubber

echo "Software built. Now packaging..."
cd ../__debian/
dpkg-deb -b ${PKGDIR}
echo "Package built"
