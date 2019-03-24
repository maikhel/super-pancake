build: install-dependencies
	env GOOS=linux go build -ldflags="-s -w" -o main main.go
	mkdir -p bin
	mv main bin/

install-dependencies:
	dep ensure

local-run:
	go run main.go

test:
	godotenv -f .env go test  ./...
