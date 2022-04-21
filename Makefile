LDFLAGS="-extldflags '-static' $(shell ./build/version.sh)"

fmt: ## Run go fmt against code.
	goimports -w -local github.com/27149cheo .

vet: ## Run go vet against code.
	go vet ./...

.PHONY: build
build:
	go build -ldflags ${LDFLAGS} -o helmtool ./cmd/helmtool/main.go
