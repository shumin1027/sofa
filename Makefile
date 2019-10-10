GIT_SHA=$(shell git rev-parse HEAD)
GIT_CLOSEST_TAG=$(shell git describe --abbrev=0 --tags)
DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

BUILD_INFO_IMPORT_PATH=xtc/sofa/pkg/version
BUILD_INFO=-ldflags "-X $(BUILD_INFO_IMPORT_PATH).commitSHA=$(GIT_SHA) -X $(BUILD_INFO_IMPORT_PATH).version=$(VERSION) -X $(BUILD_INFO_IMPORT_PATH).date=$(DATE)"

VERSION=v1.0.0

.PHONY: build
build:
	CGO_ENABLED=0 installsuffix=cgo go build -tags sofa -o ./bin/sofa $(BUILD_INFO) ./main.go

.PHONY: clean
clean:
	rm -rf ./bin/*

.PHONY: echo-version
echo-version:
	@echo $(GIT_CLOSEST_TAG)

