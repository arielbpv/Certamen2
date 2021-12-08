package main

import (
	"fmt"
	"time"
)

type mesa struct {
	comensales [5]comensal
}

type comensal struct {
	nombre string
	canal  chan bool
}

func newMesa(comensales [5]comensal) *mesa {
	Corr := mesa{comensales: comensales}
	return &Corr
}

func (Corr *mesa) reposar(C int) {
	time.Sleep(2 * time.Second)
	fmt.Println(Corr.comensales[C].nombre, "solto los cubiertos")
	Corr.comensales[C].canal <- true
}

func (Corr *mesa) comer(C int) {
	for i := 0; ; i++ {
		fmt.Println(Corr.comensales[C].nombre, "comiendo")
		<-Corr.comensales[C].canal
	}
}

func main() {
	chanel := make(chan bool)
	veces := 2
	comensal1 := comensal{nombre: "comensal 1", canal: chanel}
	comensal2 := comensal{nombre: "comensal 2", canal: chanel}
	comensal3 := comensal{nombre: "comensal 3", canal: chanel}
	comensal4 := comensal{nombre: "comensal 4", canal: chanel}
	comensal5 := comensal{nombre: "comensal 5", canal: chanel}

	comida := newMesa([5]comensal{comensal1, comensal2, comensal3, comensal4, comensal5})

	go comida.comer(0)
	go comida.comer(1)
	go comida.comer(2)
	go comida.comer(3)
	go comida.comer(4)

	for i := 0; i <= veces; i++ {
		comida.reposar(0)
		comida.reposar(1)
		comida.reposar(2)
		comida.reposar(3)
		comida.reposar(4)
	}
}
