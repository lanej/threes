GO_BIN ?= go
VERSION := $(shell git describe --tags)
DIST_DIRS := find * -type d -exec

build:
	${GO_BIN} build -o threes -ldflags "-X main.version=${VERSION}" threes.go

install: build
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./threes ${DESTDIR}/usr/local/bin/threes

test:
	${GO_BIN} test . ./board

testv:
	${GO_BIN} test -v . ./board

integration-test:
	${GO_BIN} build
	./threes up
	./threes install

clean:
	rm -f ./threes.test
	rm -f ./threes
	rm -rf ./dist

bootstrap-dist:
	${GO_BIN} get -u github.com/franciscocpg/gox
	cd ${GOPATH}/src/github.com/franciscocpg/gox && git checkout dc50315fc7992f4fa34a4ee4bb3d60052eeb038e
	cd ${GOPATH}/src/github.com/franciscocpg/gox && ${GO_BIN} install


build-all:
	gox -verbose \
		-ldflags "-X main.version=${VERSION}" \
		-os="linux darwin windows freebsd openbsd netbsd" \
		-arch="amd64 386 armv5 armv6 armv7 arm64" \
		-osarch="!darwin/arm64" \
		-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}" .

dist: build-all
	cd dist && \
		$(DIST_DIRS) cp ../LICENSE {} \; && \
		$(DIST_DIRS) cp ../README.md {} \; && \
		$(DIST_DIRS) tar -zcf threes-${VERSION}-{}.tar.gz {} \; && \
		$(DIST_DIRS) zip -r threes-${VERSION}-{}.zip {} \; && \
		cd ..


.PHONY: build test install clean bootstrap-dist build-all dist integration-test
