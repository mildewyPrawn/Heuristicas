package funciones

import (
	"testing"
	"fmt"
)

// Test para obtenerLatLon
func TestObtenerLatLon(t *testing.T) {
	lat1 := radianes(35.685000000000002273)
	lon1 := radianes(139.75100000000000477)
	lat7 := radianes(14.604200000000000514)
	lon7 := radianes(120.98199999999999931)
	var id = []int{1,7}
	for i := range id {
		latI, lonI := obtenerLatLon(id[i])
		if i == 0 {
			if lat1 != latI && lon1 != lonI {
				fmt.Println("Error en latitudes/longitudes")
			}
		}
		if i == 1 {
			if lat7 != latI && lon7 != lonI {	
				fmt.Println("Error en latitudes/longitudes")			
			}
		}
	}

}

// Test para la gráfica completa de 1-10
func TestCompleta(t *testing.T) {
	var unoDiez = []int{1,2,3,4,5,6,7,8,9,10}
	ciudades1a10 := completa(unoDiez)
	var miUnoDiez = [10][10]float64{
		{0,0,0,0,0,0,2999396.2299999999813,0,1158707.3100000000559,0},
		{0,0,0,0,0,0,1829270.9099999999161,0,890547.02000000001858,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,1089251.9799999999814,1085175.2399999999907,0,0,0,0},
		{0,0,0,0,0,7596.0600000000004002,4756785.7000000001864,
			4339406.4999999999998,4684841.339999999851,0},
		{0,0,0,0,0,0,4757152.5899999998511,0,4689521.1500000003727,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
	}
	if ciudades1a10 == nil {
		fmt.Println("Es nil")
	}
	if len(ciudades1a10) != len(miUnoDiez) {
		fmt.Println("No es la misma longitud")
	}
	for i := 0; i < len(unoDiez); i++ {
		for j := 0; j < len(unoDiez); j++ {
			if ciudades1a10[i][j] != miUnoDiez[i][j] {
				fmt.Println("No son iguales.")
			}
		}
	}
}
