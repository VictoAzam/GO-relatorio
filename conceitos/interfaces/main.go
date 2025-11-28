// Interfaces: definem um conjunto de métodos que um tipo deve implementar.
package main

import "fmt"

type Formato interface {
	Area() float64
}

type Retangulo struct{ L, A float64 }
func (r Retangulo) Area() float64 { return r.L * r.A }

type Circulo struct{ R float64 }
func (c Circulo) Area() float64 { return 3.14159 * c.R * c.R }

func imprimeArea(f Formato) { fmt.Printf("Área: %.2f\n", f.Area()) }

func main() {
	imprimeArea(Retangulo{L: 3, A: 4})
	imprimeArea(Circulo{R: 2.5})
}
