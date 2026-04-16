package main

import (
	"fmt"
	"sort"
)

/*
 * The standard library
 */

func main() {
	fmt.Println("=========== The Standard Library =============")

	// strings package
	//greeting := "Hello there friend!"
	//
	//fmt.Println(strings.Contains(greeting, "Hello"))
	//fmt.Println(strings.ReplaceAll(greeting, "Hello", "Good Afternoon"))
	//fmt.Println(strings.ToUpper(greeting))
	//fmt.Println(strings.Index(greeting, "llo"))
	//fmt.Println(strings.Split(greeting, " "))

	// sort package
	ages := []int{26, 28, 37, 18, 25, 27, 39, 45, 57}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 26) // searches the sorted slice
	fmt.Println(index)

	names := []string{"Yoshi", "Danesh", "Richard", "Mario", "Bighead"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Bighead"))

}
