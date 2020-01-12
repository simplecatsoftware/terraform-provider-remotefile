.PHONY: install
install:
	go get -v -t -d ./...

.PHONY: build
build:
	GOOS=windows GOARCH=amd64 go build -o terraform-provider-remotefile-windows-64bit
	GOOS=linux GOARCH=amd64 go build -o terraform-provider-remotefile-linux-64bit
	GOOS=darwin GOARCH=amd64 go build -o terraform-provider-remotefile-darwin-64bit

.PHONY: test
test:
	TF_ACC=true go test -v ./...

format:
	go fmt ./...
