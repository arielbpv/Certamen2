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

func newCajero(name int, terminado chan bool) cajero {
	return cajero{name: name, terminado: terminado}
}

//clientes
type cliente struct {
	name              int
	TiempoAtencionCli int
}

func newCliente(name int, TiempoAtencionCli int, terminado chan bool) cliente {
	return cliente{name: name, TiempoAtencionCli: TiempoAtencionCli}
}

//banco
type banco struct {
	clientes []cliente
	cajeros  []cajero
	canal    chan bool
	last     int
	ncliente int
}

func (banco *banco) getNcliente() int{
	return banco.ncliente
}
func newBanco(cajeros []cajero, clientes []cliente) *banco{
	return &banco{cajeros: cajeros, clientes: clientes}
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
func (banco *banco) addCajero(cajero cajero){
	//banco.cajas[len(banco.cajas)-1] = cajero;
  banco.cajeros = append(banco.cajeros, cajero)
  fmt.Println("Crea al cajero ",cajero.name)
}

func (b *banco) CrearCajeros(Ncajas int) {
	for i := 0; i < Ncajas; i++ {
		var aux cajero = newCajero(i+1,make(chan bool))
		b.addCajero(aux)
	}
	b.WakeUp()
}
func (b *banco) CrearClientes() {
	for {
		b.encola()
		b.WakeUp()
		b.sleep()
	}
}
func (b *banco) desencola() cliente {
	var salida cliente
	salida = b.clientes[0]
	b.clientes = b.clientes[1:]
	fmt.Printf("El cliente %d sale de la fila\n", salida.name)
	return salida

}

func (banco *banco) addCliente(cliente cliente){

	banco.clientes = append(banco.clientes, cliente)
	banco.ncliente = banco.ncliente + 1
	fmt.Println("Crea al cliente ",cliente.name)

}

func (b *banco) encola() {
	if(b.ncliente != 0){
		var aux cliente = newCliente(b.ncliente, rand.Intn(10)+1, make(chan bool))
   		b.addCliente(aux)
	}else{
		b.ncliente = b.ncliente+1
	}
}

func (c *cajero) Atender(b *banco) {
	c.atendiendo = b.desencola()
	if c.name != 0{
		if c.atendiendo.name != 0{
			fmt.Printf("el cajero %d atendiendo al cliente %d \n", c.name, c.atendiendo.name)
	
			for i := c.atendiendo.TiempoAtencionCli; i > 0; i-- {
				if c.atendiendo.TiempoAtencionCli > 0 {
					time.Sleep(time.Second * 1)
					c.atendiendo.TiempoAtencionCli = c.atendiendo.TiempoAtencionCli - 1
				}
			}
			fmt.Printf("El cliente %d fue atendido y se retira del banco\n", c.atendiendo.name)
		}
	}
	b.WakeUp()
	c.sleep()
}


func main() {
	nCajeros := flag.Int("CrearCajero", 1, "NumCajeros")
	flag.Parse()

	CajerosArray := make([]cajero,1)
	SliceCajeros := CajerosArray
	ClientesArray := make([]cliente,1)
	SliceClientes := ClientesArray

	banco := newBanco(SliceCajeros, SliceClientes)
	banco.canal = make(chan bool)

	go banco.CrearCajeros(*nCajeros)
	banco.sleep()
	go banco.CrearClientes()
	banco.sleep()

	for {
		for i := 0; i < len(banco.cajeros); i++ {
			go banco.cajeros[i].Atender(banco)
			banco.sleep()
		}
		banco.WakeUp()
		banco.sleep()
	}
}
