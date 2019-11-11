package main

import (
	"fmt"
	"math/rand"
	"os"
	arg "github.com/Heuristicas/SSS/argumentos"
	mvo "github.com/Heuristicas/SSS/funciones"
)

func main() {
	a, d := arg.Leer_grafica(os.Args[1])
	fmt.Print("ACREDORES: ")
	fmt.Println(a)
	fmt.Print("DEUDORES: ")
	fmt.Println(d)

	seed := 64
	rand.Seed(int64(seed)) // setear la seed a todo el programa
	// mvo.CreaAristas(a,d)
	u := mvo.PrimerUniverso(a,d)
	mvo.CreaAristas(u)
	
}
