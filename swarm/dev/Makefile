.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/epvchain/go-epvchain/cmd/swarm
	go build -o bin/gepv     github.com/epvchain/go-epvchain/cmd/gepv
	go build -o bin/bootnode github.com/epvchain/go-epvchain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/epvchain/go-epvchain/swarm/...
