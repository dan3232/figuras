package figuras

import "fmt"

type Operaciones interface {
	area() float64
	perimetro() float64
}



func Medidas(o Operaciones) {
	fmt.Println("-----------------------------")
	fmt.Println("Medida:",o)
	fmt.Println("Area:",o.area())
	fmt.Println("Perimetro:",o.perimetro())
	fmt.Println("-----------------------------")
}