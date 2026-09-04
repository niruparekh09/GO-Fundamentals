package main

import "fmt"

func main() {
	boolean()
	integer()
	floating()
	strings()
}

func strings() {
	/*------------------STRING------------------*/

	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}

func floating() {
	/*------------------FLOAT------------------*/

	var x2 float32 = 123.78
	var y2 float32 = 3.4e+38

	fmt.Println()
	fmt.Printf("Type: %T, value: %v\n", x2, x2)
	fmt.Println()

	fmt.Printf("Type: %T, value: %v", y2, y2)

	var x3 float64 = 1.7e+308
	fmt.Println()
	fmt.Printf("Type: %T, value: %v", x3, x3)
}

func integer() {
	/*------------------INTEGER------------------*/

	// Signed Integers
	var x int = 500
	var y int = -4500
	fmt.Println()
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Println()
	fmt.Printf("Type: %T, value: %v", y, y)

	// Unsigned Integers
	var x2 uint = 500
	var y2 uint = 4500
	fmt.Println()
	fmt.Printf("Type: %T, value: %v", x2, x2)
	fmt.Println()
	fmt.Printf("Type: %T, value: %v", y2, y2)

	fmt.Println()
}

func boolean() {
	/*------------------BOOL------------------*/

	var b1 bool = true
	var b2 = true
	var b3 bool
	b4 := true

	fmt.Println(b1) // t
	fmt.Println(b2) // t
	fmt.Println(b3) // f
	fmt.Println(b4) // t
}
