package funciones

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

// Query para seleccionar conexiones
const QUERY_IDS = "SELECT distance FROM connections WHERE id_city_1 = ? AND id_city_2 = ?"
// Query para seleccionar latitud y longitud
const QUERY_LAT_LON = "SELECT latitude, longitude FROM cities WHERE id = ?"

// Conexiona a la base de datos
var database, _ = sql.Open("sqlite3", "../base/tsp.db") // No sé si funciona chido

// Regresa la latitud y longitud de una ciudad dada por su ID
func obtenerLatLon(i int) (latitud, longitud float64) {
	var lat, lon float64
	rows, _ := database.Query(QUERY_LAT_LON, i)
	for rows.Next() {
		rows.Scan(&lat, &lon)
	}
	return radianes(lat), radianes(lon)
}

// Regresa la representación de la gráfica dados los ID
// Solo son las ciudades que están conectadas
func completa(ciudades []int) [][]float64 {
	var matriz = [][]float64{}
	for i := 0; i < len(ciudades); i++ {
		adyacentes := make([]float64, len(ciudades))
		for j := 0; j < len(ciudades); j++ {
			var distance float64
			rows, _ := database.Query(QUERY_IDS, ciudades[i],
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
