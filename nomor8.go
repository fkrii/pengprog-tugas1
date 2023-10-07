package main
import (
	"fmt"
)

//Konversi Gram ke Kilogram
func main() {
	//Deklarasi Variabel
	var b1 int
	var b2 int
	var b3 int

	//Baris Input
	fmt.Print("Masukkan berat 3 buah belerang: ")
	fmt.Scan(&b1, &b2, &b3)

	//Program
	var b1kg = b1 / 1000
	var b2kg = b2 / 1000
	var b3kg = b3 / 1000

	//Baris Output
	fmt.Print(b1, "g,", b2, "g,", b3, "g")
	fmt.Print(b1kg, "Kg,", b2kg, "Kg,", b3kg, "Kg")
}