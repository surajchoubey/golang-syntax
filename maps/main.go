package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for key, value := range menu {
		fmt.Println(key, "->", value)
	}

	// ints as key type
	phonebook := map[int]string{
		8652099501: "peach",
		8879314511: "mario",
		9867543211: "luigi",
	}

	fmt.Println(phonebook[8652099501])
	fmt.Println(phonebook)

	phonebook[8879314511] = "bowser"
	phonebook[9867543211] = "yoshi"

	fmt.Println(phonebook)

}
