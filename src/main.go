package main

import "fmt"

//func somar(a int, b int) int {
//return a + b
//}

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Apresentar() {
	fmt.Println("Olá, meu nome é", p.Nome, "e tenho", p.Idade, "anos.")
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("não é possível dividir por zero")
	}
	return a / b, nil
}

func main() {

	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado da divisão:", resultado)
	}

	//produto1 := Product{Name: "Notebook", Price: 3500.00}
	//produto1.DisplayInfo()

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
