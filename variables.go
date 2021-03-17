package main

import "fmt"

func variables() {
	//tipos

	//int -> int8, int16, int32, int64
	//uint: enteros positivos -> igual a int
	//float32 y float64
	//string
	//bool
	//complex: numeros imaginarios complex64: 32bits, complex128: 64bits -> 1i

	//Declaracion de constantes
	const pi float64 = 3.14 //64 bits
	const pi2 = 3.14

	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	//Declaracion de variables
	base := 12          //Primera forma se crea y se inicializa valor
	var altura int = 14 //Segunda forma se crea y se especifica explicitamente su tipo
	var area int        //Go no compila si las variables no son usadas

	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna vaalores a variables vacías
	var a int     // valor defecto 0
	var b float64 // valor defecto 0
	var c string  // valor defecto ""
	var d bool    // valor defecto false

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)
}
