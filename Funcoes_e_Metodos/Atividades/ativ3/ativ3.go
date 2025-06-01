package main

import "fmt"

type Estudante struct {
	Nome  string
	Idade int
}

type Turma struct {
	Alunos []Estudante
}

func (e Estudante) Exibir() string {
	msg := fmt.Sprintf("Nome: %s, Idade %d: ", e.Nome, e.Idade)
	return msg
}

func (n Turma) NomeDuplicado(nomeRecebido string) bool {
	for _, aluno := range n.Alunos {
		if aluno.Nome == nomeRecebido {
			return true
		}
	}
	return false
}

func main() {
	var quantidade int
	var alunosTurma Turma

	fmt.Println("Qual a quantidade de alunos da turma: ")
	fmt.Scanln(&quantidade)

	for i := 0; i < quantidade; i++ {
		var nome string
		var idade int

		for {
			fmt.Println("Qual é seu nome: ")
			fmt.Scanln(&nome)
			duplicado := alunosTurma.NomeDuplicado(nome)

			if !duplicado {
				fmt.Println("Aluno cadastrado com sucesso!")
				break
			} else {
				fmt.Println("Aluno já existe, tente outro nome!")
			}
		}
		fmt.Println("Qual é a sua idade: ")
		fmt.Scanln(&idade)

		novoAluno := Estudante{Nome: nome, Idade: idade}
		alunosTurma.Alunos = append(alunosTurma.Alunos, novoAluno)

	}

	for _, aluno := range alunosTurma.Alunos {
		msg := aluno.Exibir()
		fmt.Println(msg)

	}
}
