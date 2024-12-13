test:
	go test
build:
	go build main.go
dockerbuild:
	docker build .
lint:
	golangci-lint run
