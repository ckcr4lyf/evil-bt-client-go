.PHONY: linux

make linux:
	go build -i ./bin/main ./cmd/app/main.go