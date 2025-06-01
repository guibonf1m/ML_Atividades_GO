package main

import "fmt"

func main() {
	
	a := 0
	b := 1

	for b <= 10 {
		fmt.Println(b)
		proximo := a + b
		a = b
		b = proximo
	}
}
