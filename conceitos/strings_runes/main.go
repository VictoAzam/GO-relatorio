// Strings e runes: cada rune pode ter mais de 1 byte.
package main

import "fmt"

func main() {
	s := "Olá"
	for i, r := range s {
		fmt.Printf("Posição=%d Código=%U Caractere=%c\n", i, r, r)
	}
}
