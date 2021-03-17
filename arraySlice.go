package main

import "fmt"

func arraySlices() {
	//array dimensionado
	var array [4]int
	array[0] = 1
	array[1] = 2
	//len = longitud, cap = capacidad array
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])  //imprima hasta 3 elementos
	fmt.Println(slice[2:4]) //del 2 al 4
	fmt.Println(slice[4:])  //del 4 en adelante el que va despues de los dos puntos ese numero es exclusivo el primero inclusivo quiere decir un intervalo cerrado, abierto

	//Append
	slice = append(slice, 100)

	//Add slice to slice
	slice = append(slice, slice...) // iterando ... o sea no hay una lista dentro de una lista si no que se recorre los elementos

	fmt.Println(slice)
}
