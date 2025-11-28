// range: percorre slice e map.
package main

import "fmt"

func main() {
	for i, v := range []string{"a", "b", "c"} {
		fmt.Println("Índice:", i, "Valor:", v)
	}
	m := map[string]int{"x": 1, "y": 2}
	for k, v := range m {
		fmt.Println("Chave:", k, "Valor:", v)
	}
}
