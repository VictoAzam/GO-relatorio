// Recursão: função chama a si mesma.
package main

import "fmt"

func fatorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fatorial(n-1)
}
func main() { fmt.Println("Fatorial de 5:", fatorial(5)) }
