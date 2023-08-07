all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./stubber.exe

prepare:
	sudo /opt/bin/install_golang.sh ___GOVERSION___ ___ARCHITECTURE___
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/___BINARYNAME___.exe $(DESTDIR)$(PREFIX)/bin/___BINARYNAME___
	strip $(DESTDIR)$(PREFIX)/bin/___BINARYNAME___
