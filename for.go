package main

import "fmt"

func main() {

	for counter := 0; counter < 20; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"eko", "Kurniawan", "Khannedy"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
