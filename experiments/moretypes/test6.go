package main

import (
	"fmt"
)

type Pessoa struct {
	nome string
	idade int
	fuma bool
}

func main() {
	i, j := 41, 14

	p := &i
	fmt.Println(p, *p, i, j)

	*p = 144 * 10
	fmt.Println(p, *p)
	
	p1 := Pessoa{"Lucas", 23, false}
	fmt.Println("Pessoa 1 se chama:", p1.nome)

	p2 := Pessoa{"Maria", 20, true}
	pp2 := &p2
	pp2.nome = "Mario"
	fmt.Println(p2, *pp2)

	//slices 
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	//slices literais
	q := []int{2, 3, 4, 5}
	fmt.Println(q)

	r := []bool{true, false, false}
	fmt.Println(r)

	s1 := []struct {
		i int 
		b bool
	}{
		{2, true},
		{6, false},
	}
	fmt.Println(s1) 

}