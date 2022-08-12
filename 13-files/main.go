package main

import (
	"fmt"
	"os"
)

func readFile() {
	archivo, err := os.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

func saveFile() {
	archivo, err := os.Create("./saved.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Carlos")
	fmt.Fprintln(archivo, "Enrique")
	fmt.Fprintln(archivo, "Santander")
	fmt.Fprintln(archivo, "Ruiz")
	archivo.Close()
}

func main() {
	readFile()

	saveFile()
}
