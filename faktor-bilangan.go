package main

import "fmt"

func main() {
	var bilangan int = 6

	// menggunakan perulangan
	// inisial ade :sama dengan 1 apakah ade kurang dari sama dengan nilai bilangan jika ia ulangi ade++
	for ade := 1; ade <= bilangan; ade++ {
		// kondisi jika bilangan modulo ade == 0 terpenuhi maka cetak ade
		if bilangan % ade == 0 { // 0 mencari
			fmt.Println(ade)
		}
	}
}