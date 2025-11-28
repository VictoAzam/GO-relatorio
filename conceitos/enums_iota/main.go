// Enums com iota: constantes sequenciais e bitflags simples.
package main

import "fmt"

const (
	Domingo = iota
	Segunda
	Terca
	Quarta
	Quinta
	Sexta
	Sabado
)

const (
	PermLeitura = 1 << iota // 001
	PermEscrita             // 010
	PermExecucao            // 100
)

func main() {
	fmt.Println("Dias:", Domingo, Segunda, Terca, Quarta, Quinta, Sexta, Sabado)
	perms := PermLeitura | PermEscrita
	fmt.Printf("Permissões: leitura=%t escrita=%t execução=%t\n",
		perms&PermLeitura != 0, perms&PermEscrita != 0, perms&PermExecucao != 0)
}
