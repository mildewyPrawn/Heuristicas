package funciones

import (
	"math"
	"sort"
	"fmt"
)

type Ciudadeser interface {
	PrintCiudad()
	GetNormalizador() float64
	// GetDistTotal() float64
	GetDistancias() [][]float64
	GetAristasE() []float64
	SetId([]int)
	FunCosto()
	AceptacionPorUmbrales(float64) []int
	TemperaturaInicial(float64) float64
}

type Ciudades struct {
	Id []int
	Distancias [][]float64
	AristasE []float64
	Costo float64 // Sin normalizador
	Normalizador float64
}

func (c Ciudades) PrintCiudad() {
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
	// end := len(aristasE)-len(ciudadesId)
	end := 0
	if  len(aristasE) < len(ciudadesId)-1{
		end = len(aristasE)
	} else {
		end = len(aristasE)-len(ciudadesId)
	}
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
func (c *Ciudades) FunCosto() {
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

func FunCostoSolucion(id []int, distancias [][]float64,
	aristas []float64) float64 {
	suma := 0.0
	for i := 1; i < len(id); i++ {
		if (distancias[i][i-1] == 0 && distancias[i-1][i] == 0) {
			suma += pesoAumentado(id[i], id[i-1],
				aristas[len(aristas)-1])
		} else {
			// No sabemos en qué parte de la matriz esta
			suma += distancias[i][i-1] + distancias[i-1][i]
		}
	}
	return suma
}

// func calculaLote(t float64, ciudades *) (float64, []int) {
func calculaLote(t float64, ciu*Ciudades) (float64, []int) {
	fmt.Println("CALCULA LOTE")
	distancias := ciu.Distancias
	normaliz := ciu.Normalizador
	c := 0
	i := 0
	r := 0.0
	s := ciu.Id
	for c < L && i < L*L{
	// for c < L {
		fmt.Printf("C(%d)-I(%d)\n", c, i)
		sP := vecino(s)
		fsP := FunCostoSolucion(s, distancias, ciu.AristasE)/normaliz
		if fsP <= ciu.Costo/normaliz {
			ciu.SetId(sP)
			c++
			r = r + fsP
		}
		i++;
	}
	return r/L, s
}

// func AceptacionPorUmbrales(t float64, sol*Ciudades) []int{
func (sol *Ciudades) AceptacionPorUmbrales(t float64) []int{
	fmt.Println("ACEPTACION POR UMBRALES")
	s := sol.Id
	p := 0.0
	for t > EPSILON {
		q := math.MaxFloat64
		for p <= q {
			q = p
			p, s = calculaLote(t, sol)
			// printSol(p, s)
		}
		t = PHI*t
	}
	return s
}

func porcentajeAceptados(sol*Ciudades, t float64) float64 {
	fmt.Println("PORCENTAJE ACEPTADOS")
	c := 0
	for i := 0; i < len(sol.Id); i++ {
		sP := vecino(sol.Id)
		fsP := FunCostoSolucion(sol.Id, sol.Distancias, sol.AristasE)
		norm := sol.Normalizador
		if fsP/norm < sol.Costo/norm + t {
			c++
			sol.SetId(sP)
		}
	}
	return float64(c)/float64(len(sol.Id))
}

func busquedaBinaria(s*Ciudades, t1, t2 float64) float64{
	fmt.Println("BUSQUEDA BINARIA")
	tm := (t1+t2)/2
	if t2 - t1 < EPSILONP {
		return tm
	}
	p := porcentajeAceptados(s, tm)
	if math.Abs(P - p) < EPSILONP {
		return tm
	}
	if p > P {
		return busquedaBinaria(s, t1, tm)
	} else {
		return busquedaBinaria(s, tm, t2)
	}
}

// func temperaturaInicial(s*Ciudades, t float64) float64 {
func (s *Ciudades) TemperaturaInicial(t float64) float64 {
	fmt.Println("TEMPERATURA INICIAL")
	p := porcentajeAceptados(s, t)
	var t1, t2 float64
	if math.Abs(P - p) <= EPSILONP {
		return t
	}
	if p < P {
		for p < P {
			t = 2*t
			p  = porcentajeAceptados(s, t)
		}
		t1 = t/2
		t2 = t
	} else {
		for p > P {
			t = t/2
			p = porcentajeAceptados(s, t)
		}
		t1 = t
		t2 = 2*t
	}
	return busquedaBinaria(s, t1, t2)
}

func printSol(p float64, s []int) {
	fmt.Printf("Costo: %2.15f\t con ciudades:\n", p)
	fmt.Println(s)
}

func (c *Ciudades) SetId(id []int) {
	c.Id = id
}

func (c *Ciudades) GetDistancias() [][]float64 {
	return c.Distancias
}

func (c *Ciudades) GetAristasE() []float64 {
	return c.AristasE
}

func (c *Ciudades) GetNormalizador() float64 {
	return c.Normalizador
}

func NewCiudades(ciudadesId []int) Ciudadeser {
	dista := completa(ciudadesId)
	arist := totalAristas(ciudadesId, dista)
	norma := GetNormalizador(arist, ciudadesId)
	return &Ciudades{
		Id: ciudadesId,
		Distancias: dista,
		AristasE: arist,
		Normalizador: norma,
	}
}
