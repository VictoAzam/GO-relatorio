// Switch: seleção baseada em valor.
package main

import "fmt"

func main() {
	dia := 3
	switch dia {
	case 1:
		fmt.Println("Segunda-feira")
	case 2:
		fmt.Println("Terça-feira")
	case 3, 4:
		fmt.Println("Meio da semana")
	default:
		fmt.Println("Outro dia")
	}
}
