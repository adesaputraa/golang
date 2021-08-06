package main

import "fmt"

// membuat fungsi atau function
func ade1(nama string, studentScore int)  {
		//proses menggunakan pengkondisian
		var ade string = ""
		// jika nilai studentScore lebih besar sama dengan 80 dan jika nilai studentScore kurang dari sama dengan 100 maka cetak nilainya
		if studentScore >= 80 && studentScore <= 100 {
		// cetak nilai
			ade = "A"
		}else if studentScore >= 65 && studentScore <= 79 {
			ade = "B+"
		}else if studentScore >= 50 && studentScore <= 64 {
			ade = "B"
		}else if studentScore >= 35 && studentScore <= 49 {
			ade = "C"
		}else if studentScore >= 0 && studentScore <= 34 {
			ade = "D"
		// else jika kondisinya tidak terpenuhin 
		} else {
			ade = "Invalid"
		}
		fmt.Println(nama, ade)
}

func main (){

	// memanggil function
	ade1("ade", 100)
	
	// var studentScore int = 79 // nilai input

	// //proses menggunakan pengkondisian
	// var ade string = ""
	// // jika nilai studentScore lebih besar sama dengan 80 dan jika nilai studentScore kurang dari sama dengan 100 maka cetak nilainya
	// if studentScore >= 80 && studentScore <= 100 {
	// // cetak nilai
	// 	ade = "A"
	// }else if studentScore >= 65 && studentScore <= 79 {
	// 	ade = "B+"
	// }else if studentScore >= 50 && studentScore <= 64 {
	// 	ade = "B"
	// }else if studentScore >= 35 && studentScore <= 49 {
	// 	ade = "C"
	// }else if studentScore >= 0 && studentScore <= 34 {
	// 	ade = "D"
	// // else jika kondisinya tidak terpenuhin 
	// } else {
	// 	ade = "Invalid"
	// }

	// fmt.Println(ade)
}

