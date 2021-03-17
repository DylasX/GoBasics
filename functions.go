package main

import "fmt"

func calculateRectangleArea(b, h int) {
	rectangleArea := b * h
	fmt.Println("El área del rectángulo es:", rectangleArea)
}

func calculateTrapezeArea(b1, b2, h1 float64) {
	trapezeArea := (b1 + b2) * h1 / 2
	fmt.Println("El área del trapecio es:", trapezeArea)
}

func calculateCircleArea(r float64) {
	const pi float64 = 3.14
	circleArea := pi * (r * r)
	fmt.Println("El área del circulo es:", circleArea)
}

func doubleReturn(a int) (c int, d int) {
	return a, a * 2

}

func closureExpose() {
	var num int = 0

	//esto es un closure
	incrementar := func() int {
		num++
		return num
	}

	fmt.Println("Num: ", incrementar()) // 1
}

// Funcion que retorna una funcion anonima
func increment() func() int {
	var num int = 0

	return func() int {
		num++
		return num
	}
}
