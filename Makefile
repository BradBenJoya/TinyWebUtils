default:
	cmake -S tools/corruptor -B bin/corruptor -G Ninja  -DCMAKE_BUILD_TYPE=Release
	cmake --build bin/corruptor
	go build -o bin/tinywebutils src/backend/main.go 
	./bin/tinywebutilsbin/tinywebutils
