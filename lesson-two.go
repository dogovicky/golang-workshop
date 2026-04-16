package main

import (
	"fmt"
)

/*
 * The "fmt" package
 */

func main() {

	name := "Dogo Vicky"
	age := 21

	// Print
	fmt.Print("Lesson two \n")
	fmt.Print("Lesson two \n")

	// Println
	fmt.Println("Hello Goers!")
	fmt.Println("Happy to join")

	//fmt.Println("My age is ", age, " and my name is ", name)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("My age is %v and my name is %v\n", age, name) // v for variable
	fmt.Printf("My age is %v and my name is %q\n", age, name) // q places the strings in quotes
	fmt.Printf("Age is of type %T\n", age)                    // Returns the type of variable
	fmt.Printf("You scored %0.1f points \n", 22.55)           // Float numbers

	// Springf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %q\n", age, name)
	fmt.Println(str)
}
