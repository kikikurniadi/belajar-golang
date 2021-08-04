package main

import "fmt"

func main() {
	var name = "Kurniawan"

	if name == "EKO" {
		fmt.Println("Hello eko")
	} else if name == "Joko" {
		fmt.Println("Hello joko")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("terlalu Panjang")
	} else {
		fmt.Println("Nama terlalu panjang")
	}
}
