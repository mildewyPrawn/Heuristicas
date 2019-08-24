// package sql
package main

import (
	"database/sql"
	"fmt"
	"sort"
	// "strconv"
	_ "github.com/mattn/go-sqlite3"
	"math"
)

var radio = 6373000

type MatricesDistancias struct {
	values [][]float64
}

// Dados los ID's, regresa la matriz de la gráfica completa de los ID's
func ciudades(ciudades []int) [][]float64{ //hay quepasar las ciudades, pero las tengo abajo
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
	}
	return matriz
}

// Para cada par no ordenado si la arista está en distancias(en tsp.sql) la
// agregamos a una lista, la cual está ordenada de menor a mayor.
func totalAristas(distancias [][]float64, ciudades []int) []float64{
	var totalAristasE []float64
	for i := 0; i < len(ciudades); i++ {
		for j := 0; j < len(ciudades); j++ {
			if distancias[i][j] != 0 {
				totalAristasE = append(totalAristasE, distancias[i][j])
			}
		}
	}
	sort.Float64s(totalAristasE) // sorted ascending
	return totalAristasE
}

// Regresa la suma de las últimas k aristas
func normalizador(aristas []float64, k int) float64 {
	suma := 0.0
	end := len(aristas)-k
	for i := len(aristas)-1; i > end; i-- {
		suma += aristas[i]
	}
	return suma
}

// Regresa la distancia natural entre dos (u,v) ciudades dadas por su ID.
// func distanciaNatural(u, v int, max float64) float64 {
func distanciaNatural(u, v int, ciudadesID []int) float64 {
	a := obtenerA(u, v, ciudadesID)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return float64(radio) * c
}

// Regresa la coordenada dada en radianes.
func radianes(coordenada float64) float64 {
	return (coordenada*math.Pi)/180
}

// Obtiene A la fórmula del PDF.
func obtenerA (u, v int, ciudadesID []int) float64{
	latV, lonV := obtenerLatLon(v, ciudadesID)
	latU, lonU := obtenerLatLon(u, ciudadesID)
	
	sin1 := math.Pow(math.Sin((latV-latU)/2), 2)
	sin2 := math.Pow(math.Sin((lonV-lonU)/2), 2)
	cos1 := math.Cos(latU)
	cos2 := math.Cos(latV)

	return sin1 + cos1 * cos2 * sin2

}

// Obtiene la longitud y longitud de la ciudad con el ID i.
func obtenerLatLon(i int, ciudadesID []int) (latitud, longitud float64) {
	database, _ := sql.Open("sqlite3", "../base/tsp.db")
	var lat, lon float64
	rows, _ := database.Query("SELECT latitude, longitude FROM cities WHERE id = ?", i)
	for rows.Next() {
		rows.Scan(&lat, &lon)
	}
	return radianes(lat), radianes(lon)
}


func main() {
	// var ciudadesID = []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,332,333,489,490,491,492,493,496,652,653,654,656,657,792,815,816,817,820,978,979,980,981,982,984}
	var ciudadesID = []int{1,2,3,4,5,6,7,8,9,11,12,14,16,17,19,20,22,23,25,26,27,74,75,77,163,164,165,166,167,168,169,171,172,173,174,176,179,181,182,183,184,185,186,187,297,326,327,328,329,330,331,332,333,334,336,339,340,343,344,345,346,347,349,350,351,352,353,444,483,489,490,491,492,493,494,495,496,499,500,501,502,504,505,507,508,509,510,511,512,520,652,653,654,655,656,657,658,660,661,662,663,665,666,667,668,670,671,673,674,675,676,678,792,815,816,817,818,819,820,821,822,823,825,826,828,829,832,837,839,840,978,979,980,981,982,984,985,986,988,990,991,995,999,1001,1003,1004,1037,1038,1073,1075}
	
	var distancias MatricesDistancias
	var normaliz, maximaDist float64
	distancias.values = ciudades(ciudadesID)
	// distancias.values = ciudades()
	
	aristasE := totalAristas(distancias.values, ciudadesID)
	// fmt.Println(distancias.values)

	fmt.Printf("TOTAL ARISTAS:\t %d\n", len(aristasE))
	maximaDist = aristasE[len(aristasE)-1]
	fmt.Printf("MAXIMA DISTANCIA\t %2.15f\n", maximaDist)
	
	if len(aristasE) < len(ciudadesID) {
		normaliz = normalizador(aristasE, len(aristasE))
	}
	normaliz = normalizador(aristasE, len(ciudadesID))
	fmt.Printf("NORMALIZADOR:\t %2.15f\n", normaliz)
	// fmt.Println(aristasE)

	// fmt.Printf("distanciaNatural:\t%2.15f",distanciaNatural(0,6, ciudadesID))
	fmt.Printf("DISTANCIA NATURAL(%d-%d):\t%2.15f",1 ,7, distanciaNatural(1,7, ciudadesID))

}
