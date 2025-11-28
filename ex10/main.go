// Exercício 10: resolve a Torre de Hanoi usando recursão.
package main

import "fmt"

func hanoi(n int, origem, auxiliar, destino string) {
	if n == 0 {
		return
	}
	hanoi(n-1, origem, destino, auxiliar)
	fmt.Printf("Mover disco %d de %s para %s\n", n, origem, destino)
	hanoi(n-1, auxiliar, origem, destino)
}

func main() {
	var n int
	fmt.Print("Quantidade de discos: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Valor inválido")
		return
	}
	hanoi(n, "A", "B", "C")
}
