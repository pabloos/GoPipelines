
ifndef VERBOSE
.SILENT:
endif

# PROGRAM_NAME=pipeline

clean-test-cache:
	echo "cleaning test cache..."
	go clean -testcache

test: clean-test-cache
	go test -v -race -timeout 30s

trace: clean-test-cache
	go test -v -race -timeout 30s 2> trace.out
	go tool trace trace.out

# compile:
# 	go build -race -o ${PROGRAM_NAME}

# run:
# 	./${PROGRAM_NAME}

# clean:
# 	rm ${PROGRAM_NAME}

bench:
	go test -bench=.
