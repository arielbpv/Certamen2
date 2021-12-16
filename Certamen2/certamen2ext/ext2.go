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

func newCajeros(name int, atendiendo int) cajeros {
	return cajeros{name: name, atendiendo: atendiendo}
}

//estructura clientes y creacion
type clientes struct {
	name              int
	TiempoAtencionCli int
}

func newClientes(name int, TiempoAtencionCli int) clientes {
	return clientes{name: name, TiempoAtencionCli: TiempoAtencionCli}
}
/*
func adding(clienteSlice []clientes, nclientes int) chan int{
	r := make(chan int)

	go func(){
		defer close(r)
		fmt.Println("Nuevo cliente entrando al banco")
		var demora int
		demora = rand.Intn(10)
		clienteSlice = append(clienteSlice, newClientes(nclientes, demora))
		//clienteSlice = &nuevoClienteSlice
		nclientes++
		time.Sleep(time.Second * 15)
		r <- rand.Intn(10)
	}()
	return r
}

//funcion async para la llegada de clientes a la fila
//funcion async para la llegada de clientes a la fila
func AsyncFila(nclientes int, clienteSlice []clientes) <-chan int {
    //r := make(chan int)

   for i:=0; i<2; i++{
        if len(clienteSlice) < 20 {
            go adding(clienteSlice, nclientes)
            fmt.Println(len(clienteSlice), cap(clienteSlice))
			//fmt.Println(<-val)
        }else{
			fmt.Println(clienteSlice)
		}
    }
	//return r
}*/

//funcion async para la llegada de clientes a la fila
func AsyncFila(nclientes int, clienteSlice []clientes) <-chan int {
    r := make(chan int)

    //for {
        if len(clienteSlice) <= 20 {
            go func() {
                defer close(r)

                fmt.Println("Nuevo cliente entrando al banco")
                var demora int
                demora = rand.Intn(10)
                nuevoClienteSlice := append(clienteSlice, newClientes(nclientes, demora))
                clienteSlice = nuevoClienteSlice
                nclientes++
                time.Sleep(time.Second * 15)
                r <- rand.Intn(10)
            }()

            fmt.Println(len(clienteSlice), cap(clienteSlice))
        }
		return r
   // }
}


//fmt.Println("entre aca")
//time.Duration(rand.Intn(5))

//el primer cliente en la fila se atiende, avanza la fila
func popQueue(clienteSlice []clientes) clientes{
	var aux clientes
	aux = clienteSlice[0]
	clienteSlice = clienteSlice[1:]
	fmt.Printf("El cliente %d sale de la fila\n",aux.name)
	return aux
}

func atent(cajero *cajeros, clienteSlice []clientes){
	if cajero.atendiendo == 1 {
		var aux = popQueue(clienteSlice)
		cajero.atendiendo = 0
		//lo atiende
		fmt.Printf("El cajero %d, esta atendiendo al cliente %d ", cajero.name, aux.name)
		//var demoraCajero = aux.TiempoAtencionCli
		time.Sleep(time.Second * 20)
		cajero.atendiendo = 1
		fmt.Printf("\t Cliente satisfecho\n")
	}/*else{
		time.Sleep(time.Second * 14)
	}*/
}

//atencion a los clientes
func AsyncAtencion(caja []cajeros, clienteSlice []clientes) {
	//while true
	for j:=0 ; j<200; j++{
		for i := 0; i < len(caja); i++ {
			atent(&caja[i], clienteSlice)
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
	var ArrCajas = make([]cajeros, 1)

	//se crean los cajeros
	for i := 1; i <= icaja; i++ {
		ArrCajas = append(ArrCajas, newCajeros(i, 1))
		fmt.Printf("Se crea el cajero %d con estado %d\n",ArrCajas[i].name, ArrCajas[i].atendiendo)
	}

	var nclientes int = 1
	var clientes = make([]clientes, 1)

	for j:=0; j<100; j++{
		go AsyncFila(nclientes, clientes)
		go AsyncAtencion(ArrCajas, clientes)
		fmt.Println(clientes)
	}
	fmt.Println(ArrCajas)
	fmt.Println(clientes)
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
