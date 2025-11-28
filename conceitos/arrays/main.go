// Array: tamanho fixo definido na declaração.
package main

import "fmt"

func main() {
	var a [3]int
	a[0], a[1], a[2] = 1, 2, 3
	fmt.Println("Array:", a, "Tamanho:", len(a))
}
