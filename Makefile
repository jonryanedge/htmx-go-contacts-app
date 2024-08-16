# Include variable from .envrc file
include .envrc

# Include server ops 
include ./servers/Makefile

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## run: run the web application
.PHONY: run
run:
	go run ./cmd/web

## air: autorun using AIR
.PHONY: air
air:
	air -c .air.toml

## build: build minimzed app and save to .tmp directory
.PHONY: build
build:
	go build -ldflags='-s' -o=./tmp/bin ./cmd/web
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./tmp/linux_amd64/igmp ./cmd/web

## host/ssh: ssh to host server
.PHONY: host/ssh
host/ssh:
	ssh -i ${USERKEY} root@${HOST}

## host/push: copy linux build to host server
.PHONY: host/push
host/push:
	scp -i ${USERKEY} ./tmp/linux_amd64/igmp root@${HOST}:/srv/igmp/igmp

## host/svc: copy service file to host server
.PHONY: host/svc
host/svc:
	scp -i ${USERKEY} ./igmp.service root@${HOST}:/etc/systemd/system/igmp.service
