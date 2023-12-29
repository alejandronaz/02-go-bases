package main

import "fmt"

func main() {
	letrasPalabra("Hola")
	letrasPalabra("Golang")
}

func letrasPalabra(palabra string) {

	fmt.Println("Letras de la palabra: ", palabra)

	for _, palabra := range palabra {
		fmt.Println(string(palabra))
	}

}
