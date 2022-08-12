package main

import "fmt"

func main() {
	var numero int = 2
	exponencia(numero)
}

func exponencia(numero int) {
	if numero > 100000000 {
		return
	}
	fmt.Println(numero)
	exponencia(numero * 2)
}
