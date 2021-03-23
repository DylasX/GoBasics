package main

import "fmt"

// <- ese canal solo va a ser de entrada
func sayChannel(text string, c chan<- string) {
	//vamos a ingresar el texto a ese canal
	// en caso de asignar text = <- c
	c <- text
}

func exploreChannels() {
	// va a manejar solo 1 string a la vez aunque el limite puede ser opcional
	c := make(chan string, 1)
	// sayChannel("World", c)
	// go sayChannel("WORLD", c)

	//saque el string del canal y muestrelo
	//fmt.Println(<-c)

	c <- "hola"
	fmt.Println(len(c), cap(c))

	//close lo unico que hace es decirle a go que va a cerrar el canal
	//close(c)

	//Esto seria un error por que ya se cerro
	//c <- "Hola"

	//antes del range debe ser cerrada
	close(c)

	for message := range c {
		fmt.Println(message)
	}

	//select

	email1 := make(chan string, 1)
	email2 := make(chan string, 1)

	go sayChannel("hola", email1)
	go sayChannel("hola2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido en email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido en email2", m2)
		}
	}
}
