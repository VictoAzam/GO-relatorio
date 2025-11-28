package main

import "fmt"

func main() {
	var n int
	fmt.Print("Quantidade de números: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Nada a processar")
		return
	}
	freq := map[int]int{}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		freq[nums[i]]++
	}
	moda := nums[0]
	max := freq[moda]
	for v, f := range freq {
		if f > max {
			moda = v
			max = f
		}
	}
	fmt.Printf("Moda: %d (freq %d)\n", moda, max)
}
