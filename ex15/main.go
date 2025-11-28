package main

import "fmt"

func main() {
	var base, altura float64
	fmt.Print("Base e altura: ")
	fmt.Scan(&base, &altura)
	fmt.Printf("Área: %.2f\n", base*altura)
}
