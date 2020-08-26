package main

import (
    "fmt"
    "time"
    "./bls12_381"
)

func main() {
    var x, y bls12_381.Element
    /*
        0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff == 39402006196394479212279040100143613805079739270465446667948293404245721771497210611414266254884915640806627990306815
    */
    x.SetString("39402006196394479212279040100143613805079739270465446667948293404245721771497210611414266254884915640806627990306815")
    y.SetString("39402006196394479212279040100143613805079739270465446667948293404245721771497210611414266254884915640806627990306815")

	start := time.Now()
    for i := 0; i < 10000000; i++ {
        x.Mul(&x, &y)
    }
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println("took: ", elapsed.String())
    fmt.Println("result: ", x.String())
}
