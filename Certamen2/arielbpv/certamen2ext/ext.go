package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

//cajeros
type cajero struct {
	name       int
	atendiendo cliente
	terminado  chan bool
}

func newCajero(name int, terminado chan bool) *cajero {
	return &cajero{name: name, terminado: terminado}
}

//clientes
type cliente struct {
	name              int
	TiempoAtencionCli int
}

func newCliente(name int, TiempoAtencionCli int, terminado chan bool) *cliente {
	return &cliente{name: name, TiempoAtencionCli: TiempoAtencionCli}
}

//banco
type banco struct {
	clientes [20]cliente
	cajeros  []cajero
	canal    chan bool
	last     int
	ncliente int
}

func (b *banco) sleep() {
	<-b.canal

}
func (b *banco) WakeUp() {
	time.Sleep(1 * time.Second)
	b.canal <- true
}
func (c *cajero) sleep() {
	<-c.terminado

}
func (c *cajero) WakeUp() {
	time.Sleep(1 * time.Second)
	c.terminado <- true
}

//funciones
func (b *banco) CrearCajeros(Ncajas int) {
	for i := 0; i < Ncajas; i++ {
		cajero := newCajero(i, make(chan bool))
		b.cajeros = append(b.cajeros, *cajero)
		fmt.Printf("se creo el cajero %d \n", b.cajeros[i].name)
	}
	b.WakeUp()
}
func (b *banco) CrearClientes() {
	for {
		if b.last != -1 {
			var demora = rand.Intn(5)
			cliente := newCliente(b.ncliente, demora, make(chan bool))
			if b.last < 20 {
				b.clientes[b.last] = *cliente
				b.last++
				b.ncliente++
				fmt.Printf("se creo el cliente %d \n", b.clientes[b.last-1].name)
			}
			b.WakeUp()
			b.sleep()
		} else {
			var demora = rand.Intn(10)
			cliente := newCliente(b.ncliente, demora, make(chan bool))
			b.clientes[0] = *cliente
			b.last = 1
			b.ncliente = 1
		}
	}

}
func (b *banco) desencola() cliente {
	var salida cliente
	if b.last != (-1) {
		for i := 0; i < b.last-1; i++ {
			if i == 0 {
				salida = b.clientes[i]
			}
			b.clientes[i] = b.clientes[i+1]
		}
		b.last--
	}
	return salida
}
func (b *banco) encola() {
	if b.last < len(b.clientes) {
		var demora = rand.Intn(10)
		cliente := newCliente(b.ncliente, demora, make(chan bool))
		b.clientes[b.last] = *cliente
		b.last++
		b.ncliente++
	}
}

func (b *banco) Atender() {
	fmt.Println("entre a atender")
	for {
		for i := 0; i < len(b.cajeros); i++ {
			var client cliente

			if b.cajeros[i].atendiendo == client {
				b.cajeros[i].atendiendo = b.desencola()
				fmt.Printf("el cajero %d atendiendo al cliente %d \n", b.cajeros[i].name, b.cajeros[i].atendiendo.name)
				b.WakeUp()
				b.cajeros[i].sleep()
			} else {

				if b.cajeros[i].atendiendo.TiempoAtencionCli > 0 {
					fmt.Printf("restandole tiempo al cliente %d \n", b.cajeros[i].atendiendo.name)
					b.cajeros[i].atendiendo.TiempoAtencionCli--
				} else {
					b.cajeros[i].atendiendo = client

				}
				b.WakeUp()
				b.cajeros[i].sleep()
			}
		}
	}
}

func main() {
	nCajeros := flag.Int("CrearCajero", 1, "NumCajeros")
	flag.Parse()
	var banco banco
	banco.canal = make(chan bool)

	go banco.CrearCajeros(*nCajeros)
	banco.sleep()
	go banco.CrearClientes()
	banco.sleep()

	for {
		for i := 0; i < len(banco.cajeros); i++ {
			go banco.Atender()
			banco.sleep()
		}
		banco.WakeUp()
		banco.sleep()
	}
}
