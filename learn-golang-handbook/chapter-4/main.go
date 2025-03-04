package main

import (
	"errors"
	"fmt"
)

/*
   Go Leverages Guard Clauses (https://blog.boot.dev/clean-code/guard-clauses/)
*/

func main() {
	//add
	fmt.Printf("add(2, 3) = %d\n", add(2, 3))

	//Ignore Return Values
	// ignore y value
	x, _ := getPoint()
	fmt.Printf("x = %d\n", x)
}

func add(x, y int) int {
	return x + y
}

func getPoint() (x int, y int) {
	return 3, 4
}

func getCoords() (x, y int) {
	// x and y are initialized with zero values

	return // automatically returns x and y
}

func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
}
