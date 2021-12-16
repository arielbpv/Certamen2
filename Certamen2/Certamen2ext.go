package main

import (
	"fmt"
	"math/rand"
)

type cajeros struct {
	name int
}

func newCajeros(name int) *cajeros {
	return &cajeros{name: name}
}

func (c *cajeros) atencion() {
	fmt.Println("está siendo atendido")
	/*for {
		dirige := <-*c.atencionCa
		fmt.Print("El cliente: ", dirige)
		fmt.Println("se dirige con el cajero")
	}*/

}

type clientes struct {
	name              int
	TiempoAtencionCli int
	terminado         chan bool
}

func newClientes(name int, TiempoAtencionCli int, terminado chan bool) *clientes {
	return &clientes{name: name, TiempoAtencionCli: TiempoAtencionCli, terminado: terminado}
}

func (cli *clientes) nuevo(nclientes int, clienteSlice *[]*clientes, max int) {
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

func main() {
	var nclientes = 0
	var clientes [2]*clientes
	clienteSlice := clientes[0:2]

}

/*
	for i := 0; i < max; i++ {
		fmt.Println("El cliente ", i)
		cli.TiempoAtencionCli=1
	}

	*cli.terminado <- true
*/

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
