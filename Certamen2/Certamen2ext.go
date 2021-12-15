package main

import "fmt"

type cajeros struct {
	atencionCa *chan int
}

func newCajeros(atencionCa *chan int) *cajeros {
	return &cajeros{atencionCa: atencionCa}
}

func (c *cajeros) atencion() {
	fmt.Println("Nuevo cliente entrando al banco")
	for {
		dirige := <-*c.atencionCa
		fmt.Print("El cliente: ", dirige)
		fmt.Println("se dirige con el cajero")
	}
}

type clientes struct {
	atencionCli *chan int
	terminado   *chan bool
}

func newClientes(atencionCli *chan int, terminado *chan bool) *clientes {
	return &clientes{atencionCli: atencionCli, terminado: terminado}
}

func (cli *clientes) nuevo(max int) {
	fmt.Println("Nueva atención")
	for i := 0; i < max; i++ {
		fmt.Println("El cliente ", i)
		fmt.Println("está siendo atendido")
		*cli.atencionCli <- i
	}
	*cli.terminado <- true
}

func main() {

}

/*Deberá crear un framework de simulación con el que deberá programar un ejemplo de simulación de la atención a clientes en un banco (una cola de clientes con múltiples cajeros). En la figura 3 se aprecia una corrutina generadora de clientes, que permite representar sus llegadas al banco. Por otra parte se tiene c cajeros, a los que atienden a los clientes que están en la fila. Los clientes, una vez atendidos, se retiran del banco.
I)Cajeros cuya pausa en la ejecución permitirá simular el tiempo de atención al cliente
II) Generador de clientes que simulará la creación de clientes, cuya pausa permite simular el tiempo que transcurre entre el arribo de un cliente y otro. Para la ejecución de la simulación se debe iterar de forma cíclica sobre una estructura que contiene referencias a las corrutinas antes mencionadas. En cada iteración se determinará aleatoriamente (con diferente probabilidad) si a la corrutina se la despierta o no, representando de esta forma que los clientes llegan a intervalos diferentes de tiempo y que las atenciones de los cajeros pueden variar en tiempo también.*/
