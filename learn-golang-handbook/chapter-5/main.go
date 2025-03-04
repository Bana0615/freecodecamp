package main

import "fmt"

type car struct {
	make       string
	model      string
	height     int
	width      int
	frontWheel Wheel
	backWheel  Wheel
}

type Wheel struct {
	radius   int
	material string
}

type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	car
	bedSize int
}

func main() {
	myCar := car{}
	myCar.frontWheel.radius = 5
	fmt.Println(myCar)

	//truck
	lanesTruck := truck{
		bedSize: 10,
		car: car{
			make:  "toyota",
			model: "camry",
		},
	}

	fmt.Println(lanesTruck.bedSize)

	// embedded fields promoted to the top-level
	// instead of lanesTruck.car.make
	fmt.Println(lanesTruck.make)
	fmt.Println(lanesTruck.model)

	//Struct Methods
	r := rect{
		width:  5,
		height: 10,
	}

	fmt.Println(r.area())
	// prints 50
}

type rect struct {
	width  int
	height int
}

// area has a receiver of (r rect)
func (r rect) area() int {
	return r.width * r.height
}
