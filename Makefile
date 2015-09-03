osx:
	@GOOS=darwin GOARCH=amd64 go build

linux32:
	@GOOS=linux GOARCH=386 go build

linux64:
	@GOOS=linux GOARCH=amd64 go build

test:
	@go test ./...
