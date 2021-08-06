package main

import "fmt"

func main() {
	var bilangan int = 6

	// menggunakan perulangan
	// inisial ade :sama dengan bilanga apakah ade kurang dari sama dengan nilai bilangan jika ia ulangi ade--
	for ade := bilangan; ade >= 1; ade-- {
		// kondisi jika bilangan modulo ade == 0 terpenuhi maka cetak ade
		if bilangan % ade == 0 { // 0 mencari
			fmt.Println(ade)
		}
	}
}