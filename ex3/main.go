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
	fmt.Printf("Antecessor: %d Sucessor: %d\n", n-1, n+1)
}
