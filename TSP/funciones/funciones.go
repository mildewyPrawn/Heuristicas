package funciones

import (
	"sort"
	"fmt"
)

type Ciudadeser interface {
	TotalAristas()
	FunCosto()
	PrintCiudad()
	Normalizador()
	// Init()
}

type ciudades struct {
	id []int
	distancias [][]float64
	aristasE []float64
	costo float64 // Sin normalizador
	normalizador float64
}

func (c ciudades) PrintCiudad() {
	fmt.Println(c.id)
	fmt.Printf("MAX DIST:\t%2.15f\nDISTOT:\t%2.15f\n", 
		c.aristasE[len(c.aristasE)-1], c.costo)
	fmt.Printf("NORMALIZADOR:\t%2.15f\n",c.normalizador)
	fmt.Printf("FUNCOSTO:\t%2.15f\n",c.costo/c.normalizador)
}

// Para cada par no ordenado, si la arista está en las distancias (tsp.sql), la
// agregamos a una lista
// Regresa todas las aristas en E
func (c *ciudades) TotalAristas() {
	var totalAristas []float64
	for i := 0; i < len(c.id); i++ {
		// ¿j = i?
		for j := 0; j < len(c.id); j++ {
			// Si está en las distancias agregamos
			if c.distancias[i][j] != 0 {
				totalAristas = append(totalAristas,
					c.distancias[i][j])
			}
		}
	}
	sort.Float64s(totalAristas) // Ordenadonn de menor a mayor
	c.aristasE = totalAristas
}

// Regresa la suma de las últimas k aristas
// Recibe todas las aristas en E y todas las ciudades
func (c *ciudades) Normalizador() {
	suma := 0.0
	end := len(c.aristasE)-len(c.id)
	for i := len(c.aristasE)-1; i > end; i-- {
		suma += c.aristasE[i]
	}
	c.normalizador = suma
	// return suma
}

// Funcion que calcula el numerador de la funcion de costo, hay que divirlo
// entre el normalizador
// Recibe los id de las ciudades
// Recibe la matriz con las distancias
// Regresa la suma de las distancias de la arista (i, i-1)
func (c *ciudades) FunCosto() {
	suma := 0.0
	for i := 1; i < len(c.id); i++ {
		if (c.distancias[i][i-1] == 0 && c.distancias[i-1][i] == 0) {
			suma += pesoAumentado(c.id[i], c.id[i-1],
				c.aristasE[len(c.aristasE)-1])
		} else {
			// No sabemos en qué parte de la matriz esta
			suma += c.distancias[i][i-1] + c.distancias[i-1][i]
		}
	}
	c.costo = suma
}

func NewCiudades(ciudadesId []int) Ciudadeser {
	// return &ciudades{id: ciudadesId}
	return &ciudades{
		id: ciudadesId,
		distancias: completa(ciudadesId),
	}
}
