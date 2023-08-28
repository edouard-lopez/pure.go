.PHONY: install
install:
	go install gotest.tools/gotestsum@latest
	go install github.com/goreleaser/goreleaser/...@latest

.PHONY: test
test:
	gotestsum \
		--format testname \
	$$(go list ./... | grep -v '/cmd')	

.PHONY: test-watch
test-watch:
	clear
	gotestsum \
		--format testname \
		--watch \
	./...

.PHONY: dev
dev:
	go run cmd/cli.go

.PHONY: run
run: build
	./pure --last-command-status $$?

.PHONY: build
build:
	GOOS=linux \
	GOARCH=amd64 \
	go build \
		-o pure \
		cmd/cli.go
	chmod u+x pure

.PHONY: demo
demo: build
	./pure --last-command-status 0
	./pure --last-command-status 1


.PHONY: test-coverage
test-coverage:
	go test -v -count=1 -race -coverprofile=./coverage-unit.txt -covermode=atomic ./...


.PHONY: test-all
test-all: build
	gotestsum \
		--format testname \
	./...

.PHONY: test-e2e # only the e2e tests
test-e2e: build
	gotestsum \
		--format testname \
	./cmd