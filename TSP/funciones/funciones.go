package funciones

import (
	// "math/rand"
	"math"
	"sort"
	"fmt"
)

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
	// nueva []int
	// nuevaC float64
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
	fmt.Printf("\t%2.15f\n\n", s.normalizador)
	fmt.Println("MEJOR:")
	fmt.Println(s.mejor)
	fmt.Printf("\tcosto: %2.15f\n\n", s.mejorC)
	fmt.Print()
	fmt.Println("ACTUAL:")
	fmt.Println(s.actual)
	fmt.Printf("\tcosto: %2.15f\n\n", s.actualC)
	fmt.Println("NUEVA:")
	// fmt.Println(s.nueva)
	// fmt.Printf("\tcosto: %2.15f\n\n", s.nuevaC/s.normalizador)
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
	aristas []float64, norm float64) float64 {
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
	return suma/norm
}

func porcentajeAceptados(ciu *Ciudades, t float64, sol *Solucion) float64 {
	c := 0
	// fs := FunCostoSolucion(s, ciu.Distancias, ciu.AristasE)
	for i := 0; i < 1000; i++ {
		sp, _, _ := vecino(sol.actual)
		// sp, _, _ := vecino(sol.mejor)
		fsP := FunCostoSolucion(sp, ciu.Distancias, ciu.AristasE,
			sol.normalizador)
		if fsP <= sol.actualC + t {
			c++
			sol.actual = sp // s <- s'
			sol.actualC = fsP
		}
	}
	return float64(c)/float64(1000)
}

func busquedaBinaria(s *Ciudades, t1, t2 float64, sol *Solucion) float64{
	// fmt.Printf("T1, T2: %2.15f\t%2.15f\n", t1, t2)
	tm := (t1+t2)/2
	// fmt.Printf("TM: %2.15f\n", tm)
	if t2 - t1 < EPSILONP {
		return tm
	}
	p := porcentajeAceptados(s, tm, sol)
	// fmt.Printf("P:A: %2.15f\n", p)
	if math.Abs(P - p) < EPSILON { // P
		return tm
	}
	if p > P {
		return busquedaBinaria(s, t1, tm, sol)
	} else {
		return busquedaBinaria(s, tm, t2, sol)
	}
}

func (s *Ciudades) TemperaturaInicial(t float64, sol *Solucion) float64 {
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




// func calculaLote(t float64, ciudades *) (float64, []int) {
func calculaLote(t float64, ciu *Ciudades, sol *Solucion) (float64, []int) {
	c := 0
	i := 0
	r := 0.0
	// s := CopiarCiudades(sol.actual)
	// fs := FunCostoSolucion(s, ciu.Distancias, ciu.AristasE)
	for c < L && i < L*L {
		sP, a, b := vecino(sol.actual)
		// fmt.Println(i)
		fsP := FunCostoSolucion(sP, ciu.Distancias, ciu.AristasE,
			sol.normalizador)
		if fsP <= sol.actualC + t {
			sol.actual = sP // s <- s'
			sol.actualC = fsP
			// fmt.Printf("\n\t(%d,%d) .... randoms\t\n",a, b)
			if fsP < sol.mejorC {
				// fmt.Println(i)				
				sol.mejor = sP
				sol.mejorC = fsP
				printSol(sol.mejorC, sol.mejor)
				fmt.Printf("\n\t(%d,%d) .... randoms\t\n",a, b)
			}
			c++
			r = r + fsP
			// fmt.Printf("R: %2.15f\n", r)
		}
		i++;
	}
	// fmt.Println(s)
	return r/L, sol.actual
}

// func AceptacionPorUmbrales(t float64, sol*Ciudades) []int{
func (c *Ciudades) AceptacionPorUmbrales(t float64, sol *Solucion) []int{
	// fmt.Println("ACEPTACION POR UMBRALES")
	s := c.Id
	p := 0.0
	for t > EPSILON {
		fmt.Println(t)
		q := math.MaxFloat64
		for p < q {
			q = p
			p, s = calculaLote(t, c, sol)
			fmt.Println("\n\n------------------------------------UP")
			fmt.Printf("\nP: %2.15f\nQ: %2.15f\n",p, q)
			fmt.Println("------------------------------------DW\n\n")
			// fmt.Printf("P: %2.15f\n", p)
			// fmt.Println(s)
		}
		t = PHI*t
	}
	return s
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
	costo := FunCostoSolucion(id, dist, aristas, norm)
	return &Solucion{
		// temperatura: 0.0,
		normalizador: norm,
		mejor: id,
		mejorC: costo,
		actual: id,
		actualC: costo,
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
