package funciones

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"

)

const QUERY_SELECT_IDS = "SELECT distance FROM connections WHERE id_city_1 = ? AND id_city_2 = ?"
const QUERY_LAT_LON = "SELECT latitude, longitude FROM cities WHERE id = ?"

func ObtenerLatLon(i int) (latitud, longitud float64) {
	database, _ := sql.Open("sqlite3", "./tsp.db")
	var lat, lon float64
	rows, _ := database.Query(QUERY_LAT_LON, i)
	for rows.Next() {
		rows.Scan(&lat, &lon)
	}
	return radianes(lat), radianes(lon)
}
