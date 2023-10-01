package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["España"] = "Madrid"
	paises["Inglaterra"] = "Londres"
	fmt.Println(paises)
	fmt.Println(paises["España"])

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Sevilla":     37,
		"At.Madrid":   30,
	}

	fmt.Println(campeonato)

	for equipo, puntos := range campeonato {
		fmt.Printf("Equipo %s, tiene puntos: %d \n", equipo, puntos)
	}

	delete(campeonato, "Barcelona")
	fmt.Println(campeonato)

	puntos, existe := campeonato["Sevilla"]
	fmt.Printf("Los puntos son %d, y el equipo existe = %t\n", puntos, existe)
}
