package main

import "fmt"

//defer, continue, break

func workingWithThatReserved() {
	//Defer Print limitar solo 1 defer por funcion
	defer fmt.Println("Probando defer #1") // Este se imprimira despues util para romper la conexion de una BD
	fmt.Println("Probando defer #2")

	var i uint8 = 0
	for i < 10 {
		fmt.Println(i)
		i++

		// Continue
		if i == 2 {
			fmt.Println("¡El número que sigue es par!")
			continue
		}

		if i == 9 {
			fmt.Println("Esto no va a continuar va a romperse")
			break
		}

	}
}
