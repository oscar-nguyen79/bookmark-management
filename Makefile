.PHONY: run swagger test dev-run

run:
	go run cmd/api/main.go

swagger:
	swag init -g cmd/api/main.go

test:
	go test ./...

dev-run: swagger run