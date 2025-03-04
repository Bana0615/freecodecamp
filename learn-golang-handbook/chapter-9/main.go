package main

import "fmt"

func main() {
	//Arrays - Arrays are fixed in size
	var myInts [10]int
	fmt.Println(myInts)

	fmt.Println("Primes")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//Slices are dynamic
	/* Syntax for slices
		arrayname[lowIndex:highIndex]
		arrayname[lowIndex:]
		arrayname[:highIndex]
		arrayname[:]

	    Either lowIndex or highIndex or both can be omitted to use the entire array on that side.
	*/

	fmt.Println("Slices")
	// func make([]T, len, cap) []T
	mySlice := make([]int, 5, 10)
	fmt.Println(mySlice)

	// the capacity argument is usually omitted and defaults to the length
	mySlice2 := make([]int, 5)
	fmt.Println(mySlice2)

	mySlice3 := []string{"I", "love", "go"}
	fmt.Println(mySlice3)
	fmt.Println("mySlice3 length = ", len(mySlice3))
	fmt.Println("mySlice3 capacity = ", cap(mySlice3))

	// Variadic Functions
	final := concat("Hello ", "there ", "friend!")
	fmt.Println(final)
	// Output: Hello there friend!

	// Spread operator
	names := []string{"bob", "sue", "alice"}
	printStrings(names...)

	// Range Over a Slice
	fruits := []string{"apple", "banana", "grape"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}

// Variadic Functions - A variadic function receives the variadic arguments as a slice.
func concat(strs ...string) string {
	final := ""
	// strs is just a slice of strings
	for _, str := range strs {
		final += str
	}
	return final
}

// Spread operator - The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}
