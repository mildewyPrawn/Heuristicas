package sql

import (
	_ "github.com/mattn/go-sqlite3"
)

const QUERY_SELECT_IDS = "SELECT distance FROM connections WHERE id_city_1 = ? AND id_city_2 = ?"
const QUERY_LAT_LON = "SELECT latitude, longitude FROM cities WHERE id = ?"

