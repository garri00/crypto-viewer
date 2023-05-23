hello:
	echo "Hello"

lint:
	golangci-lint run

test :
	go test ./...

run :
	go run ./bin/main.go

all : hello lint test run