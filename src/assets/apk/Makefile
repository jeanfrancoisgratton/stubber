all: 
	cd ../src
	CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -linkmode=internal -buildid=" . -o ./{{ BINARY NAME }}.exe

prepare:
	sudo /opt/bin/install_golang.sh {{ GO VERSION }} amd64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/{{ BINARY NAME }}.exe $(DESTDIR)$(PREFIX)/bin/{{ BINARY NAME }}
