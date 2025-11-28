package main

import "fmt"

func pergunta(p string) bool {
	var r string
	fmt.Print(p, " (s/n): ")
	fmt.Scan(&r)
	return r == "s" || r == "S"
}

func main() {
	fmt.Println("Árvore de decisão: identificar animal")
	if pergunta("É um mamífero?") {
		if pergunta("Voa?") {
			fmt.Println("Possível: Morcego")
		} else if pergunta("É doméstico?") {
			fmt.Println("Possível: Cachorro/Gato")
		} else {
			fmt.Println("Outro mamífero")
		}
	} else {
		if pergunta("Tem penas?") {
			fmt.Println("É uma ave")
		} else if pergunta("Vive na água?") {
			fmt.Println("Pode ser peixe")
		} else {
			fmt.Println("Outro tipo de animal")
		}
	}
}
