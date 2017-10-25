package main

import (
	"fmt"
	"runtime"
)

func main() {
	var cont int
	var a int

	cont = 0

	for i := 0; i < 10; i++ {
		cont += 1
		fmt.Println(cont)
	}

	cont = 1

	for ; cont < 10; {
		cont += cont
		fmt.Println(cont)
	}	

	cont = 1
	
	for cont < 10 {
		cont += cont
		fmt.Println(cont)
	}

	a = 15

	if a > 10 {
		fmt.Println("a é maior que 10");
	}

	if b := 20; b > a {
		fmt.Println("b é maior que a");
	}

	switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Boa garoto!!");
		default:
			fmt.Println(os);
	}

}