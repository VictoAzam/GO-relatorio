// Struct: agrupamento de campos nomeados.
package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() { p := Pessoa{"Ana", 30}; fmt.Println("Pessoa:", p.Nome, p.Idade) }
