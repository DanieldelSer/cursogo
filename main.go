package main

import (
	"fmt"
	"runtime"
	// "github.com/cursogo/variables"
)

func main() {
	// estado, texto := variables.ConviertoaTexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("No es windows")
	} else {
		fmt.Println("Es windows", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
