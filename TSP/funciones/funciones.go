package funciones

import (
	// "math/rand"
	"math"
	"sort"
	"fmt"
)

type Ciudadeser interface {
	PrintCiudad()
	GetNormalizador() float64
	GetDistancias() [][]float64
	GetAristasE() []float64
	SetId([]int)
	GetId() []int
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

type Solucion struct {
	// temperatura float64 // Temperatura usada
	normalizador float64 // Normalizador usado 
	mejor []int
	mejorC float64
	actual []int
	actualC float64
	nueva []int
	nuevaC float64
}

// Solo imprime una ciudad con sus datos
func (c Ciudades) PrintCiudad() {
	fmt.Println(c.Id)
	fmt.Printf("MAX DIST:\t%2.15f\nDISTOT:\t%2.15f\n", 
		c.AristasE[len(c.AristasE)-1], c.Costo)
	fmt.Printf("NORMALIZADOR:\t%2.15f\n",c.Normalizador)
	fmt.Printf("FUNCOSTO:\t%2.15f\n",c.Costo/c.Normalizador)
}

// func (s *Solucion) PrintData(c *Ciudades) {
func (s *Solucion) PrintData() {
	fmt.Print("NORMALIZADOR:")
	fmt.Printf("\t%2.15f\n", s.normalizador)
	fmt.Println("MEJOR:")
	fmt.Println(s.mejor)
	fmt.Printf("\tcosto: %2.15f\n", s.mejorC/s.normalizador)
	fmt.Print()
	fmt.Println("ACTUAL:")
	fmt.Println(s.actual)
	fmt.Printf("\tcosto: %2.15f\n", s.actualC/s.normalizador)
	fmt.Println("NUEVA:")
	fmt.Println(s.nueva)
	fmt.Printf("\tcosto: %2.15f\n", s.nuevaC/s.normalizador)
	// fmt.Printf("\tcosto: %2.15f", FunCostoSolucion(s.nueva, c.Distancias, c.AristasE)/s.normalizador)
	
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
	end := 0
	if  len(aristasE) < len(ciudadesId)-1{
		end = len(aristasE)
	} else {
		end = len(aristasE)-len(ciudadesId)
	}
	for i := len(aristasE)-1; i > end; i-- {
		suma += aristasE[i]
	}
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
func calculaLote(t float64, ciu *Ciudades, sol *Solucion) (float64, []int) {
	dist := ciu.Distancias
	norm := sol.normalizador
	c := 0
	i := 0
	r := 0.0
	// s := ciu.Id
	s := sol.actual
	for c < L {
		fmt.Println("en el c<L")
		// fmt.Printf("C(%d)-I(%d)\n", c, i)
		sP := vecino(s)
		fsP := FunCostoSolucion(s, dist, ciu.AristasE)/norm
		fmt.Printf("%2.15f\n", fsP/norm)
		fmt.Printf("%2.15f\n", sol.actualC/norm)
		sol.PrintData()
		fmt.Println(i)
		if fsP <= sol.actualC/norm + t {
			fmt.Println("en el fsP < actual")
			sol.actual = sP
			sol.actualC = fsP
			if fsP < sol.mejorC {
				sol.mejor = sP
				sol.mejorC = fsP
			}
			s = sP
			c++
			r = r + fsP
		}
		i++;
		if i > L*L {
			return r/L, s
		}
	}
	return r/L, s
}

// func AceptacionPorUmbrales(t float64, sol*Ciudades) []int{
func (c *Ciudades) AceptacionPorUmbrales(t float64, sol *Solucion) []int{
	fmt.Println("ACEPTACION POR UMBRALES")
	s := c.Id
	p := 0.0
	for t > EPSILON {
		fmt.Println(t)
		fmt.Println(EPSILONP)
		q := math.MaxFloat64
		fmt.Println("Antes del p<q")
		for p < q {
			fmt.Println("en el p<q")
			q = p
			p, s = calculaLote(t, c, sol)
			
		}
		t = PHI*t
	}
	return s
}

func porcentajeAceptados(ciu *Ciudades, t float64, sol *Solucion) float64 {
	c := 0
	norm := sol.normalizador
	for i := 0; i < len(ciu.Id); i++ {
		sp := vecino(sol.actual)
		fsP := FunCostoSolucion(sp, ciu.Distancias, ciu.AristasE)
		sol.nueva = sp
		sol.nuevaC = fsP
		if fsP/norm <= sol.actualC/norm + t {
			c++
			sol.actual = sol.nueva
			sol.actualC = sol.nuevaC
			if sol.nuevaC < sol.mejorC {
				sol.mejor = sol.nueva
				sol.mejorC = sol.nuevaC
			}
		}
		sol.PrintData()
	}
	return float64(c)/float64(len(ciu.Id))
}

func busquedaBinaria(s *Ciudades, t1, t2 float64, sol *Solucion) float64{
	tm := (t1+t2)/2
	if t2 - t1 < EPSILONP {
		return tm
	}
	p := porcentajeAceptados(s, tm, sol)
	if math.Abs(P - p) < EPSILONP {
		return tm
	}
	if p > P {
		return busquedaBinaria(s, t1, tm, sol)
	} else {
		return busquedaBinaria(s, tm, t2, sol)
	}
}

// func temperaturaInicial(s*Ciudades, t float64) float64 {
func (s *Ciudades) TemperaturaInicial(t float64, sol *Solucion) float64 {
	fmt.Println("TEMPERATURA INICIAL")
	p := porcentajeAceptados(s, t, sol)
	var t1, t2 float64
	if math.Abs(P - p) <= EPSILONP {
		return t
	}
	if p < P {
		for p < P {
			t = 2*t
			p  = porcentajeAceptados(s, t, sol)
		}
		t1 = t/2
		t2 = t
	} else {
		for p > P {
			t = t/2
			p = porcentajeAceptados(s, t, sol)
		}
		t1 = t
		t2 = 2*t
	}
	return busquedaBinaria(s, t1, t2, sol)
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

func (c *Ciudades) GetId() []int {
	return c.Id
}

func (c *Ciudades) GetAristasE() []float64 {
	return c.AristasE
}

func (c *Ciudades) GetNormalizador() float64 {
	return c.Normalizador
}

func NewSolucion(aristas []float64, id []int, dist [][]float64) *Solucion {
	norm := GetNormalizador(aristas, id)
	costo := FunCostoSolucion(id, dist, aristas)
	return &Solucion{
		// temperatura: 0.0,
		normalizador: norm,
		mejor: id,
		mejorC: costo,
		actual: id,
		actualC: costo,
		nueva: id,
		nuevaC: costo,
	}
}

func NewCiudades(ciudadesId []int) *Ciudades {
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

/*
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
}*/
