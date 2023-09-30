package ejercicios

import "strconv"

func ConvertirAEntero(nombre string) (int, string) {
	if valor, err := strconv.Atoi(nombre); err != nil {
		return 0, "Ha habido un error"
	} else if valor > 100 {
		return valor, "Es mayor a 100"
	} else {
		return valor, "Es menor a 100"
	}
}
