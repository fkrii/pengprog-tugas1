package main
import (
	"fmt"
)

//Menghitung Y
func main() {
	//Deklarasi Variabel
	var x float64
	var y float64

	//Baris Input
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	//Program
	y = (3 * x - 5) * (2 * x + 1)

	//Baris Output
	fmt.Println("Nilai y adalah", y)
}