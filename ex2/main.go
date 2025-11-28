package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Digite dois inteiros (dividendo divisor): ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Entrada inválida")
		return
	}
	if b == 0 {
		fmt.Println("Divisor não pode ser zero")
		return
	}
	quociente := float64(a) / float64(b)
	resto := a % b
	fmt.Printf("Quociente: %.2f Resto: %d\n", quociente, resto)
}
