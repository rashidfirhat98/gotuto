package main

import "fmt"

func main() {
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"

	fmt.Println(nameFour)

	//ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint = 25

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 987972034293845.7
	// scoreThree := 1.5
}