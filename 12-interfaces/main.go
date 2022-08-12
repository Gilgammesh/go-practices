package main

import "fmt"

type human interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	carnivoro() bool
}

type vegetal interface {
	clasificacion() string
}

/* Human Gender */
type genderHuman struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	isMan      bool
}
type man struct{ genderHuman }
type woman struct{ genderHuman }

func (g *genderHuman) respirar() { g.respirando = true }
func (g *genderHuman) pensar()   { g.pensando = true }
func (g *genderHuman) comer()    { g.comiendo = true }
func (g *genderHuman) sexo() string {
	if g.isMan {
		return "Hombre"
	} else {
		return "Mujer"
	}
}

func humanRespirando(h human) {
	h.respirar()
	fmt.Printf("Soy un/una %s, y ya estoy respirando \n", h.sexo())
}

/* Animal Gender */
type genderAnimal struct {
	respirando  bool
	comiendo    bool
	isCarnivoro bool
}
type perro struct{ genderAnimal }
type pato struct{ genderAnimal }

func (g *genderAnimal) respirar()       { g.respirando = true }
func (g *genderAnimal) comer()          { g.comiendo = true }
func (g *genderAnimal) carnivoro() bool { return g.isCarnivoro }

func animalRespirando(a animal) {
	a.respirar()
	fmt.Println("Soy un animal, y estoy respirando")
}

func animalesCarnivoros(a animal) int {
	if a.carnivoro() {
		return 1
	}
	return 0
}

func main() {
	Pablo := new(man)
	Pablo.isMan = true
	humanRespirando(Pablo)

	Maria := new(woman)
	humanRespirando(Maria)

	Dogo := new(perro)
	Dogo.isCarnivoro = true
	animalRespirando(Dogo)

	Lucas := new(pato)
	animalRespirando(Lucas)

	Terrier := new(perro)
	Terrier.isCarnivoro = true
	animalRespirando(Terrier)

	var totalCarnivoros int
	totalCarnivoros += animalesCarnivoros(Dogo)
	totalCarnivoros += animalesCarnivoros(Lucas)
	totalCarnivoros += animalesCarnivoros(Terrier)
	fmt.Printf("Total de animales carnivoros %d \n", totalCarnivoros)
}
