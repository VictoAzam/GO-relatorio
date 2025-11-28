package main

import "fmt"

func ehPrimo(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Digite um inteiro: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Entrada inválida")
		return
	}
	if ehPrimo(n) {
		fmt.Printf("%d é primo\n", n)
	} else {
		fmt.Printf("%d não é primo\n", n)
	}
}
