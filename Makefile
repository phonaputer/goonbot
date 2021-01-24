CLI_BIN=cli_bin

GO_FLAGS=

cli:
	go build $(GO_FLAGS) -o $(CLI_BIN) cmd/cli.go 

format:
	go $(GO_FLAGS) fmt ./...

test:
	go test $(GO_FLAGS) -count=1 ./...

clean:
	rm -f *_bin
