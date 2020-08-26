package main

import (
    "fmt"
    "time"
    "./bls12_381"
)

func main() {
    var x, y bls12_381.Element
    x.SetString("2")
    y.SetString("2")

    // TODO assert mont_form(2) * mont_form(2) == 1514052131932888505822357196874193114600527104240479143842906308145652716846165732392247483508051665748635331395571
	start := time.Now()

    for i := 0; i < 10000000; i++ {
        x.Mul(&x, &y)
    }

    x.Mul(&x, &y)
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println("took: ", elapsed.String())
    fmt.Println("result: ", x.String())
}
