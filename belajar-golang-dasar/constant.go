package main

import "fmt"

func main() {
	// const variable yang tidak dapat dirubah isinya, namun tidak error jika tidak digunakan

	const firstName string = "Rosyid"
	const lastName = "Wibowo"

	// Multiple const
	const (
		namaDepan string = "Rosyid"
		namaAkhir        = "Wibowo"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	fmt.Println(namaDepan, namaAkhir)
}
