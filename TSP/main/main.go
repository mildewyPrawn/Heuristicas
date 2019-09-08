package main

import (
	"fmt"
	city "github.com/Heuristicas/TSP/funciones"
	arg "github.com/Heuristicas/TSP/argumentos"
	"os"
	// heur "github.com/Heuristicas/TSP/heuristica"
)

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("TSP: Falta el archivo de ciudades D:")
		os.Exit(1)
	}
	nombre := os.Args[1] // Nombre archivo con ciudades
	ciudades, saludo := arg.Leer(nombre) // ciudades y nombre limpio
	fmt.Println(saludo)
	
	c := city.NewCiudades(ciudades)	
	c.FunCosto() //Esta así y no desde el "constructor" para poder seguirla calculando
	c.PrintCiudad()	
}
