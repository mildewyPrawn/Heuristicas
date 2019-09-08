package main

import (
	"fmt"
	"crypto/rand"
	"math/big"
	city "github.com/Heuristicas/TSP/funciones"
)

// Genera un número random
 // Regresa un número entre [0, i)
func randInt(i int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(i)))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return int(n)
}

// Copia un arreglo
// Regresa la copia del arreglo
func copiarCiudades(actual []int) []int {
	nuevo := make([]int, len(actual))
	for i := 0; i < len(actual); i++ {
		nuevo[i] = actual[i]
	}
	return nuevo
}

// Swapea un arreglo
// Recibe dos índices
// Regresa el arreglo con los índices swapeados
func swap(i, j int, actual []int) []int {
	nuevo := copiarCiudades(actual)
	nuevo[j] = actual[i]
	nuevo[i] = actual[j]
	return nuevo;
}

func Vecino(actual []int) ([]int, int, int) {
	i := randInt(len(actual))
	j := randInt(len(actual))
	nuevo := swap(i, j, actual)
	return nuevo, i, j
}

func main() {
	var ciudades40 = []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,
		332,333,489,490,491,492,493,496,652,653,654,656,657,792,815,816,
		817,820,978,979,980,981,982,984}

	c := city.NewCiudades(ciudades40)
	c.TotalAristas()
	c.GetNormalizador()
	c.FunCosto()
	// c.PrintCiudad()

	// copyCiudades := copiarCiudades(ciudades40)
	// fmt.Println(copyCiudades)
	
	// randArray := swap(0,1, ciudades40)
	// fmt.Println(randArray)

	vecinoC40, i, j := Vecino(ciudades40)
	fmt.Printf("%d----%d\n", i, j)
	fmt.Println(vecinoC40)

	v := city.NewCiudades(vecinoC40)
	v.TotalAristas()
	v.GetNormalizador()
	v.FunCosto()
	v.PrintCiudad()
	costoV := v.GetDistTotal()
	fmt.Println("------------------------%2.15f---------------", costoV)
	// fmt.Println(v.costo/v.normalizador)
	
}
