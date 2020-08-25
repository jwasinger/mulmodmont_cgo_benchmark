all:
	rm -rf build && mkdir -p build/bin
	cd cgo_wrapper && go build && mv cgo_wrapper ../build/bin/go_benchmark
	clang -O3 -o build/bin/c_benchmark c/montmulbench.c
