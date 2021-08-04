package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "928374529378450928734"
	var marredStatus Married = false

	fmt.Println(noKtpEko)
	fmt.Println(marredStatus)
}
