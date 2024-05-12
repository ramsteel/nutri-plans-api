.PHONY: test coverage
test:
	go test ./usecases -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out

.PHONY: run-container
run-container:
	docker compose up -d --build

.PHONY: stop-container
stop-container:
	docker compose down