#!/usr/bin/env bash

PKGDIR="{{ SOFTWARE NAME }}-{{ PACKAGE VERSION }}-{{ PACKAGE RELEASE }}_{{ ARCHITECTURE }}"

mkdir -p ${PKGDIR}/opt/bin ${PKGDIR}/DEBIAN
mkdir -p ${PKGDIR}/opt/bin ${PKGDIR}/DEBIAN
for i in control preinst prerm postinst postrm;do
  mv $i ${PKGDIR}/DEBIAN/
done

echo "Building binary from source"
cd ../src
go build -o ../__debian/${PKGDIR}/opt/bin/{{ BINARY NAME }} .
strip ../__debian/${PKGDIR}/opt/bin/{{ BINARY NAME }}
chown 0:0 ../__debian/${PKGDIR}/opt/bin/{{ BINARY NAME }}

echo "Binary built. Now packaging..."
cd ../__debian/
dpkg-deb -b ${PKGDIR}
