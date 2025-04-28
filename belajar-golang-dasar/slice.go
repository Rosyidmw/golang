package main

import "fmt"

func main() {
	name := [...]string{"Rosyid", "Mukti", "Wibowo", "Isma", "Fathanah", "Rotibowif"}

	slice1 := name[4:6]
	fmt.Println(slice1)

	slice2 := name[:3]
	fmt.Println(slice2)

	slice3 := name[3:]
	fmt.Println(slice3)

	slice4 := name[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)

	daySlice1 := days[5:] // Sabtu dan Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru") // Untuk manambahkan data
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(daySlice1)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Rosyid"
	newSlice[1] = "Mukti"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Wibowo")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Rotibow"
	fmt.Println(newSlice2)
	fmt.Println(newSlice2)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Perbedaan cara buat array dan slice => Array [...], Slice []

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
