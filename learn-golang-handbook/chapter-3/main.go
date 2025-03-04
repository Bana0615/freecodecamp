package main

import "fmt"

func main() {
	a := fmt.Sprintf("I am %v years old", 10)
	print(a + "\n")
	// I am 10 years old

	b := fmt.Sprintf("I am %v years old", "way too many")
	print(b + "\n")
	// I am way too many years old

	//%s - Interpolate a string
	c := fmt.Sprintf("I am %s years old", "way too many")
	print(c + "\n")
	// I am way too many years old

	//%d - Interpolate an integer in decimal form
	d := fmt.Sprintf("I am %d years old", 10)
	print(d + "\n")
	// I am 10 years old

	//%f - Interpolate a decimal
	e := fmt.Sprintf("I am %f years old", 10.523)
	print(e + "\n")
	// I am 10.523000 years old

	// The ".2" rounds the number to 2 decimal places
	f := fmt.Sprintf("I am %.2f years old", 10.523)
	print(f + "\n")
	// I am 10.53 years old

	//Contitionals
	height := 8

	if height > 6 {
		fmt.Println("You are super tall!")
	} else if height > 4 {
		fmt.Println("You are tall enough!")
	} else {
		fmt.Println("You are not tall enough!")
	}

	//initial statement of an if block
	//The variable(s) created in the initial statement are only defined within the scope of the if body.
	email := ""

	if length := len(email); length < 1 {
		fmt.Println("Email is invalid")
	}
}
