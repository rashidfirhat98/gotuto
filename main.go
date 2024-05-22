package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string)  {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
	
}

func main() {
	fn1, sn1 := getInitials("uzumaki naruto")
	fn2, sn2 := getInitials("uchiha sasuke")
	fn3, sn3 := getInitials("gaara")
	
	fmt.Println(fn1, sn1)
	fmt.Println(fn2, sn2)
	fmt.Println(fn3, sn3)

}