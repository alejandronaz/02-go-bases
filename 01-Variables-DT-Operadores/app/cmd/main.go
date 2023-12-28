package main

import (
	// nombreDelModulo/pathDelPaquete
	"clase1/app/internals/calculator"
	"fmt"
)

func main() {
	fmt.Println("Suma de 2 + 3 =", calculator.Suma(2, 3))
	fmt.Println("Nombre =", calculator.Nombre)
}
