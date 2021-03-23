package main

import (
	"fmt"
	"math/big"
)

func test() {

	a := big.NewInt(1)
	b := big.NewInt(2)
	fmt.Println(a, b, a.Cmp(b))

	a = big.NewInt(2)
	b = big.NewInt(2)
	fmt.Println(a, b, a.Cmp(b))

	a = big.NewInt(3)
	b = big.NewInt(2)
	fmt.Println(a, b, a.Cmp(b))
}
