VERSION=0.1.0
USER=bells17
REPOSITORY=hexip

all: gom ghr bundle build

init: gom ghr

gom:
	go get -u github.com/mattn/gom

ghr:
	go get -u github.com/tcnksm/ghr

bundle:
	gom -test install

build:
	gom build -ldflags '-X main.BuildVersion=${VERSION}' -o bin/hexip

install:
	install hexip /usr/local/bin/hexip

fmt:
	gom exec go fmt ./...

test:
	gom exec go test -v .

build-cross:
	GOOS=linux GOARCH=amd64 gom build -ldflags '-X main.BuildVersion=${VERSION}' -o bin/hexip-linux-amd64
	GOOS=darwin GOARCH=amd64 gom build -ldflags '-X main.BuildVersion=${VERSION}' -o bin/hexip-darwin-amd64

dist: build-cross
	cd bin && \
		tar cvf release/hexip-linux-amd64-${VERSION}.tar hexip-linux-amd64 && \
		zopfli release/hexip-linux-amd64-${VERSION}.tar && \
		rm release/hexip-linux-amd64-${VERSION}.tar
	cd bin && \
		tar cvf release/hexip-darwin-amd64-${VERSION}.tar hexip-darwin-amd64 && \
		zopfli release/hexip-darwin-amd64-${VERSION}.tar && \
		rm release/hexip-darwin-amd64-${VERSION}.tar

clean:
	rm -f bin/hexip*
	rm -f bin/release/hexip*

tag:
	git checkout master
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin master

release: clean dist
	rm -f bin/release/.gitkeep && \
		ghr -u ${USER} -r ${REPOSITORY} ${VERSION} bin/release && \
		touch bin/release/.gitkeep
