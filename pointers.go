package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}

//coge el valor en memoria que tiene la clase para duplicarla *
func (myPC *pc) duplicateDisk() {
	myPC.disk = myPC.disk * 2
}

//manera de personalizar el output de una clase o struct
func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de Disco y es un sistema %s.", myPC.ram, myPC.disk, myPC.brand)
}

func pointers() {
	//& es el puntero direccion de memoria
	//* es el valor de la memoria
	var a int = 50
	b := &a

	//aunque modifique b a igual sera modificada
	*b = 100
	fmt.Println(a)

	fmt.Println(*b, b)

	myPC := pc{ram: 16, disk: 200, brand: "asus"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRam()
	myPC.duplicateDisk()

	fmt.Println(myPC)
}
