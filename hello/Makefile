
pkgs = $(shell go list ./... )

.PHONY: test_ci
test:
	go test -race $(pkgs)


.PHONY: install
install:
	go install ./cmd/hello

.PHONY: download
download:
	go mod download