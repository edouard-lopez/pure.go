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

.PHONY: run
run:
	go run cmd/cli.go