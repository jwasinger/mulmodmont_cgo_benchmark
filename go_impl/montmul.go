package main

/*
#include "../include/montmul.h"
*/
import "C"

import (
    "fmt"
    "unsafe"
    "time"
)

func mulmodmont(out *[6]C.ulong, x *[6]C.ulong, y *[6]C.ulong) {
	C.mulmodmont((*C.ulong)(unsafe.Pointer(out)), (*C.ulong)(unsafe.Pointer(x)), (*C.ulong)(unsafe.Pointer(y)))
}

func main() {
	x := [6]C.ulong{0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff}
	y := [6]C.ulong{0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff,0xffffffffffffffff}

	start := time.Now()
	for i := 0; i < 100000; i++ {
		mulmodmont(&x, &x, &y)
	}
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println("took: ", elapsed.String())

	fmt.Print("result: ")
	for i := 0; i < 6; i++ {
		fmt.Print(x[i], ", ")
	}
	fmt.Print("\n")
}
