package main

import (
	myPackage "basic/mypackage"
	"fmt"
)

func main() {
	//variables
	variables()

	//Arithmethic operator
	fmt.Println("Calling operators")
	fmt.Println("Result add 1 + 2 : ", add(1, 2))
	fmt.Println("Result subtract 2 - 1 : ", subtract(2, 1))
	fmt.Println("Result division 2 / 1 : ", division(2, 1))
	fmt.Println("Result multiply 2 * 1 : ", multiply(2, 1))
	useFmt()
	calculateRectangleArea(10, 10)
	calculateTrapezeArea(10, 10, 2)
	calculateCircleArea(12)
	// si necesitas hacer un escape por que no necesitas el segundo valor se usa el guion bajo como segundo _
	//func1return, _ := doubleReturn(2)
	func1return, func2return := doubleReturn(2)
	fmt.Printf("Los valores regresados son %d y %d", func1return, func2return)
	looping()

	//condiciones
	if isPair(6) {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is odd")
	}
	if isValidUser("Alpha5", "MyPassword") {
		fmt.Println("Credentials are valid")
	} else {
		fmt.Println("Credentials aren't valid")
	}

	switches()
	arraySlices()
	fmt.Println(palindrome("anita lava la tina"))
	makingMaps()
	initializeStruct()

	//caling mypackage
	var myCar = myPackage.Car
	myCar.Brand = "Ferrari 2"
	fmt.Println(myCar)

	//Los comparadores y operadores logicos, se usan similar a js exceptuando por la cohersion que en go no existe
}
