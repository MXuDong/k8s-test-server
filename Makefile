# Author : A-Donga

# Make project

include Makefile.var

# get local arch
LOCAL_ARCH := $(shell uname -m)
ifeq ($(LOCAL_ARCH),x86_64)
    ARCH ?= amd64
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 5),armv8)
    ARCH ?= arm64
else ifeq ($(LOCAL_ARCH),aarch64)
    ARCH ?= arm64
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 4),armv)
    ARCH ?= arm
else
    $(error This system's architecture $(LOCAL_ARCH) isn't supported)
endif


build: pre-file
	go build -o ${OUTPUT_FILE}/app k8s-test-backend/cmd

# docker tag local image
.PHONY: local-image
local-image:
	docker build \
		--build-arg GOPROXY=$(GOPROXY) \
		-t ${HUB}/kts-common-test:${VERSION}\
		-f Dockerfile.build .

.PHONY: pre-file
pre-file:
	# clean dist file
	rm -rf ${OUTPUT_FILE}
	# make directory
	mkdir ${OUTPUT_FILE}

# for drone environment
createTime ?= ${DRONE_BUILD_FINISHED}
version ?= $(DRONE_BUILD_NUMBER)
platform ?= DRONE

# build for drone platform
.PHONY: drone-build
drone-build:
	# do package with go build
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=v0.0.$(version) -X main.BuildPlatform=$(platform) -X main.BuildStamp=$(createTime)" -o ./dist/app k8s-test-backend/cmd
