package funciones

import (
	// "math/rand"
	"math"
	"sort"
	"fmt"
)

type General struct {
	Init []int
	Distancias [][]float64 // Grafica completa
	AristasE []float64 // Aristas en E
}

type Solucion struct {
	normalizador float64 // Normalizador usado 
	mejor []int
	mejorC float64
	actual []int
	actualC float64
	init []int
	initC float64
	i,j int
}

// Solo imprime una ciudad con sus datos
func (g *General) PrintGenData() {
	fmt.Println("GRÁFICA COMPLETA")
	fmt.Printf("\t%2.15f\n\n", g.Distancias)
	fmt.Println("ARISTASE")
	fmt.Printf("\t%2.15f\n\n", g.AristasE)
	fmt.Println("MÁXIMA DISTANCIA")
	fmt.Printf("\t%2.15f\n\n", g.AristasE[len(g.AristasE)-1])
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
}

// Para cada par no ordenado, si la arista está en las distancias (tsp.sql), la
// agregamos a una lista
// Regresa todas las aristas en E
func totalAristas(ciudadesId []int, g *General) []float64 {
	var totalAristas []float64
	for i := 0; i < len(ciudadesId); i++ {
		// ¿j = i?
		for j := 0; j < len(ciudadesId); j++ {
			// Si está en las distancias agregamos
			if g.Distancias[i][j] != 0 {
				totalAristas = append(totalAristas, g.Distancias[i][j])
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
func FunCostoSolucion(id []int, norm float64, g *General) float64 {
	suma := 0.0
	for i := 1; i < len(id); i++ {
		if (g.Distancias[i][i-1] == 0 && g.Distancias[i-1][i] == 0) {
			suma += pesoAumentado(id[i], id[i-1], g.AristasE[len(g.AristasE)-1])
		} else {
			// No sabemos en qué parte de la matriz esta
			suma += g.Distancias[i][i-1] + g.Distancias[i-1][i]
		}
	}
	return suma/norm
}

func porcentajeAceptados(g *General, t float64, sol *Solucion) float64 {
	c := 0
	// fs := FunCostoSolucion(s, ciu.Distancias, ciu.AristasE)
	for i := 0; i < 1000; i++ {
		sp, _, _ := vecino(sol.actual)
		// sp, _, _ := vecino(sol.mejor)
		fsP := FunCostoSolucion(sp, sol.normalizador, g)
		if fsP <= sol.actualC + t {
			c++
			sol.actual = sp // s <- s'
			sol.actualC = fsP
		}
	}
	return float64(c)/float64(1000)
}

func busquedaBinaria(g *General, t1, t2 float64, sol *Solucion) float64{
	// fmt.Printf("T1, T2: %2.15f\t%2.15f\n", t1, t2)
	tm := (t1+t2)/2
	// fmt.Printf("TM: %2.15f\n", tm)
	if t2 - t1 < EPSILONP {
		return tm
	}
	p := porcentajeAceptados(g, tm, sol)
	// fmt.Printf("P:A: %2.15f\n", p)
	if math.Abs(P - p) < EPSILON { // P
		return tm
	}
	if p > P {
		return busquedaBinaria(g, t1, tm, sol)
	} else {
		return busquedaBinaria(g, tm, t2, sol)
	}
}

func TemperaturaInicial(t float64, sol *Solucion, g *General) float64 {
	p := porcentajeAceptados(g, t, sol)
	var t1, t2 float64
	if math.Abs(P - p) <= EPSILONP {
		return t
	}
	if p < P {
		for p < P {
			t = 2*t
			p  = porcentajeAceptados(g, t, sol)
		}
		t1 = t/2
		t2 = t
	} else {
		for p > P {
			t = t/2
			p = porcentajeAceptados(g, t, sol)
		}
		t1 = t
		t2 = 2*t
	}
	return busquedaBinaria(g, t1, t2, sol)
}




// func calculaLote(t float64, ciudades *) (float64, []int) {
func calculaLote(t float64, sol *Solucion, g *General) (float64, []int) {
	c := 0
	i := 0
	r := 0.0
	// s := CopiarCiudades(sol.actual)
	// fs := FunCostoSolucion(s, ciu.Distancias, ciu.AristasE)
	for c < L && i < L*L {
		sP, a, b := vecino(sol.actual)
		// fmt.Println(i)
		fsP := FunCostoSolucion(sP, sol.normalizador, g)
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
func AceptacionPorUmbrales(t float64, sol *Solucion, g *General) []int{
	// fmt.Println("ACEPTACION POR UMBRALES")
	s := g.Init
	p := 0.0
	for t > EPSILON {
		fmt.Println(t)
		q := math.MaxFloat64
		for p < q {
			q = p
			p, s = calculaLote(t, sol, g)
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

func NewSolucion(id []int, g *General) *Solucion {
	norm := GetNormalizador(g.AristasE, id)
	costo := FunCostoSolucion(id, norm, g)
	return &Solucion{
		// temperatura: 0.0,
		normalizador: norm,
		mejor: CopiarCiudades(id),
		mejorC: costo,
		actual: CopiarCiudades(id),
		actualC: costo,
		init: CopiarCiudades(id),
		initC: costo,
	}
}

func NewGeneral(ciudadesId []int) *General {
	var newGeneral = new(General)
	newGeneral.Distancias = completa(ciudadesId)
	newGeneral.AristasE = totalAristas(ciudadesId, newGeneral)
	newGeneral.Init = ciudadesId
	return newGeneral
}
