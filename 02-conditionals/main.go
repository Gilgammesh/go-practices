package main

import "fmt"

var estado bool

func main() {
	// estado = true
	if estado = true; estado {
		fmt.Println("El estado es verdadero")
	} else {
		fmt.Println("El estado es falso")
	}

	// numero := 5
	switch numero := 7; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Mayor que 5")
	}

}
