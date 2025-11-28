package main

import (
	"fmt"
	"strings"
)

func normaliza(s string) string {
	s = strings.ToLower(s)
	rs := strings.Builder{}
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			rs.WriteRune(r)
		}
	}
	return rs.String()
}

func ehPalindromo(s string) bool {
	s = normaliza(s)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	var entrada string
	fmt.Print("Digite palavra/frase curta: ")
	fmt.Scanln(&entrada)
	if ehPalindromo(entrada) {
		fmt.Println("É palíndromo")
	} else {
		fmt.Println("Não é palíndromo")
	}
}
