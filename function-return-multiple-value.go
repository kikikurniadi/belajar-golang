package main

import "fmt"

func getFullName() (string, string, int) {
	return "Eko", "Khannedy", 12
}

func main() {
	firstName, lastName, _ := getFullName()
	fmt.Println(firstName, lastName)
}
