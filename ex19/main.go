package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Digite texto (linha única):")
	reader := bufio.NewReader(os.Stdin)
	linha, _ := reader.ReadString('\n')
	vogais := "aeiouáéíóúâêîôûãõ" // simplificado
	vCount, cCount := 0, 0
	for _, r := range []rune(linha) {
		if unicode.IsLetter(r) {
			rLower := unicode.ToLower(r)
			if strings.ContainsRune(vogais, rLower) {
				vCount++
			} else {
				cCount++
			}
		}
	}
	fmt.Printf("Vogais: %d Consoantes: %d\n", vCount, cCount)
}
