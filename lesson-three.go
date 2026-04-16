package main

import (
	"fmt"
)

/*
 * Arrays and slices
 */

func main() {

	fmt.Println("============= Arrays and Slices =============")

	// Arrays
	// The length of an array is fixed

	var ages = [5]int{10, 10, 30, 18, 25}
	fmt.Println(len(ages))
	fmt.Println(ages)

	names := [4]string{"Yoshi", "Mario", "Peach", "Bouser"}
	names[1] = "Luigi"
	fmt.Println(len(names))
	fmt.Println(names)

	// Slices
	// For slices, the size is not defined, the slice can scale
	scores := []int{12, 20, 49, 49}
	scores[2] = 46
	scores = append(scores, 85) // We can append to slices, but it returns a new slice

	fmt.Println(scores)
	fmt.Println(len(scores))

	// Slice ranges
	range1 := names[1:3] //includes position 1 and excludes position 3
	fmt.Println(range1)

	range2 := names[2:] // Goes up to the ed of the array/slice
	fmt.Println(range2)

	range3 := names[:3] // From position 0 to 3, does not include the third one
	fmt.Println(range3)

	range1 = append(range1, "Cooper")
	fmt.Println(range1)

}
