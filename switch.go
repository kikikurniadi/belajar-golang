package main

import "fmt"

func main() {

	name := "kurniadi"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hallo Joko")
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan Dong")
		fmt.Println("Kenalan Dong")
	}

	/*switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Pnjang")
	case false:
		fmt.Println("Nama Sesuai")
	}*/

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lamayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
