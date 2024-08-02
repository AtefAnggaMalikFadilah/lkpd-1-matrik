package main

import "fmt"

func main(){
	var total_detik int

	fmt.Print("Masukkan total detik: ")
	fmt.Scan(&total_detik)

	total_jam := total_detik / 3600
	sisa_detik := total_detik % 3600
	total_menit := sisa_detik / 60
	detik := sisa_detik % 60
	
	fmt.Print("Hasilnya adalah :\n", total_jam, " jam ", total_menit, " menit ", detik, " detik")
}