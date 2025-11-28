package main

import "fmt"

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mmc(a, b int) int { return a / mdc(a, b) * b }

func main() {
	var a, b int
	fmt.Print("Dois inteiros: ")
	fmt.Scan(&a, &b)
	if a == 0 || b == 0 {
		fmt.Println("Zero não permitido")
		return
	}
	fmt.Printf("MMC(%d,%d) = %d\n", a, b, mmc(a, b))
}
