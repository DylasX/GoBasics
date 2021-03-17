package main

import "fmt"

func useFmt() {
	helloMessage := "Hello"
	worldMessage := "world"

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Hola"
	contador := 500
	// Con valores seguros
	fmt.Printf("%s tiene más de %d contador\n", nombre, contador)
	// Con valores inseguros
	fmt.Printf("%v tiene más de %v contador\n", nombre, contador)

	// Sprintf
	message := fmt.Sprintf("%v tiene más de %v contador\n", nombre, contador)
	fmt.Println(message)

	// Tipo de datos:
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("contador: %T\n", contador)
}
