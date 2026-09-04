package main

import "fmt"

func main() {
	// Defined Length
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// Inferred Length
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr3)
	fmt.Println(arr4)

	// Accessing values at index
	fmt.Println(arr1[2])

	// Changing val at index
	arr1[1] = 20
	fmt.Println(arr1)

	// ARRAY Initialization
	arr5 := [5]int{}              // Not Initialized
	arr6 := [5]int{1, 2}          // Partially Initialized
	arr7 := [5]int{1, 2, 3, 4, 5} // Fully Initialized
	arr8 := [5]int{1: 40, 2: 22}  // Initialize Only Specific Elements

	fmt.Println(arr5) // [0 0 0 0 0]
	fmt.Println(arr6) // [1 2 0 0 0]
	fmt.Println(arr7) // [1 2 3 4 5]
	fmt.Println(arr8) // [0 40 22 0 0]

	// Length of Array
	fmt.Println(len(arr4))

}
