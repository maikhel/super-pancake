build: install-dependencies
	go build -o super-pancake

install-dependencies:
	dep ensure

run: install-dependencies
	go build -o super-pancake && ./super-pancake

test: install-dependencies
	go test

test-with-report: install-dependencies
	go test -coverprofile=coverage.txt -covermode=atomic

test-html-report: test-with-report
	go tool cover -html=coverage.txt
