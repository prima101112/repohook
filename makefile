# Check for required command tools to build or stop immediately
EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
     $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

ROOT_DIR:=${GOPATH}/bin 

BINARY=repohook
VERSION=0.1
BUILD=`git rev-parse HEAD`
PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

default: build

all: clean build install

build:
	go build ${LDFLAGS} -o ${BINARY}

install:
	go install ${LDFLAGS}

# Remove only what we've created
clean:
	find ${ROOT_DIR} -name '${BINARY}' -delete

.PHONY: check clean install build_all all