package main
import (
	"fmt"
	"math"
)

//Menghitung Fungsi X
func main() {
	//Deklarasi Variabel
	var x float64
	var fx float64

	//Baris Input
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	//Program
	fx = (math.Pow(x, 3) + (3 * x)) / (math.Pow(x, 4) - (3 * math.Pow(x, 2)) + 4)

	//Baris Output
	fmt.Println("Nilai fungsi x adalah", fx)
}