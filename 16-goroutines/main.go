package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	name := "Carlos"
	go showSlowName(name)
	fmt.Println("Estoy aquí")
	var x string
	fmt.Scanln(&x)
}

func showSlowName(name string) {
	letras := strings.Split(name, "")
	fmt.Println(letras)
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
