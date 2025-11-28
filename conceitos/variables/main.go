// Variáveis: global e local, alteração de valor.
package main

import "fmt"

var global int = 5

func main() {
	local := 3
	fmt.Println("Global:", global, "Local:", local)
	local = 10
	fmt.Println("Local alterado:", local)
}
