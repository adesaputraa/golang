package main

import "fmt"

// menggunakan fungsi. fungsi digunakan untuk mengambil sebuah nilai atau data jika terperlukan
func luasSegitga(ade1 int, ade2 int) int  {

	// membuat variabel untuk proses hitung
	// var ade3 int = ade1 * ade2 / 2

	// mengembalikan kevariabel
	// return ade3

	// bisa juga dengan cara ini 
	return ade1 * ade2 / 2
	

	
	
}

func main() {

	// memanggil sebuah fungsi
	ade3 := luasSegitga(20, 25)
	// mencetak hasil
	fmt.Println(ade3)


	// cara 2
	// membuat variabel
	var alas float64 = 20
	var tinggi float64 = 25
	// variabel hasil untuk proses menghitung
	var hasil = alas * tinggi / 2

	//cetak hasil perhitungan
	fmt.Println(hasil)
}