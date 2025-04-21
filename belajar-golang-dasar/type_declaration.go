package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRosyid NoKTP = "111111111111"

	var contoh string = "2222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpRosyid)
	fmt.Println(contohKTP)
}
