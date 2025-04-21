package main

import "fmt"

func main() {
	// var name string bisa seperti ini, atau bisa seperti ini

	var name = "Rosyid Mukti"
	fmt.Println(name)

	name = "Rosyid Mukti Wibowo"
	fmt.Println(name)

	// Tanpa menggunakan var, untuk deklarasi pertama bisa seperti ini
	nama := "Rotibow"
	fmt.Println(nama)

	// Multiple Variable
	var (
		firstName  = "Rosyid"
		middleName = "Mukti"
		lastName   = "Wibowo"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	// Semua variable harus digunakan agar tidak error
}
