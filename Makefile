test:
	@go test -race .
test-coverage:
	@go test -cover -coverprofile=coverage.out && go tool cover -html=coverage.out