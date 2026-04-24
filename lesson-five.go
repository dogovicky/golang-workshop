package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("============== Loops =================")

	//x := 0

	// Simple while loop

	//for x < 5 {
	//	fmt.Println("Value: ", x)
	//	x++
	//}

	// Traditional for loop

	//for i := 0; i < 5; i++ {
	//	fmt.Println("Value: ", i)
	//}

	names := []string{"Mario", "Luigi", "Danesh", "Patel"}
	sort.Strings(names)

	// Looping through slices
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, name := range names {
		fmt.Println("Name: ", name)
	}

}
