package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hello %s \n", text)
}

func goRoutine() {
	var wg sync.WaitGroup
	// say("world")
	//Agrego una go routine para que espere su ejecucion
	wg.Add(1)
	//ejecutado de forma concurrente
	go say("WORLD", &wg)

	go func() {
		fmt.Println("Annonymous function call")
	}()

	wg.Wait()
	//es como un async de js podemos hacer esperar haciendo

	//time.Sleep(1 * time.Second)
}
