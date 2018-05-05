GOBIN = $(shell pwd)/bin/bin
GO ?= latest

gepv:
	bin/build.sh go run bin/ep.go install ./command/gepv
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gepv\" to launch gepv."

clean:
	rm -fr bin/_workspace/pkg/ $(GOBIN)/*
