// Slice: estrutura dinâmica baseada em array.
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println("Slice:", s, "Tamanho:", len(s))
}
