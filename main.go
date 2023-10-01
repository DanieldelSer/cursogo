package main

import (
	"github.com/cursogo/webserver"
	// "github.com/cursogo/goroutines"
	// "github.com/cursogo/defer_panic"
	// e "github.com/cursogo/ejer_interfaces"
	// "github.com/cursogo/modelos"
	// "github.com/cursogo/users"
	// "github.com/cursogo/mapas"
	// "github.com/cursogo/arreglos_slices"
	// "github.com/cursogo/funciones"
	// "github.com/cursogo/files"
	// "github.com/cursogo/ejercicios"
	// "github.com/cursogo/teclado"
	// "github.com/cursogo/variables"
	// "github.com/cursogo/iteraciones"
)

func main() {
	// estado, texto := variables.ConviertoaTexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)
	// if os := runtime.GOOS; os == "linux" || os == "OS X." {
	// 	fmt.Println("No es windows")
	// } else {
	// 	fmt.Println("Es windows", os)
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	// numero, texto := ejercicios.ConvertirAEntero("101")
	// fmt.Println(numero, texto)
	// fmt.Println(ejercicios.ConvertirAEntero("101"))

	// teclado.IngresosNumeros()

	// iteraciones.Iterar()

	// fmt.Println(ejercicios.MostrarTabla())
	// files.GrabarTabla()

	// files.SumaTabla()
	// files.LeoArchivo()

	// funciones.Calculos()
	// funciones.LlamarClosure()

	// funciones.Exponencia(2)

	// arreglos_slices.Capacidad()

	// mapas.MostrarMapas()

	// users.AltaUsuario()

	// Daniel := new(modelos.Hombre)
	// e.HumanosRespirando(Daniel)

	// Maria := new(modelos.Mujer)
	// e.HumanosRespirando(Maria)

	// defer_panic.VemosDefer()
	// defer_panic.EjemploPanic()

	// canal1 := make(chan bool)
	// go goroutines.MiNombreLento("Daniel del Ser", canal1)
	// defer func() {
	// 	<-canal1
	// }()
	// fmt.Println("Estoy aqui")

	webserver.MiWebServer()
}
