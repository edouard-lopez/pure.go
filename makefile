# Install development tools (gotestsum for testing, goreleaser for releases)
.PHONY: install
install:
	go install gotest.tools/gotestsum@v1.11.0 
	go install github.com/goreleaser/goreleaser/...@v1.21.2

# Run unit tests (excludes e2e tests in cmd/)
.PHONY: test
test:
	gotestsum \
		--format testname \
	$$(go list ./... | grep -v '/cmd')	

# Run tests in watch mode, re-running on file changes
.PHONY: test-watch
test-watch:
	clear
	gotestsum \
		--format testname \
		--watch \
	./...

# Run the CLI directly without building
.PHONY: dev
dev:
	go run cmd/cli.go

# Build and run the prompt with current shell exit status
.PHONY: run
run: build
	./pure --last-command-status $$?

# Build the Linux amd64 binary
.PHONY: build
build:
	GOOS=linux \
	GOARCH=amd64 \
	go build \
		-o pure \
		cmd/cli.go
	chmod u+x pure

# Demo the prompt with success (0) and failure (1) exit codes
.PHONY: demo
demo: build
	./pure --last-command-status 0
	./pure --last-command-status 1


# Run tests with race detection and generate coverage report
.PHONY: test-coverage
test-coverage:
	go test -v -count=1 -race -coverprofile=./coverage-unit.txt -covermode=atomic ./...


# Run all tests including e2e tests
.PHONY: test-all
test-all: build
	gotestsum \
		--format testname \
	./...

# Run only end-to-end tests in cmd/
.PHONY: test-e2e
test-e2e: build
	gotestsum \
		--format testname \
	./cmd