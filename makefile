.PHONY: test
test:
	gotestsum \
		--format testname \
	./...

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
