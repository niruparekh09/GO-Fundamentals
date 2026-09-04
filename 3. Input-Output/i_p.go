package main

import "fmt"

func main() {
	op()
	ip()
}

func op() {
	var i, j string = "Hello", "World"

	// Same line print
	fmt.Print(i)
	fmt.Print(j, "\n")
	fmt.Print(i, " ", j)

	fmt.Println()

	// Different line print
	fmt.Println(i)
	fmt.Println(j)

	// Printf() function first formats its argument based on the given formatting verb and then prints them
	// %v is used to print the value of the arguments
	// %T is used to print the type of the arguments

	var x string = "Hello"
	var y int = 15

	fmt.Printf("i has value: %v and type: %T\n", x, x)
	fmt.Printf("j has value: %v and type: %T", y, y)

}

func ip() {
	fmt.Println("\n\n--- Input Section ---")

	// 1. Using fmt.Scan() (Reads space-separated values, stops at a space or newline)

	var firstName string
	var age int
	fmt.Print("Enter your first name and age (space-separated): ")
	// Pass pointers (&) so Scan can modify the variables directly
	fmt.Scan(&firstName, &age)
	fmt.Printf("Scan Result -> Name: %v, Age: %v\n", firstName, age)

	// 2. Using fmt.Scanln() (Similar to Scan, but strictly stops when you press Enter/newline)
	var lastName string
	fmt.Print("Enter your last name: ")
	fmt.Scanln(&lastName)
	fmt.Printf("Scanln Result -> Last Name: %v\n", lastName)
}
