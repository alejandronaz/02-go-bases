package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("Empleados: ", employees)

	fmt.Println("Edad de Benjamin: ", employees["Benjamin"])

	var mayores21 int
	for _, value := range employees {
		if value > 21 {
			mayores21++
		}
	}
	fmt.Println("Mayores de 21: ", mayores21)

	// New employee
	employees["Federico"] = 25
	fmt.Println("Federico nuevo empleado: ", employees)

	// Delete employee
	delete(employees, "Pedro")
	fmt.Println("Eliminar a Pedro: ", employees)
}
