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

}