.DEFAULT_GOAL:=build
BINARY:=server
PACKAGE:=github.com/cudneys/password-validator

VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
COMMIT_HASH="$(git rev-parse --short HEAD)"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

clean:
	rm -rvf dist/

prep:
	mkdir -p dist
	mkdir -p dist/linux/amd64
	mkdir -p dist/linux/arm
	mkdir -p dist/darwin/amd64
	mkdir -p dist/darwin/m1
	mkdir -p dist/windows/amd64

build: clean prep
	swag init
	GOOS=linux GOARGH=amd64 go build -o dist/linux/amd64/${BINARY} --ldflags="-X ${PACKAGE}/version.Version=$(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//') -X ${PACKAGE}/version.BuildTimestamp=$(shell date '+%Y-%m-%dT%H:%M:%S') -X ${PACKAGE}/version.CommitHash=$(shell git rev-parse --short HEAD)"
	GOOS=linux GOARGH=arm64 go build -o dist/linux/arm/${BINARY} --ldflags="-X ${PACKAGE}/version.Version=$(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//') -X ${PACKAGE}/version.BuildTimestamp=$(shell date '+%Y-%m-%dT%H:%M:%S') -X ${PACKAGE}/version.CommitHash=$(shell git rev-parse --short HEAD)"
	GOOS=darwin GOARGH=amd64 go build -o dist/darwin/and64/${BINARY} --ldflags="-X ${PACKAGE}/version.Version=$(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//') -X ${PACKAGE}/version.BuildTimestamp=$(shell date '+%Y-%m-%dT%H:%M:%S') -X ${PACKAGE}/version.CommitHash=$(shell git rev-parse --short HEAD)"
	GOOS=darwin GOARGH=arm64 go build  -o dist/darwin/m1/${BINARY} --ldflags="-X ${PACKAGE}/version.Version=$(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//') -X ${PACKAGE}/version.BuildTimestamp=$(shell date '+%Y-%m-%dT%H:%M:%S') -X ${PACKAGE}/version.CommitHash=$(shell git rev-parse --short HEAD)"
	GOOS=windows GOARGH=amd64 go build  -o dist/windows/and64/${BINARY}.exe --ldflags="-X ${PACKAGE}/version.Version=$(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//') -X ${PACKAGE}/version.BuildTimestamp=$(shell date '+%Y-%m-%dT%H:%M:%S') -X ${PACKAGE}/version.CommitHash=$(shell git rev-parse --short HEAD)"

docker: build
	docker build -t ${PACKAGE}:${VERSION}