package main

import "fmt"

// a := 1 <---- Will give error
var a int
var b int = 2
var c = 3

func main() {
	basics()
	adv()
	constants()
}

func basics() {
	var student1 string = "John" // Type given

	// Type inferred
	var student2 = "Jack"
	student3 := "Dean"

	fmt.Println(student1, "--", student2, "--", student3)

	// Decalre without initial value
	var a string
	var b int // <--- Will override the global variable
	var c uint64

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Value assignment after decalration
	var student4 string
	student4 = "Sam"
	fmt.Println(student4)

	// Changing value of variable after declaration
	student1 = "Harry"
	fmt.Println(student1)
}

func adv() {
	// Declaring multiple variables at once
	var a, b, c, d int = 1, 2, 3, 4

	e, f, g, h := 5, 6, 7, 8
	var j, k, l, m = "Sam", "Dean", "Cass", "Jack"

	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g, h)
	fmt.Println(j, k, l, m)

	// Declaring multiple variables of different type at once
	// var n, o int = 9, "Cool" <----- Error
	var n, o = 9, "Cool"
	p, q := true, 3.14

	fmt.Println(n, o)
	fmt.Println(p, q)

	// Multible variable devlaration in a block
	var (
		r int
		s int  = 1
		t bool = false
	)

	fmt.Println(r, s, t)
}

/*------------CONTSANTS------------*/

const PI = 3.14
const SIN30 float32 = 0.5

// Multiple const
const (
  X int = 1
  Y = 3.14
  Z = "Hi!"
)

func constants() {
	fmt.Println("PI", PI)
	fmt.Println("Sin 30°:", SIN30)

	// PI = 3.1 <---- Error

	// Can be declared inside a function
	const A int = 1
	fmt.Println(A)
}
