package main

import (
	"fmt"
	"strconv"
	"time"
)

var procesoMatar int = 20

func mostrar_opciones() string {
	var opcion string
	fmt.Println("Selecciona una opción")
	fmt.Println("1) Agregar proceso")
	fmt.Println("2) Mostrar procesos")
	fmt.Println("3) Terminar proceso")
	fmt.Println("0) Salir")
	fmt.Scan(&opcion)
	return opcion
}

func Proceso(ch chan string, id int) {
	i := uint64(0)
	for {
		if procesoMatar == id {
			return
		}
		ch <- strconv.Itoa(id) + ":" + strconv.FormatUint(i, 10)
		i = i + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func eliminar() int {
	var opcion string
	fmt.Print("Id a eliminar? ")
	fmt.Scan(&opcion)
	value, err := strconv.Atoi(opcion)
	if err == nil {
		return value
	}
	return 0
}

func main() {
	c := make(chan string)
	var opcion string
	idActual := 0
	imprimir := false

	go func() {
		for {
			msg := <-c
			if imprimir {
				fmt.Println(msg)
			}
		}
	}()

	for {
		opcion = mostrar_opciones()
		switch opcion {
		case "1":
			idActual++
			go Proceso(c, idActual)
		case "2":
			imprimir = true
			var scan string
			fmt.Scanf("%s", &scan)
			fmt.Scanf("%s", &scan)
			imprimir = false
		case "3":
			procesoMatar = eliminar()
		case "0":
			return
		default:
			fmt.Println("Opción incorrecta")
		}
	}
}
