package main

import "fmt"

func pangkat(base, pangkat int) int {
	ade := 1 //inisial niali awal

	for ade1 :=1; ade1 <= pangkat; ade1++{
		ade = ade * base
	}
	return ade


}

func main() {
	fmt.Println(pangkat(2,3)) //8
	fmt.Println(pangkat(5,3)) //125
	fmt.Println(pangkat(10,3)) //100
	fmt.Println(pangkat(2,5)) //32
	fmt.Println(pangkat(7,3)) //343
}