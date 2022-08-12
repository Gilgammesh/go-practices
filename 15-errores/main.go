package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/* DEFER */
	archivo := "prueba.txt"
	f, err := os.Open(archivo)
	if err != nil {
		fmt.Println("Error abriendo el archivo")
		// os.Exit(1)
	}
	defer f.Close()

	/* PANIC */
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrió un error que generó Panic \n %v \n", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontró el valor de 1")
	}
}
