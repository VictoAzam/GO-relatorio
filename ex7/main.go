package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Print("Digite uma sequência de caracteres: ")
	fmt.Scan(&s)
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	fmt.Println("Ordenado:", string(runes))
}
