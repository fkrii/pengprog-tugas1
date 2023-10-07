package main
import (
	"fmt"
)

//Rumus Konversi Celcius ke Kelvin
func main() {
	//Deklarasi Variabel
	var celcius float64
	var kelvin float64

	//Baris Input
	fmt.Print("Masukkan suhu celcius: ")
	fmt.Scan(&celcius)

	//Program
	kelvin = celcius + 273

	//Baris Output
	fmt.Println("Suhu dalam kelvin adalah", kelvin, "K")
}