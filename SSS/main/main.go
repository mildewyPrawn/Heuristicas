package main

import (
	arg "github.com/Heuristicas/SSS/argumentos"
	"os"
	"fmt"
)

func main() {
	a, d := arg.Leer_grafica(os.Args[1])
	fmt.Print("ACREDORES: ")
	fmt.Println(a)
	fmt.Print("DEUDORES: ")
	fmt.Println(d)
}
