
ifndef VERBOSE
.SILENT:
endif

TEST_TIMEOUT = 30s
COVER_PROFILE = coverage.out

TEST = go test ./... -v -race -timeout ${TEST_TIMEOUT}
COVER = go tool cover

test: clean-test-cache
	${TEST}

coverage: clean-test-cache
	${TEST} -covermode=atomic -coverprofile=${COVER_PROFILE}
	${COVER} -func ${COVER_PROFILE}
	${COVER} -html=${COVER_PROFILE}
	rm ${COVER_PROFILE}

trace: clean-test-cache
	${TEST} 2> trace.out
	go tool trace trace.out

bench:
	go test -bench=.

clean-test-cache:
	echo "cleaning test cache..."
	go clean -testcache
