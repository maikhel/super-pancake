build: install-dependencies
	go build -o super-pancake

install-dependencies:
	dep ensure

local-run: install-dependencies
	go build -o super-pancake && ./super-pancake

test: install-dependencies
	godotenv -f .env.test go test  ./...

test-with-report: install-dependencies
	godotenv -f .env.test go test  ./... -coverprofile=coverage.txt -covermode=atomic

test-html-report: test-with-report
	go tool cover -html=coverage.txt
