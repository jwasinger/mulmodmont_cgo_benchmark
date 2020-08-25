#include "../include/montmul.h"
#include<stdio.h>
#include<time.h>

int main(int argc, char **argv) {
	struct timespec start, end;
	uint64_t x[6] = {0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff};
	uint64_t y[6] = {0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff};

	clock_gettime(CLOCK_REALTIME, &start);
	for (int i = 0; i < 10000000; i++) {
		mulmodmont(&x,&x,&y);
	}
	clock_gettime(CLOCK_REALTIME, &end);
	  double took = ( end.tv_sec - start.tv_sec )
                 + ( end.tv_nsec - start.tv_nsec )
                 / 1E9;
	printf( "took: %lfms\n", took * 1000.0 );
	printf("result: ");
	for (int i = 0; i < 6; i++) {
		printf("%lu, ", x[i]);
	}
	printf("\n");
}
