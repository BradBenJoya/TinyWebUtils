.PHONY: default
default:
	go build -o bin/tinywebutils/server ./src/backend
	./bin/tinywebutils/server