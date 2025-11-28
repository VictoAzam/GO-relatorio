// Closure: função interna acessa variável externa.
package main

import "fmt"

func gerador() func() int { v := 0; return func() int { v++; return v } }
func main()               { next := gerador(); fmt.Println(next(), next(), next()) }
