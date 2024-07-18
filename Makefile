BINARY_NAME=rotdetector
CMDS=$(filter-out internal, $(notdir $(wildcard cmd/*)))
	

.PHONY: all test build clean run tidy

all: test build

test:
	go test -v -race -cover ./...

build: $(CMDS)

${CMDS}:
	go build -o bin/${BINARY_NAME} cmd/${BINARY_NAME}/*.go

clean:
	go clean
	rm -f bin/${BINARY_NAME}

run: build
	./bin/${BINARY_NAME}
