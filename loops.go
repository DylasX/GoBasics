package main

import (
	"fmt"
)

func looping() {

	//for condicional
	for i := 0; i <= 10; i++ {
		fmt.Printf("Posición %d ", i)
	}

	//for while
	counter := 10
	for counter >= 0 {
		fmt.Println("For while inverted")
		counter--
	}

	//for forever like cron
	// for {
	// 	fmt.Println("Esto es un cron, para la aplicación")
	// 	time.Sleep(1 * time.Second)
	// }

	//para mas explicacion de esto mirar array
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}
}
