package main

import (
	"bufio"
	"fmt"
	"os"
)

var number1 int
var number2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("Enter number 1:")
	// fmt.Scanf("%d", &number1)
	fmt.Scanln(&number1)

	fmt.Println("Enter number 2:")
	// fmt.Scanf("%d", &number2)
	fmt.Scanln(&number2)

	// fmt.Println("Suma:", number1+number2)
	fmt.Println("Leyenda:")

	// Using bufio
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	resultado = number1 + number2
	fmt.Println(leyenda, resultado)

}
