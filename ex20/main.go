package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Palavra a buscar: ")
	palavra, _ := reader.ReadString('\n')
	palavra = strings.TrimSpace(strings.ToLower(palavra))
	fmt.Println("Digite texto (linha única):")
	texto, _ := reader.ReadString('\n')
	texto = strings.ToLower(texto)
	count := 0
	for _, w := range strings.Fields(texto) {
		if w == palavra {
			count++
		}
	}
	fmt.Printf("Ocorrências de '%s': %d\n", palavra, count)
}
