package funciones

import (
	"sort"
	"fmt"
)

type Ciudadeser interface {
	PrintCiudad()
	GetDistTotal() float64
	FunCosto()
}

type ciudades struct {
	Id []int
	Distancias [][]float64
	AristasE []float64
	Costo float64 // Sin normalizador
	Normalizador float64
}

func (c ciudades) PrintCiudad() {
	fmt.Println(c.Id)
	fmt.Printf("MAX DIST:\t%2.15f\nDISTOT:\t%2.15f\n", 
		c.AristasE[len(c.AristasE)-1], c.Costo)
	fmt.Printf("NORMALIZADOR:\t%2.15f\n",c.Normalizador)
	fmt.Printf("FUNCOSTO:\t%2.15f\n",c.Costo/c.Normalizador)
}

// Para cada par no ordenado, si la arista está en las distancias (tsp.sql), la
// agregamos a una lista
// Regresa todas las aristas en E
func totalAristas(ciudadesId []int, distancias [][]float64) []float64 {
	var totalAristas []float64
	for i := 0; i < len(ciudadesId); i++ {
		// ¿j = i?
		for j := 0; j < len(ciudadesId); j++ {
			// Si está en las distancias agregamos
			if distancias[i][j] != 0 {
				totalAristas = append(totalAristas,
					distancias[i][j])
			}
		}
	}
	sort.Float64s(totalAristas) // Ordenadonn de menor a mayor
	return totalAristas
}

// Regresa la suma de las últimas k aristas
// Recibe todas las aristas en E y todas las ciudades
func GetNormalizador(aristasE []float64, ciudadesId []int) float64 {
	suma := 0.0
	end := len(aristasE)-len(ciudadesId)
	for i := len(aristasE)-1; i > end; i-- {
		suma += aristasE[i]
	}
	// c.Normalizador = suma
	return suma
}

// Funcion que calcula el numerador de la funcion de costo, hay que divirlo
// entre el normalizador
// Recibe los id de las ciudades
// Recibe la matriz con las distancias
// Regresa la suma de las distancias de la arista (i, i-1)
func (c *ciudades) FunCosto() {
	suma := 0.0
	for i := 1; i < len(c.Id); i++ {
		if (c.Distancias[i][i-1] == 0 && c.Distancias[i-1][i] == 0) {
			suma += pesoAumentado(c.Id[i], c.Id[i-1],
				c.AristasE[len(c.AristasE)-1])
		} else {
			// No sabemos en qué parte de la matriz esta
			suma += c.Distancias[i][i-1] + c.Distancias[i-1][i]
		}
	}
	c.Costo = suma
}

// Creo que es un getter :o
func (c *ciudades) GetDistTotal() float64 {
	return c.Costo
}

func NewCiudades(ciudadesId []int) Ciudadeser {
	dista := completa(ciudadesId)
	arist := totalAristas(ciudadesId, dista)
	norma := GetNormalizador(arist, ciudadesId)
	return &ciudades{
		Id: ciudadesId,
		Distancias: dista,
		AristasE: arist,
		Normalizador: norma,
	}
}
