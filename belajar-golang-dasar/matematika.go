package main

import "fmt"

func main() {
	var a = 15
	var b = 5
	var d = 5
	var e = 2
	var c = a + b - d*e

	fmt.Println(c)

	//a = a + 10
	//fmt.Println(a)

	a += 10 // a = a + 10
	fmt.Println(a)

	// Unary operator
	var j = 1
	j++ // j = j + 1
	fmt.Println(j)
	j-- // j = j - 1
	fmt.Println(j)
}
