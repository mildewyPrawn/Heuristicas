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
	rand.Seed(int64(seed))
	fmt.Println(saludo)
	
	c := city.NewCiudades(ciudades)	
	c.FunCosto() //Esta así y no desde el "constructor" para poder seguirla calculando
	c.PrintCiudad()
	

	s := city.NewSolucion(c.AristasE, c.Id, c.Distancias)


	
	
	tInit := c.TemperaturaInicial(8, s)
	fmt.Printf("TF: %2.15f", tInit)
	

	fmt.Println()
	fmt.Println()
	fmt.Println()	
	res := c.AceptacionPorUmbrales(tInit, s)
	s.PrintData()
	fmt.Println(res)
	fmt.Println()
	fmt.Println()
	fmt.Println(tInit)
	// s.PrintData(c)
	// fmt.Printf("\n%2.15f\n",s.nuevaC)
	

	
	// min := city.FunCostoSolucion(res, c.GetDistancias(), c.GetAristasE())
	// fmt.Printf("EL COSTO ES DE: %2.15f", min/c.GetNormalizador())


	// [164 331 980 817 491 6 333 2 978 496 4 820 653 489 982 984 657 3 332 172 816 490 329 163 652 493 979 815 492 165 75 656 5 1 792 168 327 981 654 7]

	// [492 653 490 489 493 332 164 815 6 3 5 165 2 327 980 656 331 816 817 984 657 333 978 329 491 792 981 654 4 496 75 982 652 820 163 1 172 7 979 168]
                                                       



}
