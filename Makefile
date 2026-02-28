test:
	go test ./...

test-short:
	go test -short ./...

test-integration:
	BEADS_INTEGRATION=1 go test ./internal/beads -v

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
