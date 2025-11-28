// Exercício 1: lê dois inteiros e imprime a soma.
package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Digite dois inteiros separados por espaço: ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Entrada inválida")
		return
	}
	fmt.Printf("Soma: %d\n", a+b)
}
