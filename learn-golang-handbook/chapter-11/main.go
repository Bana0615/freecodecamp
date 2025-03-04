package main

import (
	"fmt"
	"io"
	"os"
)

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

// Closures
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

// Anonymous Functions
// doMath accepts a function that converts one int into another
// and a slice of ints. It returns a slice of ints that have been
// converted by the passed in function.
func doMath(f func(int) int, nums []int) []int {
	var results []int
	for _, n := range nums {
		results = append(results, f(n))
	}
	return results
}

func main() {
	fmt.Println(aggregate(2, 3, 4, add))
	// prints 9
	fmt.Println(aggregate(2, 3, 4, mul))
	// prints 24

	// Function Currying - Function currying is the practice of writing a function that takes a function (or functions) as input, and returns a new function.
	fmt.Println("--- Function Currying ---")
	squareFunc := selfMath(mul)
	doubleFunc := selfMath(add)

	fmt.Println(squareFunc(5))
	// prints 25

	fmt.Println(doubleFunc(5))
	// prints 10

	//Closures
	fmt.Println("--- Closures ---")
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive

	// Anonymous Functions
	fmt.Println("--- Anonymous Functions ---")
	nums := []int{1, 2, 3, 4, 5}

	// Here we define an anonymous function that doubles an int
	// and pass it to doMath
	allNumsDoubled := doMath(func(x int) int {
		return x + x
	}, nums)

	fmt.Println(allNumsDoubled)
	// prints:
	// [2 4 6 8 10]
}

// CopyFile copies a file from srcName to dstName on the local filesystem.
// Deferred functions are typically used to close database connections, file handlers and the like.
func CopyFile(dstName, srcName string) (written int64, err error) {

	// Open the source file
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// Close the source file when the CopyFile function returns
	defer src.Close()

	// Create the destination file
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	// Close the destination file when the CopyFile function returns
	defer dst.Close()

	return io.Copy(dst, src)
}
