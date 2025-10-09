all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./stubber.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.25.2 amd64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/stubber.exe $(DESTDIR)$(PREFIX)/bin/stubber
	strip $(DESTDIR)$(PREFIX)/bin/stubber
