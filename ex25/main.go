package main

import "fmt"

func main() {
	var n int
	fmt.Print("Quantidade de valores: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Nada a calcular")
		return
	}
	soma := 0.0
	for i := 0; i < n; i++ {
		var v float64
		fmt.Printf("Valor %d: ", i+1)
		fmt.Scan(&v)
		soma += v
	}
	fmt.Printf("Média: %.2f\n", soma/float64(n))
}
