package main

import "fmt"

/*
Pointer -A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to" the location of
where the data is stored NOT the actual data itself.
*/

// Pointer Method Receivers
type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

func (c car) setColor2(color string) {
	c.color = color
}

type circle struct {
	x      int
	y      int
	radius int
}

func (c *circle) grow() {
	c.radius *= 2
}

func main() {
	myString := "hello"
	myStringPtr := &myString
	fmt.Println(*myStringPtr)
	*myStringPtr = "world"
	fmt.Println(myString)

	// Pointer receiver
	fmt.Println("--- Pointer receiver ---")
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"

	// Non-pointer receiver
	fmt.Println("--- Non-pointer receiver ---")
	d := car{
		color: "white",
	}
	d.setColor2("blue")
	fmt.Println(d.color)
	// prints "white"

	//
	fmt.Println("--- Circle ---")
	e := circle{
		x:      1,
		y:      2,
		radius: 4,
	}

	// notice e is not a pointer in the calling function
	// but the method still gains access to a pointer to c
	e.grow()
	fmt.Println(e.radius)
	// prints 8
}
