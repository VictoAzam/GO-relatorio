package main

import "fmt"

func diaSemana(d, m, a int) string {
	if m < 3 {
		m += 12
		a -= 1
	}
	K := a % 100
	J := a / 100
	h := (d + (13*(m+1))/5 + K + K/4 + J/4 + 5*J) % 7
	dias := []string{"Sábado", "Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta"}
	return dias[h]
}

func main() {
	var d, m, a int
	fmt.Print("Digite dia mes ano (dd mm aaaa): ")
	fmt.Scan(&d, &m, &a)
	fmt.Println("Dia da semana:", diaSemana(d, m, a))
}
