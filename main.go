package main

import "fmt"

/*
	Struct representing a vending machine
	Stores model, address and capacity.
	String() method to override print

*/

func main() {
	var cewmachine = machine{
		model: "18239342",
		address: address{
			street: "101 MItchell St",
			city:   "Arlington",
			state:  "TX",
			zip:    "76010",
		},

		capacity: 100,
	}

	var p printer = cewmachine
	fmt.Println(p)
	// fmt.Println(cewmachine)
	// fmt.Println("Capacity: ", cewmachine.capacity)
}
