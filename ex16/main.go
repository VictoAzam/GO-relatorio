package main

import "fmt"

func main() {
	var valor float64
	var origem, destino string
	fmt.Print("Valor e unidade origem (C/F/K) e destino (C/F/K): ")
	fmt.Scan(&valor, &origem, &destino)
	origem = string(origem[0])
	destino = string(destino[0])
	c := valor
	switch origem {
	case "F":
		c = (valor - 32) * 5 / 9
	case "K":
		c = valor - 273.15
	case "C":
	default:
		fmt.Println("Origem inválida")
		return
	}
	var out float64
	switch destino {
	case "C":
		out = c
	case "F":
		out = c*9/5 + 32
	case "K":
		out = c + 273.15
	default:
		fmt.Println("Destino inválido")
		return
	}
	fmt.Printf("Resultado: %.2f %s\n", out, destino)
}
