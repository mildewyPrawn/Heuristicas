package main

import (
	"fmt"
	city "github.com/Heuristicas/TSP/funciones"
	// arg "github.com/Heuristicas/TSP/argumentos"
	"os"
	"math/rand"
)

func  main() {  

	if(len(os.Args) < 3) {
		fmt.Println("TSP: falta el archivo de ciudades D:")
		fmt.Println("O la semilla...")
		os.Exit(1)
	}
	// nombre := os.Args[1] // nombre del archivo con ciudades
	// ciudades, _, _ := arg.Leer(nombre, os.Args[2]) // ciduades y seed

	ciudades := []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,332,333,489,490,491,492,493,496,652,653,654,656,657,792,815,816,817,820,978,979,980,981,982,984}

	seed := 0
	fBest := 1000000.0
	fSeed := 0
	// for i := 1100; i < 1650; i++ { // hacemos este ciclo para usar muchas ciudades
		// basicamente nos vale el arguemnto 2, pero cuando ya tenga la
		// mejor debería borrar el for
		
	seed = 326
		rand.Seed(int64(seed)) // setear la seed a todo el programa
		fmt.Println()
		fmt.Printf("SEED: %d\n", seed)
		cities := make([]int, len(ciudades))
		for i := 0; i < len(cities); i++ {
			cities[i] = ciudades[i]
		}
		//revolver la solucion
		rand.Shuffle(len(cities), func(i, j int) {cities[i], cities[j] = cities[j], cities[i]})
		current := city.NewTSP(cities)
		fmt.Println()
		if current < fBest {
			// fSeed = i
			fSeed = seed
			fBest = current
		}
	// }
	fmt.Printf("BEST[%d]:%2.15f", fSeed, fBest)
}
