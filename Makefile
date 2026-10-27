local:
	mkdir -p bin/corruptor bin/tinywebutils
	cmake -S tools/corruptor -B bin/corruptor -G Ninja -DCMAKE_BUILD_TYPE=Release
	cmake --build bin/corruptor
	go build -ldflags="-s -w" -o bin/tinywebutils/server ./src/backend
	./bin/tinywebutils/server
docker:
	docker build -t tinywebutils .
	docker run -d -p 8000:8000 --name tinywebutils tinywebutils