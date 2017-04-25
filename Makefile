
all: test
.PHONY: all

test:
	@go test -cover ./...
.PHONY: test
