package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func MostrarTabla() string {
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 11; i++ {
		fmt.Println("Ingrese numero : ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i < 11; i++ {
		// fmt.Println(numero * i)
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
	}

	return texto
}
