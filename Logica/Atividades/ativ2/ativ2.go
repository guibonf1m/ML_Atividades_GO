package main

import "fmt"

func main() {
	resultado := 1

	for i := 5; i >= 1; i = i - 1 {
		resultado *= i
	}
	fmt.Println("O Fatorial é:", resultado)

}
