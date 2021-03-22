package main

import "fmt"

func makingMaps() {
	m := make(map[string]int)
	doubleMap := make(map[string]map[string]int)

	//el nested map se debe initializar o si no quedaria nil al momento de asignar valor
	doubleMap["Jose"] = make(map[string]int)
	doubleMap["Jose"]["Edad"] = 13

	m["Jose"] = 14
	m["Pepito"] = 20

	//si se accede a algo que no existe al map devuelve el zero value

	delete(m, "Jose")

	// el ok es busqueda, si existe
	value, ok := m["Jose"]
	fmt.Println(value, ok)

	//Recorrido simple
	for iterator, value := range m {
		fmt.Println(iterator, value)

	}

	//Recorrido doble
	for i := range doubleMap {
		for j, nv := range doubleMap[i] {
			fmt.Println(i, j, nv)
		}
	}
}
