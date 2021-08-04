package main

import "fmt"

func sayHElloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "abibah"
	sayHElloTo("Eko", "Kurniawan")
	sayHElloTo(firstName, "admin")
}
