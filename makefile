build: lint
	go build -o bin/api ./cmd/api

run: format lint test build
	./bin/api

lint:
	golangci-lint run ./...

format:
	go fmt ./...

test:
	gotestsum --format pkgname ./tests/...