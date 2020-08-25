all:
	rm -rf build && mkdir -p build/bin
	cd go_impl && go build && mv go_impl ../build/bin/go_benchmark
	clang -O3 -o build/bin/c_benchmark c_impl/montmulbench.c
