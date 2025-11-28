// Função variádica: aceita quantidade variável de argumentos.
package main

import "fmt"

func soma(nums ...int) int {
	t := 0
	for _, n := range nums {
		t += n
	}
	return t
}
func main() { fmt.Println("Soma:", soma(1, 2, 3, 4)) }
