package main

import (
	"fmt"
	"math/rand"
	"math"
)

//constantes
const Ipca = 7.64

//variaveis
var a, b, c bool

var nome, idade, fuma = "Lucas", 23, false

//retorna um inteiro
func add(x, y int) int {
	return x + y
}

//retorna dois parametros
func nome_sobrenome(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

//retorno nomeado
func calc_qualquer_1(val int) (x, y int) {
	x = val * 4
	y = x * 8
	return 
}

//Variaveis e sua inicialização
//curta :=
//declarada var x int 
func test_variaveis() {
	var d, e, f bool
	var php, java, golang = true, false, "Oh man!"
	css := "ok"
	fmt.Println(a, b, d, e, f)
	fmt.Println(nome, idade, fuma, php, java, golang, css)
}

func conversao_de_tipos() {
	var a int = 10
	var b float64 = float64(a)
	var c uint = uint(b)

	fmt.Println(a, b, c)
}


func main() {
	fmt.Println("Meu número favorito é: ", rand.Intn(10))
	
	fmt.Println("-----------------");

	fmt.Printf("Agora você tem %g problemas", math.Sqrt(7))

	fmt.Println("");
	
	fmt.Println("-----------------");

	fmt.Println("Valor de Pi:", math.Pi)

	fmt.Println("-----------------");

	fmt.Println("5 + 5 é:", add(5, 5));

	fmt.Println("-----------------");
	
	fmt.Println("10 + 12 é:", add(10, 12));

	fmt.Println("-----------------");
	
	nome, sobrenome := nome_sobrenome("Lucas", "Burg")

	fmt.Println(nome, sobrenome);
	
	fmt.Println("-----------------");

	fmt.Println(calc_qualquer_1(2));

	test_variaveis()
	
	fmt.Println("-----------------");

	conversao_de_tipos()
	
	fmt.Println("-----------------");

	fmt.Println(Ipca);
	
}