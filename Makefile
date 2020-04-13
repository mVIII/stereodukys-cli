PROJECT_NAME ?= stereodukys-cli
SHELL 	 := /bin/bash
executable := stereodukys-cli

clean:
	@rm -rf bin/*

build-cli:
	@echo Building $(executable)
	GOOS=linux GO111MODULE=on go build -o bin/$(executable) main.go

build-cli-windows:
	@echo Building $(executable)
	GOOS=windows GOARCH=386 go build -o bin/$(executable).exe main.go

help : Makefile
	@sed -n 's/^##//p' $<
