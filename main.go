package main

import "fmt"

//func somar(a int, b int) int {
//return a + b
//}

type Produto struct {
	nome  string
	preco float64
}

func (p Produto) exibirInformacoes() {
	fmt.Println("Nome do produto:", p.nome)
	fmt.Println("Preço do produto:", fmt.Sprintf("R$ %.2f", p.preco))
}

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Apresentar() {
	fmt.Println("Olá, meu nome é", p.Nome, "e tenho", p.Idade, "anos.")
}

func main() {

	produto1 := Produto{nome: "Notebook", preco: 3500.00}
	produto1.exibirInformacoes()

	//pessoa1 := Pessoa{Nome: "João", Idade: 30}
	//pessoa1.Apresentar()

	// nomes := map[string]int{
	// 	"João":  12,
	// 	"Maria": 25,
	// 	"Pedro": 30,
	// 	"Ana":   20,
	// }

	// nomes["Lucas"] = 35
	// nomes["Carla"] = 28

	// fmt.Println("Nomes:", len(nomes))

	// fmt.Println("Número de nomes:", len(nomes))

	// i := 0

	// for nome, idade := range nomes {
	// 	fmt.Println("Índice:", i, "Nome:", nome, "idade:", idade)
	// 	i++
	// }

	//resultado := somar(5, 10)
	//fmt.Println("O resultado da soma é:", resultado)

	//for numero := 1; numero <= 101-1; numero++ {
	//fmt.Println(numero)
	//}
}
