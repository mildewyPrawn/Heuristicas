package main

import (
	arg "github.com/Heuristicas/SSS/argumentos"
	"os"
	"fmt"
)

func main() {
	matriz := arg.Leer_grafica(os.Args[1])
	fmt.Println(matriz)
}
