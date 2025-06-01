package main

import (
	"fmt"
	"strconv"
)

func main() {
	//
	//atividade 1

	var n int
	fmt.Print("Enter a number N: ")
	fmt.Scanln(&n)
	//numeroPar := 10
	SomaTotal := 0
	soma := ""
	simbolo := "+"
	nome := "Oscar"
	sobrenome := "Salomon"
	idade := 35
	info := nome + " " + sobrenome + " " + strconv.Itoa(idade)

	fmt.Println(info)

	for i := 2; i <= n; i += 2 {
		SomaTotal += i
		if i == n {
			simbolo = "="
		}
		soma = soma + strconv.Itoa(i) + simbolo
	}
	fmt.Printf("Soma total: %s %d", soma, SomaTotal)
}
