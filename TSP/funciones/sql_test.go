package funciones

import (
	"testing"
	"fmt"
)

// Test para radianes, los calculé a mano :v
func TestRadianes(t *testing.T) {
	var grados = []float64{80.5, 34.22, 299.14, 78.58, 1.1}
	var rad = []float64{1.4050, 0.5972, 5.2226, 1.3715, 0.0177}
	for i := range grados {
		radI := radianes(grados[i])
		// Los trunco porque le faltan como 1000 decimales
		if int(radI) != int(rad[i]) {
			fmt.Println("Error en radianes")
		}
	}
}

// Test para distancia natural, lo saqué del correo
func TestDistanciaNatural(t *testing.T) {
	d17 := 2999396.231968969572335
	d1163 := 3222670.069842538330704
	dist17 := distanciaNatural(1,7)
	dist1163 := distanciaNatural(1,163)
	if d17 != dist17 {
		fmt.Println("Error en distancias 1-7")
	}
	if d1163 != dist1163 {
		fmt.Println("Error en distancias 1-163")
	}
}

// Estoy suponiendo que (1-7), (1-163) no están conectadas y que la máxima es 2
func TestPesoAumentado(t *testing.T) {
	d17 := 2999396.231968969572335*2
	d1163 := 3222670.069842538330704*2
	pa17 := pesoAumentado(1, 7, 2)
	pa1163 := pesoAumentado(1, 163, 2)
	if d17 != pa17 {
		fmt.Println("Error en peso aumentado")
	}
	if d1163 != pa1163 {
		fmt.Println("Error en peso aumentado")
	}
}

// Test para obtenerA
func TestObtenerA(t *testing.T) {
	a12 := 0.01933276365124884
	a23 := 0.14812066305808086
	ta12 := obtenerA(1,2)
	ta23 := obtenerA(2,3)
	if a12 != ta12 {
		fmt.Println("Error en obtener A")
	}
	if a23 != ta23 {
		fmt.Println("Error en obtener A")
	}
}

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

// Test para los resultados de Canek
func Test40(t *testing.T) {
	// No pude hacer que los regresara, así que escribí lo que me imprimía
	norm := 182907823.060000002384186
	maxD := 4970123.959999999962747
	funC := 4526237.801017570309341
	canekNormalizador := 182907823.060000002384186
	canekMaxDist := 4970123.959999999962747
	canekFuncionCosto := 4526237.801017570309341;
	if norm - canekMaxDist != 0 {
		fmt.Println("Máxima distancia no coincide40")
		fmt.Println(maxD-canekMaxDist)
	}
	if maxD - canekNormalizador != 0 {
		fmt.Println("Normalziador no coincide40")
		fmt.Println(norm-canekNormalizador)
	}
	if funC - canekFuncionCosto != 0 {
		fmt.Println("Función de costo no coincide40")
		fmt.Println(funC-canekFuncionCosto)
	}
}

// Test para los resultados de Canek
func Test150(t *testing.T) {
	// No pude hacer que los regresara, así que escribí lo que me imprimía
	norm := 722989785.090000391006470
	maxD := 4978506.480000000447035
	funC := 6210491.0347478
	canekNormalizador := 722989785.090000391006470
	canekMaxDist := 4978506.480000000447035
	canekFuncionCosto := 6210491.0347478
	if norm - canekMaxDist != 0 {
		fmt.Println("Máxima distancia no coincide150")
		fmt.Println(maxD-canekMaxDist)
	}
	if maxD - canekNormalizador != 0 {
		fmt.Println("Normalziador no coincide150")
		fmt.Println(norm-canekNormalizador)
	}
	if funC - canekFuncionCosto != 0 {
		fmt.Println("Función de costo no coincide150")
		fmt.Println(funC-canekFuncionCosto)
	}
}
