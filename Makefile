# To try different version of Go
GO := go

test:
	go test -mod=readonly -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt ./...
	go tool cover -html=coverage.txt -o coverage.html

lint:
	golangci-lint run ./...
	go fmt ./...