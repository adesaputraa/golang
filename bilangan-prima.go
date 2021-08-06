package main

import "fmt"

func primeNumber(ade int) bool {
	// membuat variabel
	var ade1 int = 0
	// perulangan
	for ade2 :=1; ade2 <= ade; ade2++{
		if ade % ade2 == 0 {
			ade1+=1
		}
	}
	if ade1 == 2 {
		return true
	}else{
		return false
	}

}

func main() {
	fmt.Println(primeNumber(11)) // true
	fmt.Println(primeNumber(13)) // true
	fmt.Println(primeNumber(17)) // true
	fmt.Println(primeNumber(20)) // false
	fmt.Println(primeNumber(35)) // false
}