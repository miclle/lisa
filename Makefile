LISA_GO_EXECUTABLE ?= go
VERSION := $(shell git describe --tags)
DIST_DIRS := find * -type d -exec

build:
	${LISA_GO_EXECUTABLE} get ./...
	${LISA_GO_EXECUTABLE} build -o lisa -ldflags "-X main.version=${VERSION}" lisa.go

install: build
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./lisa ${DESTDIR}/usr/local/bin/lisa

test:
	${LISA_GO_EXECUTABLE} test .

clean:
	rm -f ./lisa.test
	rm -f ./lisa
	rm -rf ./dist

bootstrap-dist:
	${LISA_GO_EXECUTABLE} get -u github.com/mitchellh/gox

build-all:
	gox -verbose \
	-ldflags "-X main.version=${VERSION}" \
	-os="linux darwin windows " \
	-arch="amd64 386" \
	-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}" .

dist: build-all
	cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar -zcf lisa-${VERSION}-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r lisa-${VERSION}-{}.zip {} \; && \
	cd ..


.PHONY: build test install clean bootstrap-dist build-all dist