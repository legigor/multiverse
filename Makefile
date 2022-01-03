CURRENT_DIR=$(shell pwd)
DIST_DIR=${CURRENT_DIR}/dist
BIN_NAME=multiverse

HOST_OS:=$(shell go env GOOS)
HOST_ARCH:=$(shell go env GOARCH)

build:
	@CGO_ENABLED=0 GOOS=${HOST_OS} GOARCH=${HOST_ARCH} go build -v -o ${DIST_DIR}/${BIN_NAME} ./cmd

test:
	@go test -v ./...
