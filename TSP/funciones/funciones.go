package funciones

import (
	"sort"
)

type Ciudadeser interface {
	totalAristas([][]float64)
	FunCosto([][]float64)	
}

type ciudades struct {
	aristasE []float64
	id []int
	maxDist float64
	costo float64 // Sin normalizador
}

// Para cada par no ordenado, si la arista está en las distancias (tsp.sql), la
// agregamos a una lista
// Regresa todas las aristas en E
// func totalAristas(distancias [][]float64, ciudades []int) ([]float64, float64) {
func (c ciudades) totalAristas(distancias [][]float64) {
	var totalAristas []float64
	for i := 0; i < len(c.id); i++ {
		// ¿j = i?
		for j := 0; j < len(c.id); j++ {
			// Si está en las distancias agregamos
			if distancias[i][j] != 0 {
				totalAristas = append(totalAristas,
					distancias[i][j])
			}
		}
	}
	sort.Float64s(totalAristas) // Ordenado de menor a mayor
	c.maxDist = totalAristas[len(totalAristas)-1]
	c.aristasE = totalAristas
	// return totalAristas, totalAristas[len(totalAristas)-1]
	// return totalAristas
}

// Regresa la suma de las últimas k aristas
// Recibe todas las aristas en E y todas las ciudades
func Normalizador(kAristas, id []float64) float64 {
	suma := 0.0
	end := len(kAristas)-len(id)
	for i := len(kAristas)-1; i > end; i-- {
		suma += kAristas[i]
	}
	return suma
}

// Funcion que calcula el numerador de la funcion de costo, hay que divirlo
// entre el normalizador
// Recibe los id de las ciudades
// Recibe la matriz con las distancias
// Regresa la suma de las distancias de la arista (i, i-1)
// func FunCosto(ciudadesID []int, distancias [][]float64, max float64) float64 {
func (c ciudades) FunCosto(distancias [][]float64) {
	suma := 0.0
	for i := 1; i < len(c.id); i++ {
		if (distancias[i][i-1] == 0 && distancias[i-1][i] == 0) {
			suma += pesoAumentado(c.id[i], c.id[i-1], c.maxDist)
		} else {
			// No sabemos en qué parte de la matriz esta
			suma += distancias[i][i-1] + distancias[i-1][i]
		}
	}
	c.costo = suma
	// return suma
}

func NewCiudades() Ciudadeser {
	return &ciudades{}
}
