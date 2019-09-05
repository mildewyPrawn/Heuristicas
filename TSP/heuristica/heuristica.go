// package heuristica
package main

import (
	"fmt"
	"crypto/rand"
	"math/big"
	"github.com/Heuristicas/TSP/funciones"
)
// var mejor []int // Mejor solucion
type Solucion struct {
	Ciudades []int
	Funcion float64
}

// var Actual []int // Solucion actual
// var Nueva []int // Solución nueva
// var epsilon = 

// Genera un int entre [0, 27)
func randInt(i int) int {
	// Usamos esto porque el otro random siempre da los mismos vlv
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(i)))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return int(n);
}

// Copia un arreglo y regresa la copia
func Copy(actual []int) []int {
	nuevo := make([]int, len(actual))
	for i := 0; i < len(actual); i++ {
		nuevo[i] = actual[i]
	}
	return nuevo;
}

// Recibe los índices para swapear
// Copiamos a  un nuevo arreglo
// Lo regresamos (el nuevo) y actualizamos 'nueva'
func swap(i, j int, actual []int) []int {
	nuevo := Copy(actual)
	nuevo[j] = actual[i]
	nuevo[i] = actual[j]
	// Nueva = nuevo
	return nuevo
}

// Saca una permutacion de dos elementos de la solucion actual
func Vecino(actual []int) []int{
	i := randInt(len(actual))
	j := randInt(len(actual))
	fmt.Println(i)
	fmt.Println(j)	
	nuevo := swap(i, j, actual)
	return nuevo
}

func main() {
	
	var ciudades40 = []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,
		332,333,489,490,491,492,493,496,652,653,654,656,657,792,815,816,
		817,820,978,979,980,981,982,984}

	funciones.Init(ciudades40)
	actual := Solucion{Ciudades: ciudades40,
		Funcion: funciones.FunCosto(ciudades40)}
	nueva := Solucion{Ciudades: Vecino(ciudades40),
		Funcion: 13.1}
	fmt.Println(actual)
	fmt.Println(nueva)

	// Vecino();
	// fmt.Println(Actual)
	// fmt.Println(Nueva)

}
