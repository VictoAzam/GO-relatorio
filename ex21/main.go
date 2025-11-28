package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int64
	fmt.Print("Número para fatorial: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Negativo não permitido")
		return
	}
	res := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		res.Mul(res, big.NewInt(i))
	}
	fmt.Printf("%d! = %s\n", n, res.String())
}
