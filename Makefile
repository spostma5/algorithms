build:
	@go build -o bin/algo cmd/main.go

run: build
	@go run cmd/main.go

test:
	@go test -v ./...

bench:
	@go test -bench=. ./... -benchmem
