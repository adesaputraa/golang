package main

import "fmt"

func cetakTablePerkalian(number int) {
	hitung := 0
	for ade := 1; ade <= number; ade++ {
		for ade1 := 1; ade1 <= number; ade1++ {
			hitung = ade * ade1
			fmt.Printf("%3v", hitung)
		}
		fmt.Println()
	}

}

func main() {
	cetakTablePerkalian(9)
}