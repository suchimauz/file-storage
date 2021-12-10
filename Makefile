HTTP_HOST=127.0.0.1
HTTP_PORT=8001
APP_ENV=local
STORAGE_PREFIX=storage

.EXPORT_ALL_VARIABLES:

run:
	go run cmd/app/main.go
build:
	go build -o .bin/ cmd/app/main.go
mod:
	go mod tidy
