package main

import (
    "fmt"
    "time"
    "./bls12_381"
)

func main() {
    var x, y bls12_381.Element
    x.SetRandom()
    y.SetRandom()

	start := time.Now()

    for i := 0; i < 10000000; i++ {
        x.Mul(&x, &y)
    }

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println("took: ", elapsed.String())
    fmt.Println("result: ", x.String())
}
