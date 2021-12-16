package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// structura cajero y creacion
type cajeros struct {
	name       int
	atendiendo int
}

func newCajeros(name int, atendiendo int) *cajeros {
	return &cajeros{name: name, atendiendo: atendiendo}
}

//estructura clientes y creacion
type clientes struct {
	name              int
	TiempoAtencionCli int
}

func newClientes(name int, TiempoAtencionCli int) clientes {
	return clientes{name: name, TiempoAtencionCli: TiempoAtencionCli}
}

//funcion async para la llegada de clientes a la fila
func AsyncFila(nclientes int, clienteSlice []clientes) {
	for {
		if len(clienteSlice) < 20 {
			fmt.Println("Nuevo cliente entrando al banco")
			fmt.Println(len(clienteSlice), cap(clienteSlice))
			var demora int
			demora = rand.Intn(10)
			nuevoClienteSlice := append(clienteSlice, newClientes(nclientes, demora))
			clienteSlice = nuevoClienteSlice
			nclientes++
			time.Sleep(time.Second * 15)
			//			fmt.Println(len(clienteSlice), cap(clienteSlice))
		} else {
			time.Sleep(time.Second * 40)
		}
	}
}

//fmt.Println("entre aca")
//time.Duration(rand.Intn(5))

//atencion a los clientes
func AsyncAtencion(caja []*cajeros, clienteSlice []clientes) {
	//while true
	for {
		var i = 0

		//revisamos que exista un cajero desocupado
		if caja[i].atendiendo == 0 {
			caja[i].atendiendo = 1
			//eliminamos y remplazamos el primer cliente
			var aux clientes
			aux = clienteSlice[0]
			clienteSlice = clienteSlice[1:]

			//lo atiende
			fmt.Printf("El cajero %d, esta atendiendo al cliente %d \n", caja[i].name, aux.name)

			//var demoraCajero = aux.TiempoAtencionCli
			time.Sleep(time.Second * 10)
			caja[i].atendiendo = 0
		}
		i++
		if i > len(caja) {
			i = 0
		}
	}
}

//time.Duration(demoraCajero)

func main() {
	//usuario ingresa la cantidad de cajeros
	fmt.Println("Ingrese la cantidad de cajeros en el programa: ")
	var cajas string
	fmt.Scanln(&cajas)
	icaja, _ := strconv.Atoi(cajas)

	//se crea el arreglo para guardar los cajeros
	var ArrCajas []*cajeros

	//se crean los cajeros
	for i := 0; i < icaja; i++ {
		ArrCajas = append(ArrCajas, newCajeros(i, 0))
	}

	var nclientes = 0
	var clientes = make([]clientes, 2)

	for {
		go AsyncFila(nclientes, clientes)
		go AsyncAtencion(ArrCajas, clientes)
	}

}

//C:\Users\ariel\Desktop\GO\src\github.com\arielbpv\certamen2ext

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
