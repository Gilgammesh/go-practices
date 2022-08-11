package main

import "fmt"

func main() {
	count := 5
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln(&numero)
		if numero == 55 {
			break
		}
	}

	var i = 0
	for i < 10 {
		fmt.Printf("\n Valor de i: %d", i)
		if i == 5 {
			fmt.Printf("    multiplicamos por 2 \n")
			i = i * 2
			continue // va a ir a la instruccion for
		}
		fmt.Printf("    Pasó por aquí")
		i++
	}

	var j int = 0
	var text string = "primera vez"

RUTINA:
	fmt.Printf("Inicia RUTINA %s \n", text)
	for j < 10 {
		if j == 4 {
			j = j + 3
			fmt.Println("Voy a RUTINA sumando 3 a j")
			text = "sumando 3 a j"
			goto RUTINA // va a ir a la RUTINA encima del for
		}
		fmt.Printf("Valor de j: %d \n", j)
		j++
	}
}
