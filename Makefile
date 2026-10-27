TOOLS := corruptor

REPO_corruptor := https://github.com/BradBenJoya/corruptor

.PHONY: install build run docker

install:
	@for t in $(TOOLS); do rm -rf tools/$$t; done
	git clone --depth 1 $(REPO_corruptor) tools/corruptor

build:
	@for t in $(TOOLS); do \
		cmake -S tools/$$t -B bin/$$t -G Ninja -DCMAKE_BUILD_TYPE=Release && \
		cmake --build bin/$$t || exit 1; \
	done
	mkdir -p bin/tinywebutils
	go build -ldflags="-s -w" -o bin/tinywebutils/server ./src/backend

run: build
	./bin/tinywebutils/server

docker:
	docker build -t tinywebutils .
	-docker rm -f tinywebutils
	docker run -d -p 8000:8000 --name tinywebutils tinywebutils