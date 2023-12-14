run:
	go run example/*.go

test:
	go test -v ./

test-cover:
	go test -v ./ --cover

.PHONY: run test test-cover