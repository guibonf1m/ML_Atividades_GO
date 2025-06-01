package main

import (
	"fmt"
	"strconv"
)

type Produto struct {
	id         int
	Nome       string
	Preco      float64
	Quantidade int
}
type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) descricao() string {
	return fmt.Sprintf("Nome: %s \n Idade:", p.Nome, p.Idade)
}

func (p Produto) exibir() string {
	return fmt.Sprintf("ID: %d \nNome: %s \nPreco: %.2f \nQuantidade: %d", p.id, p.Nome, p.Preco, p.Quantidade)
}

func (p Produto) valorTotal() float64 {
	if p.Preco > 0 && p.Quantidade > 0 {
		return p.Preco * float64(p.Quantidade) // converte o inteiro para float64, para que a multiplicação funcione corretamente.
	} else {
		return 0
	}
}

func main() {

	var Pessoa Pessoa

	fmt.Print("Digite o seu nome: ")
	fmt.Scanln(&Pessoa.Nome)

	fmt.Print("E qual é a sua idade: ")
	fmt.Scanln(&Pessoa.Idade)

	fmt.Println()

	Caneta := Produto{
		id:         1,
		Nome:       "Caneta",
		Preco:      2.50,
		Quantidade: 100,
	}

	Mesa := Produto{
		id:         2,
		Nome:       "Mesa",
		Preco:      50.00,
		Quantidade: 25,
	}

	Lapis := Produto{
		id:         3,
		Nome:       "Lapis",
		Preco:      150.00,
		Quantidade: 100,
	}

	fmt.Print("VOCÊ ACABA DE ACESSAR O MERCADO LIVRE, MAIOR E-COMMERCE DA AMÉRICA LATINA")
	fmt.Println()
	fmt.Println(Caneta.exibir())
	fmt.Println()
	fmt.Println(Mesa.exibir())
	fmt.Println()
	fmt.Println(Lapis.exibir())
	fmt.Println()
	fmt.Println("Me chamo, " + Pessoa.Nome + " e tenho " + strconv.Itoa(Pessoa.Idade) + " anos!")
	fmt.Println("Valor total de Canetas em estoque:", Caneta.valorTotal())
	fmt.Println("Valor total de Mesa em estoque:", Mesa.valorTotal())
	fmt.Println("Valor total de Lapis em estoque:", Lapis.valorTotal())

}
