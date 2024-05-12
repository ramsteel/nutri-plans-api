.PHONY: test coverage
test:
	go test ./usecases -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out