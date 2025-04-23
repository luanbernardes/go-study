build:
	cd project && go mod tidy
run:
	cd project && air --build.cmd "go build -o bin/api cmd/main.go" --build.bin "./bin/api"

# Docker
build-project:
	cd project && docker compose build
run-project:
	cd project && docker compose up