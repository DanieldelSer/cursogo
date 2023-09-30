package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func MostrarTabla() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero : ")
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	for i := 0; i < 11; i++ {
		fmt.Println(numero * i)
	}

}
