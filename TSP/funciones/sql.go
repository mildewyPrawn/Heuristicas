package funciones

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"sort"
)

const (
	// Query para seleccionar conexiones
	QUERY_DIS = "SELECT distance FROM connections WHERE id_city_1 = ? AND id_city_2 = ?"
	// Query para seleccionar todos los datos de una ciudad en la bd
	QUERY_DATA = "SELECT id, name, country, population, latitude, longitude FROM cities WHERE id = ?"
)

// Conexiona a la base de datos
var database, _ = sql.Open("sqlite3", "../base/tsp.db")

// Funcion que dados los ids de las ciuades, regresa un arreglo con ciudades, i.e,
// con los datos de cada ciudad
func ciudades(ids []int) []Ciudad {
	ciudades := []Ciudad{}
	for i := 0; i < len(ids); i++ {
		var id int
		var name string
		var country string
		var population int
		var latitude float64
		var longitude float64
		rows, _ := database.Query(QUERY_DATA, ids[i])
		for rows.Next() {
			rows.Scan(&id, &name, &country, &population, &latitude, &longitude)
			ci := Ciudad{i, id, name, country, population, latitude, longitude}
			ciudades = append(ciudades, ci)
		}
	}
	return ciudades
}

// Funcion que recibe un arreglo de ciudades y regresa las de forma ordenada las
// aristas consecutivas existen en el arreglo 
func TotalAristas(cis []Ciudad) []float64 {
	var aristas []float64
	for i := 0; i < len(cis); i++ {
		for j := 0; j < len(cis); j++ {
			var distance float64
			rows, _ := database.Query(QUERY_DIS, cis[i].id, cis[j].id)
			for rows.Next() {
				rows.Scan(&distance)
				aristas = append(aristas, distance)
			}
		}
	}
	sort.Float64s(aristas)
	return aristas
}

// Función que calcula el peso aumentado entre dos ciudades
func pesoAumentado(u, v Ciudad, max float64) float64 {
	if (v.id > u.id) {
		aux := v
		v = u
		u = aux
	}
	rows, _ := database.Query(QUERY_DIS, v.id, u.id)
	var distance float64
	i := 0
	for rows.Next() {
		rows.Scan(&distance)
		i++
	}
	if i > 0 {
		return distance
	} else {
		distNat := distanciaNatural(u, v)
		return distNat * max
	}
}
