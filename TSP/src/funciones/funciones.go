package funciones

import (
	"sort"
)

// Lista ordenada de mayor a menor (?) que contiene todas las distancias que
// existen en cada par ordenado de E

// Recibe la lista(arrray/slice) de los ID de las ciudades, vemos si existe la
// arista y sacamos la distancia que se agrega a l
func llenaListaL (ciudadesID []int) []float64{
	var e []float64
	for i := 0; i < len(ciudadesID); i++ {
		for j := i; j < len(ciudadesID); j++ {
			f := ciudadesID[i]
			s := ciudadesID[j]
			if (contenidasEnE(f,s)) {
				e = append(e, getDistancia(f,s))
			}
		}
	}
	sort.Float64s(e)
	return e
}

// Normalizador como número
// Recibimos la lista de las S aristas en E
// Regresamos la suma de las n-1 aristas más pesadas de E
func normalizador(listaE []float64) float64 {
	normalizador := 0.0
	for i := 0; i < len(listaE)-1; i++ {
		normalizador += listaE[i]
	}
	return normalizador
}

// Suma todos los pesos de la permutación, suma el i con el i-1
func funcionCostoSuma(ciudadesID []int) float64{
	suma := 0.0
	for i := 1; i < len(ciudadesID); i++ {
		suma += getDistancia(ciudadesID[i-1], ciudadesID[i])
	}
	return suma	
}

// Regresa la funcion de costo
// Recibe la lista(array/slice) cd ID's de ciudades
func FuncionCosto(ciudadesID []int) float64{
	listaMax := llenaListaL(ciudadesID)
	ns := normalizador(listaMax)
	suma := funcionCostoSuma(ciudadesID)
	return suma/ns
}

// Nos dice si dos aristas están en la gráfica normal
func contenidasEnE(i, j int) bool {
	// ver si están conectadas
	return true;
}

// Regresa la distancia entre dos aristas
func getDistancia(i, j int) float64 {
	// regresar la distancia entre id i-j
	return 0.0
}
