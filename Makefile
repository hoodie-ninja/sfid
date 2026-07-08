.PHONY: all
all: test

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: lint
lint:
	$(if $(shell go fix -diff ./...),$(error go fixes required; run 'go fix ./...'))
	$(if $(shell which golangci-lint),,$(error golangci-lint required for make lint))
	golangci-lint version
	golangci-lint run

.PHONY: godoc
godoc:
	$(if $(shell which pkgsite),,$(error pkgsite required for make godoc))
	pkgsite -open

.PHONY: update
update:
	rm -rf vendor/
	go get -u ./...
	go fix ./...
	go mod tidy
	go mod vendor
