PKGS    := $(shell go list ./... | grep -v /vendor/)

fmt:
	go fmt ${PKGS}
.PHONY: fmt
