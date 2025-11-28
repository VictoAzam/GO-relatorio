// Função com múltiplos retornos (quociente e resto).
package main

import "fmt"

func div(a, b int) (int, int) { return a / b, a % b }
func main()                   { q, r := div(10, 3); fmt.Println("Quociente:", q, "Resto:", r) }
