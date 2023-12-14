run-example:
	go run example/*.go

test-all:
	go test -v ./

test-cover:
	go test -v ./ --cover

.PHONY: run-example test-all test-cover