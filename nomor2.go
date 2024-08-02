package main

import "fmt"

func main(){
	bilangan1 := 10
	bilangan2 := 20
	bilangan3 := 30
	
	if bilangan1 > bilangan2 && bilangan1 > bilangan3{
		fmt.Print("bilangan terbesar adalah ", bilangan1)
	} else if bilangan2 > bilangan1 && bilangan2 > bilangan3{
		fmt.Print("bilangan terbesar adalah ", bilangan2)
	} else {
		fmt.Print("bilangan terbesar adalah ", bilangan3)
	}
}
