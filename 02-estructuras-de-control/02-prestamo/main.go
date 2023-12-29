package main

import "fmt"

func main() {

	prestamo(23, true, 2, 100000)
	prestamo(18, true, 2, 90000)
	prestamo(23, true, 2, 110000)

}

func prestamo(edad int, empleado bool, antiguedad int, sueldo float64) {

	fmt.Printf("Cliente: edad %d, empleado %t, antiguedad %d, sueldo %.2f\n", edad, empleado, antiguedad, sueldo)

	if edad > 22 && empleado && antiguedad > 1 {

		fmt.Println("Prestamo aprobado")

		if sueldo > 100000 {
			fmt.Println("No te cobramos intereses")
		}

		return
	}

	fmt.Println("Prestamo denegado")

}
