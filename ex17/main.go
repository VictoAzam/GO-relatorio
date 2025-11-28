package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	segredo := rand.Intn(100) + 1
	fmt.Println("Adivinhe o número de 1 a 100")
	for tent := 1; ; tent++ {
		var palpite int
		fmt.Print("Palpite: ")
		fmt.Scan(&palpite)
		if palpite < segredo {
			fmt.Println("Maior")
		} else if palpite > segredo {
			fmt.Println("Menor")
		} else {
			fmt.Printf("Acertou em %d tentativas!\n", tent)
			break
		}
	}
}
