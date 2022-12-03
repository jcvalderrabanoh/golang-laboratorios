package main

import "fmt"

func pruebaMultipleReturn() (a int, b int, c int, d int, e int) {
	base := 100
	altura := 150
	diametro := 100

	return base, altura, base * altura, diametro, (base * altura * base)
}

func main() {

	// Declaración de constantes
	const a = 3
	const b int = 3
	const c int = 3.32423 - .32423
	const d string = "3.32423"
	const area = a * b
	// genera la asignación de la variable
	division := a / b

	// Variables sin inialización

	var e int
	var f bool
	var g string
	var h float32
	var i float64
	var j byte

	fmt.Println("Constantes")
	fmt.Println(a, b, c, d, area, division)

	fmt.Println("Variables declaradas sin inializar")
	fmt.Println(e, f, g, h, i, j)

	fmt.Println(pruebaMultipleReturn())
	z, x, y, w, u := pruebaMultipleReturn()
	fmt.Println(z, x, y, w, u)
}
