package main

import "fmt"

func main() {
	//Loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Loops in Go can omit sections of a for loop. For example, the CONDITION (middle part) can be omitted which causes the loop to run forever.
	// for INITIAL; ; AFTER {
	// do something forever
	// }

	//There are no while loops here is the sytanx to do a "while loop"
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}
	fmt.Println("plant has grown to ", plantHeight, "inches")

	//The continue keyword stops the current iteration of a loop and continues to the next iteration. continue is a powerful way to use the "guard clause" pattern within loops.
	fmt.Println("Using continue inside of a loop")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//Break out of a loop
	fmt.Println("Break out of a loop")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
