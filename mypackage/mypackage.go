package mypackage

import "fmt"

// las mayusculas y las minisculas influyen mayuscula al inicio = publico de lo contrario privado

// Car retrieve car
type Car struct {
	Brand string
	Year  int
	Model string
	Color string
	Owner string
}

type carPrivate struct {
	barnd string
	year  int
}

func TestMethod() {
	fmt.Println("test method")
}

//asi se definen bajo un tipo asociado
// el primer parentesis es el reciver y es a que struct o clase pertenece esa funcion
func (c Car) InstanceMethodOnly() {
	fmt.Println(c.Brand)
}
