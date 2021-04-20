.PHONY: linux

make linux:
	go build -o ./bin/main ./cmd/app/main.go