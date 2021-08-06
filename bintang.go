package main

import "fmt"

func playWithAsterik(n int) {
	//loping pertama
	// ade sama dengan 1 apakah ade kurang dari sama dengan n jika iya lakukan pengulangan ade++
	for ade := 1; ade <= n; ade++ {
		//loping ke 2
		// jarak sama dengan 1 apakah jarak kurang dari sama dengan nilai n-ade jika iya ulangi jarak
		for jarak := 1; jarak <= n-ade; jarak++ {
			fmt.Printf(" ")
		}
		//loping ke 3
		for kolom :=1; kolom <= ade; kolom++{
			fmt.Printf("* ")
		}
		fmt.Println()
	}

}

func main() {
	playWithAsterik(5)
}