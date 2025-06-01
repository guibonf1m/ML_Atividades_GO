package main

import "fmt"

func main() {

	numero1 := 48
	numero2 := 18
	resto := numero1 % numero2

	for resto != 0 {
		numero1 = numero2
		numero2 = resto
		resto = numero1 % numero2
	}
	fmt.Println("O MDC é:", numero2)
}
