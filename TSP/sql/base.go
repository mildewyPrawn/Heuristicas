// package sql
package main

import (
	"database/sql"
	"fmt"
	// "strconv"
	_ "github.com/mattn/go-sqlite3"
)

type MatricesDistancias struct {
	values [][]float64
}

// func ciudades(ciudades  []int) {
func ciudades() [][]float64{ //hay quepasar las ciudades, pero las tengo abajo
	// func main() {
	var ciudades = []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,332,333,489,490,491,492,493,496,652,653,654,656,657,792,815,816,817,820,978,979,980,981,982,984}
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
	var distancias MatricesDistancias
	distancias.values = ciudades()
	fmt.Println(distancias.values)
}
