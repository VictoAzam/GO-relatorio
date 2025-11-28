package main

import "fmt"

func main() {
	var peso, altura float64
	fmt.Print("Peso(kg) Altura(m): ")
	fmt.Scan(&peso, &altura)
	if peso <= 0 || altura <= 0 {
		fmt.Println("Valores inválidos")
		return
	}
	imc := peso / (altura * altura)
	class := ""
	switch {
	case imc < 18.5:
		class = "Magreza"
	case imc < 25:
		class = "Normal"
	case imc < 30:
		class = "Sobrepeso"
	default:
		class = "Obesidade"
	}
	fmt.Printf("IMC: %.2f (%s)\n", imc, class)
}
