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
	Distancia float64
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
func Vecino(actual []int) ([]int, int, int){
	i := randInt(len(actual))
	j := randInt(len(actual))
	// fmt.Println(i)
	// fmt.Println(j)
	nuevo := swap(i, j, actual)
	return nuevo, i, j
}

// distTotal es el total de la suma de la funcion de costo antes de dividirla
// func Actualiza(vieja, nueva []int, i, j int, distTotal float64) float64{
func Actualiza(vieja []int, i, j int, distTotal float64) float64{
	// fmt.Println(funciones.Distancias[i][j])
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(vieja)
	// fmt.Println(nueva)
	if (i == 0 && j == len(vieja) - 1) { 	// Caso de los extremos
		fmt.Printf("ANTESC-1: \t%2.15f\n",distTotal)
		// Borramos la primer arista
		if (funciones.Distancias[i][i+1] == 0 && funciones.Distancias[i+1][i] == 0){
			distTotal -= funciones.PesoAumentado(vieja[i],vieja[i+1])
			fmt.Printf("ACTNOPrimeraC-1: \t%2.15f\n",distTotal)
		} else {
			distTotal -= funciones.Distancias[i][i+1] - funciones.Distancias[i+1][i]
			fmt.Printf("ACTSIPrimeraC-1: \t%2.15f\n",distTotal)
		}// Borramos la última arista
		if (funciones.Distancias[j][j-1] == 0 && funciones.Distancias[j-1][j] == 0) {
			// no esta la arista borramos la arista con peso aumentado
			distTotal -= funciones.PesoAumentado(vieja[j], vieja[j-1])
			fmt.Printf("ACTNOSegundaC-1: \t%2.15f\n",distTotal)
		} else { // Si está la arista la borramos
			distTotal -= funciones.Distancias[j][j-1] - funciones.Distancias[j-1][j]
			fmt.Printf("ACTSISegundaC-1: \t%2.15f\n",distTotal)
		}// Agregamos la primer arista
		if (funciones.Distancias[j][i+1] == 0 && funciones.Distancias[i+1][j] == 0) {
			// No está la arista
			distTotal += funciones.PesoAumentado(vieja[i+1], vieja[j])
			fmt.Printf("SumaPrimeraPESOAUMENTADOC-1: \t%2.15f\n",distTotal)
		} else {
			// Si está la arista
			distTotal += funciones.Distancias[j][i+1] + funciones.Distancias[i+1][j]
			fmt.Printf("SumaPrimeraDISTANCIANORMALC-1: \t%2.15f\n",distTotal)
		}// Agregamos la última arista
		if (funciones.Distancias[i][j-1] == 0 && funciones.Distancias[j-1][i] == 0) {
			// No está la arista, calcular Peso aumentado
			distTotal += funciones.PesoAumentado(vieja[j-1], vieja[i])
			fmt.Printf("SumaSegundoPESOAUMENTADOC-1: \t%2.15f\n",distTotal)
		} else { // Esta la arista
			distTotal += funciones.Distancias[j-1][i] + funciones.Distancias[i][j-1]
			fmt.Printf("SumaSegundaDISTANCIANORMALC-1: \t%2.15f\n",distTotal)
		}
		fmt.Printf("DESPUESC-1: \t%2.15f\n",distTotal)
		return distTotal
	} else if (i == 0 && j != 0 && j != len(vieja)) { // Caso de un extremo
		fmt.Printf("ANTESC-2: \t%2.15f\n",distTotal) //cambiar 3 aristas
		if (funciones.Distancias[i][i+1] == 0 && funciones.Distancias[i+1][i] == 0) {
			// Primer arista  no está
			distTotal -= funciones.PesoAumentado(vieja[i], vieja[i+1])
			fmt.Printf("ACTNOPrimeraC-2: \t%2.15f\n",distTotal)
		} else { // Primer arista y si está
			distTotal -= funciones.Distancias[i][i+1] - funciones.Distancias[i+1][i]
			fmt.Printf("ACTSIPrimeraC-2: \t%2.15f\n",distTotal)
		} // Segunda arista borrar
		if (funciones.Distancias[j][j+1] == 0 && funciones.Distancias[j+1][j] == 0) { // Segunda arista borrar
			distTotal -= funciones.PesoAumentado(vieja[j], vieja[j+1])
			fmt.Printf("ACTNOSegundaC-2: \t%2.15f\n",distTotal)			
		} else {
			distTotal -= funciones.Distancias[j][j+1] - funciones.Distancias[j+1][j]
			fmt.Printf("ACTSISegundaC-2: \t%2.15f\n",distTotal)
		} // Tercer arista borrar
		if (funciones.Distancias[j][j-1] == 0 && funciones.Distancias[j-1][j] == 0) {
			distTotal -= funciones.PesoAumentado(vieja[j], vieja[j-1])
			fmt.Printf("ACTNOTERCERAC-2: \t%2.15f\n",distTotal)
		} else {
			distTotal -= funciones.Distancias[j][j-1] - funciones.Distancias[j-1][j]
			fmt.Printf("ACTSITERCERAC-2: \t%2.15f\n",distTotal)
		} // Agregamos primera arista
		if (funciones.Distancias[j][i+1] == 0 && funciones.Distancias[i+1][j] == 0) {
			distTotal += funciones.PesoAumentado(vieja[j],vieja[i+1])
			fmt.Printf("SumaPrimeraPESOAUMENTADOC-2: \t%2.15f\n",distTotal)
		} else {
			distTotal += funciones.Distancias[j][i+1] + funciones.Distancias[i+1][j]
			fmt.Printf("SumaPrimeraDISTANCIANORMALC-2: \t%2.15f\n",distTotal)
		} // Agregamos segunda arista
		if (funciones.Distancias[i][j+1] == 0 && funciones.Distancias[j+1][i] == 0) {
			distTotal += funciones.PesoAumentado(vieja[i], vieja[j+1])
			fmt.Printf("SumaSegundaPESOAUMENTADOC-2: \t%2.15f\n",distTotal)
		} else {
			distTotal += funciones.Distancias[i][j+1] + funciones.Distancias[j+1][i]
			fmt.Printf("SumaSegundaDISTANCIANORMALC-2: \t%2.15f\n",distTotal)
		} // Agregamos tercera arista
		if (funciones.Distancias[i][j-1] == 0 && funciones.Distancias[j-1][i] == 0) {
			distTotal += funciones.PesoAumentado(vieja[i], vieja[j-1])
			fmt.Printf("SumaTerceraPESOAUMENTADOC-2: \t%2.15f\n",distTotal)
		} else {
			distTotal += funciones.Distancias[i][j-1] + funciones.Distancias[j-1][i]
			fmt.Printf("SumaTerceraDISTANCIANORMALC-2: \t%2.15f\n",distTotal)
		}
		fmt.Printf("DESPUESC-2: \t%2.15f\n",distTotal)
		return distTotal










		









	} else if (j == 0) { // El otro caso
		
	} else { // Caso de dos internos
		
	}
	return distTotal
}

func main() {
	
	var ciudades40 = []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,
		332,333,489,490,491,492,493,496,652,653,654,656,657,792,815,816,
		817,820,978,979,980,981,982,984}

	funciones.Init(ciudades40)
	
	actual := Solucion{
		Ciudades: ciudades40,
		Funcion: funciones.FunCosto(ciudades40),
		Distancia: funciones.FunCosto(ciudades40) *
			funciones.Normalizador()}

	permutacion, i, j := Vecino(ciudades40)
	
	nuevaDist := Actualiza(actual.Ciudades, i, j, actual.Distancia)	
	nueva := Solucion{
		Ciudades: permutacion,
		Funcion: nuevaDist/funciones.Normalizador(),
		Distancia: nuevaDist}
	
	// nuevaDist := Actualiza(actual.Ciudades, nueva.Ciudades, 0, 39, actual.Distancia)

	// nuevaDist := Actualiza(actual.Ciudades, i, j, actual.Funcion)
	fmt.Println(len(actual.Ciudades))
	fmt.Println(nuevaDist)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(actual)
	fmt.Println(nueva)
	fmt.Println(len(nueva.Ciudades))
	

	// Vecino();
	// fmt.Println(Actual)
	// fmt.Println(Nueva)

}
