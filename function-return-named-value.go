package main

import "fmt"

func getFullName2() (firstName, MiddleName, lastName string) {
	firstName = "Eko"
	MiddleName = "Kurniawan"
	lastName = "Khannedy"
	return
}

func main() {
	firstName, MiddleName, lastName := getFullName2()
	fmt.Println(firstName)
	fmt.Println(MiddleName)
	fmt.Println(lastName)

}
