package main

import "fmt"

type Car struct {
	brand string
	year  int
	model string
	color string
	owner string
}

func initializeStruct() {
	//primera manera
	myCar := Car{brand: "Toyota", year: 2018, model: "B-type", color: "Rojo", owner: "Eliaz Bobadilla"}
	fmt.Println("Los Datos de mi auto son:", myCar)

	//segunda manera
	var otherCar Car
	otherCar.brand = "Ferrari"

	fmt.Println(otherCar)
}
