// Map: chave para valor (dicionário).
package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2}
	m["c"] = 3
	delete(m, "b")
	fmt.Println("Mapa:", m)
}
