package main

import (
	"sort"
	"database/sql"
	"fmt"
	// "strconv"
	_ "github.com/mattn/go-sqlite3"
)

type MatricesDistancias struct {
	values [][]float64
}

// IMPORTANTE 
// DEBERIA HACER VARIABLES GLOBALES PARA NO TENER QUE ESTAR PASANDO SIEMPRE LAS DISTANCIAS

// Lista ordenada de mayor a menor (?) que contiene todas las distancias que
// existen en cada par ordenado de E

// Recibe la lista(arrray/slice) de los ID de las ciudades, vemos si existe la
// arista y sacamos la distancia que se agrega a l

func LlenaListaL (ciudadesID []int, distancias [][]float64) []float64{
// func LlenaListaL (distancias [][]float64) []float64{
	var e []float64
	for i := 0; i < len(ciudadesID); i++ {
	// for i := 0; i < len(distancias); i++ {
		for j := i+1; j < len(ciudadesID); j++ {
		
		// for j := 0; j < len(distancias[i]); j++ {
			// f := ciudadesID[i]
			// s := ciudadesID[j]
			// fmt.Printf("%df, %ds\n", f,s)
			
			if (contenidasEnE(i, j, distancias)) {
				// e = append(e, getDistancia(f,s, distancias))
				// fmt.Printf("([%d]/%d)-([%d]/%d)\n", ciudadesID[i],ciudadesID[i], ciudadesID[j],ciudadesID[j])
				e = append(e, getDistancia(i,j, distancias))
			}

			// fmt.Printf("(%d,%d), %E\n", i+1,j+1, distancias[i][j])

		}
	}
	fmt.Println(e)
	sort.Float64s(e) //COMO SORTEA (?)
	fmt.Println(e)
	suma := 0.0
	fmt.Println(len(e))
	// for i := 0; i < len(e)-1; i++ {
	for i := len(e)-1; i >= 0; i-- {
	// for i := len(e)-1; i > 1; i++ {	
		suma += e[i]
	}
	fmt.Printf("\n\n%2.15f\n\n", suma/2)
	return e
}

// Normalizador como número
// Recibimos la lista de las S aristas en E
// Regresamos la suma de las n-1 aristas más pesadas de E
func Normalizador(listaE []float64) float64 {
	normalizador := 0.0
	for i := 0; i < len(listaE)-1; i++ {
		normalizador += listaE[i]
	}
	return normalizador
}

// Suma todos los pesos (lineales) de la permutación, suma el i con el i-1
func FuncionCostoSuma(ciudadesID []int, distancias [][]float64) float64{
	suma := 0.0
	for i := 1; i < len(ciudadesID); i++ {
		// suma += getDistancia(ciudadesID[i-1], ciudadesID[i], distancias)
		suma += getDistancia(i-1, i, distancias)
	}

	fmt.Printf("suma: %2.15f", suma*2)
	return suma	
}

// Regresa la funcion de costo
// Recibe la lista(array/slice) cd ID's de ciudades
// func FuncionCosto(ciudadesID []int) float64{
func FuncionCosto(ciudadesID []int, distancias [][]float64) float64{
	
	listaMax := LlenaListaL(ciudadesID, distancias)

	// listaMax := LlenaListaL(distancias)

	
	ns := Normalizador(listaMax)
	suma := FuncionCostoSuma(ciudadesID, distancias)
	return suma/ns
}

// Nos dice si dos aristas están en la gráfica normal
// func contenidasEnE(i, j int) bool {
func contenidasEnE(i, j int, distancias [][]float64) bool {
	return distancias[i][j] != 0.0 || distancias[j][i] != 0.0
	// return distancias[i][j] != 0.0
	// return true;
}

// Regresa la distancia entre dos aristas
func getDistancia(i, j int, distancias [][]float64) float64 {
	// fmt.Printf("%di-%dj getDistancia \n",i,j)
	// regresar la distancia entre id i-j
	// if distancias[i][j] != 0 {
		// return distancias[i][j]
	// } else {
		// return distancias[j][i]
	// }
	return distancias[i][j]
	// return 0.0
}

func obtenerDistancias(ciudades []int) [][]float64{ //hay quepasar las ciudades, pero las tengo abajo
	// func main() {
	// var ciudades = []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,332,333,489,490,491,492,493,496,652,653,654,656,657,792,815,816,817,820,978,979,980,981,982,984}
	var matriz = [][]float64{}
	database, _ := sql.Open("sqlite3", "../base/tsp.db")
	for i := 0; i < len(ciudades); i++ {
		adyacentes := make([]float64, len(ciudades))
		for j := 0; j < len(ciudades); j++ {
			var distance float64
			rows, _ := database.Query("SELECT distance FROM connections WHERE id_city_1 = ? AND id_city_2 = ?", ciudades[i], ciudades[j])
			for rows.Next() {
				rows.Scan(&distance)
				adyacentes[j] = distance
			}
			
		}
		matriz = append(matriz, adyacentes)
		// fmt.Println(adyacentes)
	}
	// fmt.Println(matriz)
	return matriz
}

func main() {
	var ciudades = []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,332,333,489,490,491,492,493,496,652,653,654,656,657,792,815,816,817,820,978,979,980,981,982,984}
	var distancias MatricesDistancias
	distancias.values = obtenerDistancias(ciudades)
	norm := FuncionCosto(ciudades, distancias.values)
	fmt.Printf("%E NORM", norm)
	
	// fmt.Printf("%d longitud total\n",len(distancias.values))
	// for i := 0; i < len(distancias.values[0]); i++ {
		// fmt.Printf("%d, sublongitud\n", len(distancias.values[i]))
		// fmt.Printf("%E --- %d\n", distancias.values[0][i], i)
	// }
}
