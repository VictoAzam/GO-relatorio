// Ponteiros: passagem por referência.
package main

import "fmt"

func inc(p *int) { *p++ }
func main()      { x := 5; inc(&x); fmt.Println("Valor de x:", x) }
