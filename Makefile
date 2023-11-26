.PHONY: build
build:
	go build -v ./ && ./clean

.DEFAULT_GOAL := build