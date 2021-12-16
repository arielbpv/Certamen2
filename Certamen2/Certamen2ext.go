package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cajeros struct {
	name int
}

func newCajeros(name int) *cajeros {
	return &cajeros{name: name}
}

func (c *cajeros) atencion() {
	fmt.Println("está siendo atendido")

}

type clientes struct {
	name              int
	TiempoAtencionCli int
	terminado         chan bool
}

func newClientes(name int, TiempoAtencionCli int, terminado chan bool) *clientes {
	return &clientes{name: name, TiempoAtencionCli: TiempoAtencionCli, terminado: terminado}
}

func nuevo(nclientes int, clienteSlice *[]*clientes) {
	time.Sleep(3 * time.Second)
	var personasNew = rand.Intn(2)
	for i := 0; i < personasNew; i++ {
		if len(*clienteSlice) < 20 {
			fmt.Println("Nuevo cliente entrando al banco")
			var demora = rand.Intn(10)
			var terminar = make(chan bool)
			*clienteSlice = append(*clienteSlice, newClientes(nclientes, demora, terminar))
			nclientes = nclientes + 1
		}
	}
}
/*
func asyncFila(nclientes int, clienteSlice *[]*clientes) chan int{
	r := make(chan int)
	var slice []*clientes = *clienteSlice
	var id int = nclientes
	fmt.Println("Entrando a la gorutina anonima")
	
	
	go func() {
		//var randomTime int = 2*rand.Intn(10)
		defer close(r)
		fmt.Println("Nuevo cliente entrando al banco")
		var demora = rand.Intn(10)
		var terminar = make(chan bool)
		*slice = append(*slice, newClientes(id, demora, terminar))
		id = id + 1
		time.Sleep(15 * time.Second)
	}()

	fmt.Println("Termina una gorutina")
	
	return r
}*/

func main() {
	var nclientes = 1
	var clientes [2]*clientes
	clientes[0] =newClientes(nclientes, rand.Intn(10), make(chan bool))
	nclientes = nclientes + 1
	clientes[1] =newClientes(nclientes, rand.Intn(10), make(chan bool))
	nclientes = nclientes + 1
	clienteSlice := clientes[0:2]

	nuevo(nclientes, *clienteSlice)

	/*for true{
		fmt.Println("Itera el while true")
		nuevo(nclientes, clienteSlice)
	}*/
}


/*Deberá crear un framework de simulación con el que deberá programar un ejemplo de simulación
de la atención a clientes en un banco (una cola de clientes con múltiples cajeros). En la
figura 3 se aprecia una corrutina generadora de clientes, que permite representar sus llegadas
al banco. Por otra parte se tiene c cajeros, a los que atienden a los clientes que están en la fila. Los clientes, una vez atendidos, se retiran del banco.
I)Cajeros cuya pausa en la ejecución permitirá simular el tiempo de atención al cliente
II) Generador de clientes que simulará la creación de clientes, cuya pausa permite simular el
tiempo que transcurre entre el arribo de un cliente y otro. Para la ejecución de la simulación
se debe iterar de forma cíclica sobre una estructura que contiene referencias a las corrutinas
antes mencionadas. En cada iteración se determinará aleatoriamente (con diferente probabilidad)
si a la corrutina se la despierta o no, representando de esta forma que los clientes llegan a
intervalos diferentes de tiempo y que las atenciones de los cajeros pueden variar en tiempo
también.*/
