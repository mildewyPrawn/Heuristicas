package main

import (
	"fmt"
	city "github.com/Heuristicas/TSP/funciones"
	arg "github.com/Heuristicas/TSP/argumentos"
	"os"
	"math/rand"
)

func  main() {  

	if(len(os.Args) < 3) {
		fmt.Println("TSP: falta el archivo de ciudades D:")
		fmt.Println("O la semilla...")
		os.Exit(1)
	}
	nombre := os.Args[1] // nombre del archivo con ciudades
	ciudades, _, _ := arg.Leer(nombre, os.Args[2]) // ciduades y seed
	
	for i := 0; i < 1001; i++ { // hacemos este ciclo para usar muchas ciudades
		// basicamente nos vale el arguemnto 2, pero cuando ya tenga la
		// mejor debería borrar el for
		seed := i
		rand.Seed(int64(seed)) // setear la seed a todo el programa
		fmt.Println()
		fmt.Printf("SEED: %d\n", seed)
		//revolver la solucion
		rand.Shuffle(len(ciudades), func(i, j int) {ciudades[i], ciudades[j] = ciudades[j], ciudades[i]})
		city.NewTSP(ciudades)
		fmt.Println()
	}
}
