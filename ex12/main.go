package main

import "fmt"

func main() {
	var a, b string
	fmt.Print("Digite dois valores: ")
	fmt.Scan(&a, &b)
	fmt.Println(a == b)
}
