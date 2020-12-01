install-packr:
	@go get -u github.com/gobuffalo/packr/packr

test-coverage-report:
	@go test -covermode=count -coverprofile=/tmp/count.out -v ./...
	@go tool cover -html=/tmp/count.out
