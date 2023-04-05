#!/usr/bin/make -f

install:
	go install ./cmd/swapchaind

build:
	go build -o=build/swapchaind ./cmd/swapchaind

build-docker:
	@DOCKER_BUILDKIT=1 docker build -t=swapchaind .

interchaintest:
	@CGO_ENABLED=1 cd interchaintest && go test --race -v -run ^TestInterchainSwap$$ .

start-local-swapchain:
	ignite chain serve --reset-once --skip-proto

start-local-osmosis:
	sh contrib/start-local-osmosis.sh

start-relayer:
	sh contrib/start-relayer.sh

.PHONY: install build build-docker interchaintest start-local-swapchain start-local-osmosis start-relayer
