// Laços for: forma clássica e estilo while.
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("i =", i)
	}
	j := 0
	for j < 3 {
		fmt.Println("j =", j)
		j++
	}
}
