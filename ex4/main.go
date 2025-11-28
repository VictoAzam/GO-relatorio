package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um inteiro: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Entrada inválida")
		return
	}
	paridade := "ímpar"
	if n%2 == 0 {
		paridade = "par"
	}
	var sinal string
	switch {
	case n > 0:
		sinal = "positivo"
	case n < 0:
		sinal = "negativo"
	default:
		sinal = "zero"
	}
	fmt.Printf("Número %d é %s e %s\n", n, paridade, sinal)
}
