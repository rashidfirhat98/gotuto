package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "--", v)
	}

	// ints as key type
	phonebook := map[int]string{
		112341234: "mario",
		989708978: "luigi",
		923452353: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[112341234])

	phonebook[923452353] = "bowser"
	fmt.Println(phonebook)

	phonebook[989708978] = "yoshi"
	fmt.Println(phonebook)

}