package main

import (
	"fmt"
)

func main() {
	var names [3]string

	names[0] = "Rosyid"
	names[1] = "Mukti"
	names[2] = "Wibowo"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(len(names)) // Untuk mendapatkan panjang array
	fmt.Println(names)

	names[0] = "Rotibow" // Mengubah data di posisi index
	fmt.Println(names[0])

	// Membuat array secara langsung
	var values = [...]int{ // [...] jika tidak ingin menentukan berapa panjang arraynya
		90,
		85,
		80,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}
