package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	upperCaseName := strings.ToUpper(name)

	names := strings.Split(upperCaseName, " ")

	var initials []string

	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {

	fn, sn := getInitials("patel danesh")
	fmt.Println(fn, sn)
}
