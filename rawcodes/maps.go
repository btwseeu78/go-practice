package main

import "fmt"

func Maps() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 7

	fmt.Println("map:", m)

	// get a value for the key with name
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// delete a key

	delete(m, "k2")
	fmt.Println("post delete map:", m)
	_, error := m["k2"]
	fmt.Println("blank key: ", error)

}
