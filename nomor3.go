package main

import "fmt"

func main(){
	jam := 1
	menit := 30
	detik := 40

	detik_jam := jam * 3600
	detik_menit := menit * 60
	
	total_detik := detik_jam + detik_menit + detik

	fmt.Print("total detik adalah : ", total_detik)
}