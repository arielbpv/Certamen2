package main

import (
	"fmt"
	"strconv"
)

type comensal struct {
	nombre   string
	comiendo bool
}



func newComensal(numero int, comiendo bool) *comensal {
	aux := strconv.Itoa(numero)
	var c comensal
	c.nombre = "comensal" + aux
	c.comiendo = comiendo
	return &c
}

func corrutina() {

}

func main() {
	asientos := [3]struct
	comen := newComensal(1, true)
	asientos[0] = comen

	fmt.Println(asientos[0])
}
