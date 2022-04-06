lint:
	golangci-lint run

run:
	go run ./...

test:
	go test ./...

fmt:
	go fmt ./...

dependencies:
	go mod vendor -v
	go mod tidy
	go mod verify
