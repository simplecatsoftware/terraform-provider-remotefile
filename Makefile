install:
	go get

build:
	go build -o terraform-provider-remotefiles

.PHONY: test
test:
	TF_ACC=true go test -v ./...

format:
	go fmt terraform-provider-remotefiles/...
