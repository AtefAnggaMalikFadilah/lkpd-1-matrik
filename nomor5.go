package main

import "fmt"

func main(){
	for i := 1; i <= 50; i++ {
		if i % 2 == 0{
			fmt.Print(i, " adalah bilangan genap\n")
		} else {
			fmt.Print(i, " adalah bilangan ganjil\n")
		}
	}
}