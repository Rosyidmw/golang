package main

import "fmt"

func main() {
	//var person = map[string]string{}
	//person["name"] = "Rosyid"
	//fmt.Println(person)

	person := map[string]string{
		"name": "Rotibow",
		"age":  "21",
	}
	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["name"])
	person["name"] = "Rosyid"
	fmt.Println(person["name"])

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Rosyid"
	book["ups"] = "Salah"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}
