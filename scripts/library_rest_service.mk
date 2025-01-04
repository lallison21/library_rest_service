PROJECT?=github.com/lallison21/library_rest_service
NAME?=library_rest_service
VERSION?=v0.0.1
BUILD_TIME?=$(shell date '+%Y-%m-%dT%H:%M:%S')

.PHONY: build

build:
	go build \
	-ldflags "-w -s \
	-X ${PROJECT}/version.Name=${NAME} \
	-X ${PROJECT}/version.Version=${VERSION} \
	-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
	-o bin/library_rest_service cmd/library_rest_service/library_rest_service.go
