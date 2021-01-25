CLI_BIN=cli_bin
BOT_BIN=bot_bin

GO_FLAGS=

bot:
	go build $(GO_FLAGS) -o $(BOT_BIN) cmd/bot.go

cli:
	go build $(GO_FLAGS) -o $(CLI_BIN) cmd/cli.go 

format:
	go $(GO_FLAGS) fmt ./...

test:
	go test $(GO_FLAGS) -count=1 ./...

clean:
	rm -f *_bin
