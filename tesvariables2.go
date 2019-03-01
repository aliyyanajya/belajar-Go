package main

import "fmt"

/*
var (
	nama   = "David Maulana"
	umur   = 17
	alamat = "Jakarta"
)
*/

func main() {

	nama, umur, alamat := "David Maulana", 17, "Jakarta"

	fmt.Printf("%s berumur %d tahun\n", nama, umur)
	fmt.Println("Beralamat di ", alamat)
}
