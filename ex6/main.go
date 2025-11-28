package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Print("Quantidade de números: ")
	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		fmt.Println("Quantidade inválida")
		return
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)
	fmt.Println("Ordenado:", nums)
}
