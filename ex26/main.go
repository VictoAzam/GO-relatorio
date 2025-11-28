// Exercício 26: algoritmo genético simples para maximizar a função fitness.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fitness(x int) int { return -(x-7)*(x-7) + 100 }

func main() {
	rand.Seed(time.Now().UnixNano())
	popTam := 20
	geracoes := 50
	pop := make([]int, popTam)
	for i := range pop {
		pop[i] = rand.Intn(101)
	}
	melhor := pop[0]
	for g := 0; g < geracoes; g++ {
		for _, x := range pop {
			if fitness(x) > fitness(melhor) {
				melhor = x
			}
		}
		nova := make([]int, 0, popTam)
		for len(nova) < popTam {
			a := pop[rand.Intn(popTam)]
			b := pop[rand.Intn(popTam)]
			vencedor := a
			if fitness(b) > fitness(a) {
				vencedor = b
			}
			if rand.Float64() < 0.3 {
				vencedor += rand.Intn(7) - 3
				if vencedor < 0 {
					vencedor = 0
				}
				if vencedor > 100 {
					vencedor = 100
				}
			}
			nova = append(nova, vencedor)
		}
		pop = nova
	}
	fmt.Printf("Melhor solução: x=%d fitness=%d\n", melhor, fitness(melhor))
}
