#!/usr/bin/env bash

PKGDIR=stubber-2.02.00-0_amd64

mkdir -p ${PKGDIR}/opt/bin ${PKGDIR}/DEBIAN
for i in control preinst prerm postinst postrm;do
  mv $i ${PKGDIR}/DEBIAN/
done

echo "Building binary"
cd ../src
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -buildid=" -o ../__debian/${PKGDIR}/opt/bin/stubber .
sudo chown 0:0 ../__debian/${PKGDIR}/opt/bin/stubber

echo "Software built. Now packaging..."
cd ../__debian/
dpkg-deb -b ${PKGDIR}
echo "Package built"
