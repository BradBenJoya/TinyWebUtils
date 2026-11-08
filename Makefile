REPO   := https://github.com/BradBenJoya/corruptor
PORT   ?= 8000
IMAGE  ?= tinywebutils
NAME   ?= tinywebutils

.PHONY: install build run stop docker docker-stop docker-logs docker-rm clean

install:
	rm -rf tools/corruptor
	git clone --depth 1 $(REPO) tools/corruptor

build:
	cmake -S tools/corruptor -B bin/corruptor -G Ninja -DCMAKE_BUILD_TYPE=Release
	cmake --build bin/corruptor
	mkdir -p bin/tinywebutils
	go build -ldflags="-s -w" -o bin/tinywebutils/server$(EXE) ./src/backend

run: build
	$(SERVER)

clean:
	rm -rf bin

docker:
	docker build -t $(IMAGE) .
	-docker rm -f $(NAME)
	docker run -d -p $(PORT):8000 --name $(NAME) $(IMAGE)
	@echo http://localhost:$(PORT)

stop docker-stop:
	docker stop $(NAME)

docker-rm:
	-docker rm -f $(NAME)

docker-logs:
	docker logs -f $(NAME)