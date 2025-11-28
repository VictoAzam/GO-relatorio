// Métodos: função associada ao tipo.
package main

import "fmt"

type Contador struct{ v int }

func (c *Contador) Inc() { c.v++ }
func main()              { c := &Contador{}; c.Inc(); c.Inc(); fmt.Println("Contador:", c.v) }
