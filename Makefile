run:
	go run cmd/*.go

test:
	go test -v ./

.PHONY: run test