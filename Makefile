APP=hopper
MAIN=cmd/main.go
BINARY=bin/${APP}
PKG:=github.com/iamasmith/${APP}
VP:=${PKG}/internal/version
LATEST_TAG:=$(shell git describe --tags --abbrev=0 HEAD 2>/dev/null)
ifeq ($(LATEST_TAG),)
LATEST_TAG:=v0.0.0
endif
RELEASE_TAG=$(shell git tag --points-at HEAD)
CURRENT_SHA=$(shell git rev-parse HEAD)
VERSION:=${LATEST_TAG}
ifeq ($(LATEST_TAG),$(RELEASE_TAG))
BUILDTYPE:=release
else
BUILDTYPE:=snapshot-${CURRENT_SHA}
endif
ifeq ($(BUILDID),)
BUILDTYPE:=$(shell whoami)-workstation-${CURRENT_SHA}
BUILDID:=None
endif
BUILDDATE:=$(shell date -u)

all: test build

test:
	go test --cover "./internal/..."

build:
	CGO_ENABLED=0 go build -o ${BINARY} -ldflags '-X "${VP}.LName=${APP}" -X "${VP}.LVersion=${VERSION}" -X "${VP}.LBuildDate=${BUILDDATE}" -X "${VP}.LBuildType=${BUILDTYPE}" -X "${VP}.LBuildId=${BUILDID}"' ${MAIN}

clean:
	go clean
	rm ${BINARY}

coverreport:
	go test --coverprofile cover.out "./internal..."
	go tool cover -html cover.out
	rm cover.out