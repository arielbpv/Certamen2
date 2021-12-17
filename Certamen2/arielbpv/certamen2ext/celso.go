package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Generador struct {
	ch chan int
}
type Cajero struct {
	ch      chan int
	cliente Cliente
}

type Cliente struct {
	nombre int
}

type Cola struct {
	cola    [25]Cliente
	final   int
	counter int
}

//funciones de la Cola
func Insertar_cola(cl Cliente) {
	if cola.final < len(cola.cola) {
		cola.cola[cola.final] = cl
		cola.final++
		cola.counter++
	} else {
		fmt.Println("Cola Llena")
	}
}

func Desencolar() Cliente {
	var cl Cliente
	if cola.final != 0 {
		var queqe [len(cola.cola)]Cliente //cola auxiliar
		cl = cola.cola[0]                 //Cliente a sacar de la cola (primero)
		for i := 0; i < cola.final-1; i++ {
			queqe[i] = cola.cola[i+1]
		}
		cola.cola = queqe
		cola.final--
	}
	return cl
}

//corrutina de los cajeros
func (this Cajero) Atender() {
	for {
		var cliente Cliente
		if this.cliente == cliente {
			this.cliente = Desencolar()
			fmt.Printf("%d", this.cliente.nombre)
			this.Active() //acvtivo el main
			this.Hold()   //pauso el atender
		} else {
			fmt.Printf("%d", this.cliente.nombre)
			if rand.Intn(100) < 20 {
				var cliente_nulo Cliente
				this.cliente = cliente_nulo
			}
			this.Active() //activo main
			this.Hold()   //pauso atender
		}
	}
}

//Corrutina del generador de clientes
func (g Generador) generar_clientes() {
	for {
		cliente := Cliente{cola.counter}
		Insertar_cola(cliente)
		g.Active() //se activa el main
		g.Hold()   //se para el generador

	}
}

//Hold y Active de los cajeros
func (c Cajero) Hold() {
	<-c.ch
}
func (c Cajero) Active() {
	time.Sleep(1000 * time.Millisecond)
	c.ch <- 0
}

//Hold y Active de los cajeros
func (g Generador) Hold() {
	<-g.ch
}
func (g Generador) Active() {
	time.Sleep(1000 * time.Millisecond)
	g.ch <- 0
}

//inicializacion de la cola
var clientes [25]Cliente
var cola = Cola{clientes, 0, 1}

func main() {
	//incializacion de los cajeros
	n_cajeros := flag.Int("cajeros", 1, "la cantidad de cajeros")
	flag.Parse()
	var cajeros []Cajero
	for i := 0; i < *n_cajeros; i++ {
		var cajero Cajero
		cajero.ch = make(chan int)
		cajeros = append(cajeros, cajero)
	}
	//inicializacion del generador
	ch_cliente := make(chan int)
	generador := Generador{ch_cliente}
	//comienzo de las corrutinas
	go generador.generar_clientes()
	generador.Hold() //se para el main
	fmt.Println("Cola: ", cola.cola)
	for i := 0; i < len(cajeros); i++ { //itera la cantidad de cajeros
		fmt.Printf("Cajero %d: atiende al cliente ", i+1)
		go cajeros[i].Atender()
		cajeros[i].Hold() //pauso el main
		fmt.Println("\nCola: ", cola.cola)
	}
	//
	for {
		generador.Active() //activamos el generador
		generador.Hold()   //pausamos el main
		fmt.Println("\nCola: ", cola.cola)
		for i := 0; i < len(cajeros); i++ {
			fmt.Printf("Cajero %d: atiende al cliente ", i+1)
			cajeros[i].Active() //activo atender
			cajeros[i].Hold()   //pauso main
			fmt.Println("\nCola: ", cola.cola)
		}
	}
}
