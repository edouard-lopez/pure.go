.PHONY: test
test:
	gotestsum \
		--format testname \
	./...

.PHONY: test-watch
test-watch:
	gotestsum \
		--format testname \
		--watch \
	./...