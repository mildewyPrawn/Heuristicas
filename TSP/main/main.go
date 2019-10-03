package main

import (
	"fmt"
	city "github.com/Heuristicas/TSP/funciones"
	arg "github.com/Heuristicas/TSP/argumentos"
	"os"
	"math/rand"
	// heur "github.com/Heuristicas/TSP/heuristica"
)

func main() {
	if (len(os.Args) < 3) {
		fmt.Println("TSP: Falta el archivo de ciudades D:\nO la semilla")
		os.Exit(1)
	}
	nombre := os.Args[1] // Nombre archivo con ciudades
	ciudades, saludo, seed := arg.Leer(nombre, os.Args[2]) // ciudades y nombre limpio
	fmt.Println(seed)
	rand.Seed(int64(seed))
	fmt.Println(saludo)

	rand.Shuffle(len(ciudades), func(i, j int) {ciudades[i], ciudades[j] = ciudades[j], ciudades[i]})
	
	
	g := city.NewGeneral(ciudades)
	// g.PrintGenData()

	s := city.NewSolucion(ciudades, g)
	s.PrintData()


	tInit := city.TemperaturaInicial(8, s, g)
	// fmt.Printf("TF: %2.15f", tInit)
	// s.PrintData()

	fmt.Println()
	fmt.Println()
	fmt.Println()	
	res := city.AceptacionPorUmbrales(tInit, s, g)
	// s.PrintData()
	fmt.Println(res)
	fmt.Println()
	fmt.Println()
	fmt.Println(tInit)

}
