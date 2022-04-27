BINARY_NAME=ucs

build:
	go build -o bin/${BINARY_NAME} .

test:
	go test -v ./...

clean:
	go clean
	rm -rf bin