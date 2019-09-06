package funciones

import (
	// "fmt"
	
	"sort"
	"math"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const radio = 6373000
var ID []int
var Distancias [][]float64
var AristasE []float64
var MaximaDist float64


// Calcula la funcion de costo, recibe los ID's de las ciudades, las distancias y
// la máxima distancia.
// func FunCosto(ciudadesID []int) (float64, float64) {
func FunCosto(ciudadesID []int) float64 {
	suma := 0.0
	for i := 1; i < len(ciudadesID); i++ {
		if (Distancias[i][i-1]) == 0 && Distancias[i-1][i] == 0 {
			suma += PesoAumentado(ciudadesID[i], ciudadesID[i-1])
		} else {
			suma += Distancias[i][i-1] + Distancias[i-1][i]
		}
	}
	// return suma/Normalizador(), suma
	return suma/Normalizador()
}

// Regresa la suma de las últimas k aristas
func Normalizador() float64 {
	suma := 0.0
	end := len(AristasE)-len(ID)
	for i := len(AristasE)-1; i > end; i-- {
		suma += AristasE[i]
	}
	return suma
}

// Dados los ID's, regresa la matriz de la gráfica completa de los ID's
func completa(ciudades []int) [][]float64{
	var matriz = [][]float64{}
	var query = "SELECT distance FROM connections WHERE id_city_1 = ? AND id_city_2 = ?"
	database, _ := sql.Open("sqlite3", "../base/tsp.db")
	for i := 0; i < len(ciudades); i++ {
		adyacentes := make([]float64, len(ciudades))
		for j := 0; j < len(ciudades); j++ {
			var distance float64
			rows, _ := database.Query(query, ciudades[i],
				ciudades[j])
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
				totalAristasE = append(totalAristasE,
					distancias[i][j])
			}
		}
	}
	sort.Float64s(totalAristasE) // ordenado de menor a mayor
	return totalAristasE
}

// Regresa la distancia natural entre dos (u,v) ciudades dadas por su ID.
func distanciaNatural(u, v int) float64 {
	a := obtenerA(u, v)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return float64(radio) * c
}

// Regresa la coordenada dada en radianes.
func radianes(coordenada float64) float64 {
	return (coordenada*math.Pi)/180
}

// Obtiene A la fórmula del PDF.
func obtenerA (u, v int) float64{
	latV, lonV := obtenerLatLon(v)
	latU, lonU := obtenerLatLon(u)
	sin1 := math.Pow(math.Sin((latV-latU)/2), 2)
	sin2 := math.Pow(math.Sin((lonV-lonU)/2), 2)
	cos1 := math.Cos(latU)
	cos2 := math.Cos(latV)
	return sin1 + cos1 * cos2 * sin2
}

// Obtiene la longitud y longitud de la ciudad con el ID i.
func obtenerLatLon(i int) (latitud, longitud float64) {
	database, _ := sql.Open("sqlite3", "../base/tsp.db")
	var lat, lon float64
	rows, _ := database.Query(
		"SELECT latitude, longitude FROM cities WHERE id = ?", i)
	for rows.Next() {
		rows.Scan(&lat, &lon)
	}
	return radianes(lat), radianes(lon)
}

// Calcula el peso aumentado, o sea, la distancia natural por la máxima
// distancia.
func PesoAumentado(i, j int) float64 {
	dist := distanciaNatural(i, j)
	// fmt.Printf("PESO AUMENTADO(%d-%d): \t%2.15f\n",i ,j, dist * MaximaDist)
	return dist * MaximaDist
}

func Init(ciudadesID []int) {
	ID = ciudadesID
	Distancias = completa(ciudadesID)
	AristasE = totalAristas(Distancias, ID)
	MaximaDist = AristasE[len(AristasE)-1]
}
