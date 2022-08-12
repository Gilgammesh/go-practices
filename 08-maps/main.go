package main

import "fmt"

func main() {
	paises := make(map[int]string, 5)

	paises[1] = "Argentina"
	paises[2] = "Perú"
	paises[3] = "México"

	fmt.Println(paises)

	campeonatos := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  25,
		"Chivas":       14,
		"Boca Juniors": 14,
	}

	campeonatos["River Plate"] = 43
	campeonatos["Chivas"] = 55

	delete(campeonatos, "Real Madrid")

	fmt.Println(campeonatos)

	for equipo, puntaje := range campeonatos {
		fmt.Printf("El equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	puntaje1, existe1 := campeonatos["Alianza"]
	fmt.Println(puntaje1, existe1)

	puntaje2, existe2 := campeonatos["Chivas"]
	fmt.Println(puntaje2, existe2)
}
