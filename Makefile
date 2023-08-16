build:
	@go build -o ubi-init

test:
	@go clean -testcache
	@go test ./...

test-coverage:
	@GRPC_STARTSTOP=false go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
