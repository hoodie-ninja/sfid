.PHONY: all
all: test

.PHONY: test
test:
	go test -race -cover ./...
	@test -n $(shell go fix -diff ./...) || (echo "go fixes required" && exit 1)

.PHONY: lint
lint:
	@which golangci-lint > /dev/null || echo "golangci-lint required for `make lint`"
	@golangci-lint version
	@golangci-lint run

.PHONY: godoc
godoc:
	@which pkgsite > /dev/null || echo "pkgsite required for `make godoc`"
	pkgsite -open

.PHONY: update
update:
	rm -rf vendor/
	go get -u ./...
	go fix ./...
	go mod tidy
	go mod vendor
