package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func sayBye(name string) {
	fmt.Printf("Bye, %s!\n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {

	fmt.Println("============== Functions =================")
	//sayGreeting("Mario")
	//sayBye("Danesh")

	names := []string{"Mario", "Luigi", "Danesh", "Patel"}
	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)

	area1 := circleArea(7.5)
	area2 := circleArea(10.5)
	fmt.Println("Circle Area = ", area1, area2)
	fmt.Printf("Circle 1 %0.2f and Circle 2 %0.2f\n", area1, area2)

}
