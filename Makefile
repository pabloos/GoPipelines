
ifndef VERBOSE
.SILENT:
endif

TEST_TIMEOUT = 30s
TEST = go test -v -race -timeout ${TEST_TIMEOUT}

test: clean-test-cache
	${TEST}

trace: clean-test-cache
	${TEST} 2> trace.out
	go tool trace trace.out

bench:
	go test -bench=.

clean-test-cache:
	echo "cleaning test cache..."
	go clean -testcache
