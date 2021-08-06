package main

import "fmt"


func DrawXYZ(high int) string {
	ade := ""
	ade2 := (high*high)
	for ade3 := 1; ade3 <= ade2; ade3++ {
			if ade3 % 3 == 0{
				ade += "X "
			} else if ade3 % 2 == 0{
				ade += "Z "
			} else {
				ade += "Y "
			}
			if ade3 % high == 0 {
				ade += "\n"
			}
		}
		return ade

}

func main() {

	fmt.Println(DrawXYZ(5))

}