package funciones

import (
	// "math/rand"
	"math"
	"sort"
	"fmt"
)

type miniSol struct {
	s []int
	fs float64
}

type General struct {
	Init []int
	Distancias [][]float64 // Grafica completa
	AristasE []float64 // Aristas en E
}

type Solucion struct {
	normalizador float64 // Normalizador usado
	init miniSol
	best miniSol
	newS miniSol
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
	fmt.Printf("NORMALIZADOR: %2.15f\n", s.normalizador)
	fmt.Printf("INIT\n")
	fmt.Printf("C: %2.15f\n", s.init.fs)
	fmt.Println(s.init.s)
	fmt.Printf("BEST\n")
	fmt.Printf("C: %2.15f\n", s.best.fs)
	fmt.Println(s.best.s)
	fmt.Printf("NEWS\n")
	fmt.Printf("C: %2.15f\n", s.newS.fs)
	fmt.Println(s.newS.s)
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
	s := sol.init.s
	fs := sol.init.fs
	for i := 0; i < 1000; i++ {
		sp, _, _ := vecino(s)
		fsP := FunCostoSolucion(sp, sol.normalizador, g)
		if fsP <= fs + t {
			c++
			// s <- s'
			s = CopiarCiudades(sp)
			fs = fsP
		} // No es mejor, pero no importa, porque la volvemos a calcular
	}
	return float64(c)/float64(1000)
}

func busquedaBinaria(g *General, t1, t2 float64, sol *Solucion) float64{
	fmt.Println("EMPIEZA BS")
	tm := (t1+t2)/2.0
	if t2 - t1 < EPSILONP {
		return tm
	}
	p := porcentajeAceptados(g, tm, sol)
	if math.Abs(P - p) < EPSILONP {
		return tm
	}
	if p > P {
		return busquedaBinaria(g, t1, tm, sol)
	} else {
		return busquedaBinaria(g, tm, t2, sol)
	}
}

func TemperaturaInicial(t float64, sol *Solucion, g *General) float64 {
	fmt.Println("Empieza temp")
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
	return busquedaBinaria(g, t1, t2, sol) // COMO 13 segundos en calcular TI
}

func calculaLote(t float64, sol *Solucion, g *General) (float64, []int) {
	c := 0
	i := 0
	r := 0.0
	s := sol.newS.s
	fs := sol.newS.fs
	for c < L && i < L*L {
		sP, _, _ := vecino(s)
		// fmt.Println(i)
		fsP := FunCostoSolucion(sP, sol.normalizador, g)
		if fsP <= fs + t {
			// s <- s'
			s = CopiarCiudades(sP)
			fs = fsP
			// printSol(fs, s)
			// fmt.Printf("\n\t(%d,%d) .... randoms\t\n",a, b)
			if fsP < sol.best.fs {
				printSol(fs, s)
				sol.best.s = CopiarCiudades(sP)
				sol.best.fs = fsP
				// printSol(sol.mejorC, sol.mejor)
				// fmt.Printf("\n\t(%d,%d) .... randoms\t\n",a, b)
			}
			c++
			r = r + fsP
			// fmt.Printf("R: %2.15f\n", r)
		}
		i++;
	}
	// fmt.Println(s)
	return r/L, s
}

func AceptacionPorUmbrales(t float64, sol *Solucion, g *General) []int{
	// fmt.Println("ACEPTACION POR UMBRALES")
	s := sol.init.s
	p := 0.0
	for t > EPSILON {
		// fmt.Println(t)
		q := math.MaxFloat64
		for p < q {
			q = p
			p, s = calculaLote(t, sol, g)
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
	fmt.Println(g.AristasE)
	costo := FunCostoSolucion(id, norm, g)
	return &Solucion{
		// temperatura: 0.0,
		normalizador: norm,
		init: miniSol {
			s: CopiarCiudades(id),
				fs: costo,
			},
		best: miniSol {
			s: CopiarCiudades(id),
				fs: costo,
			},
		newS: miniSol {
			s: CopiarCiudades(id),
				fs: costo,
			},
	}
}

func NewGeneral(ciudadesId []int) *General {
	var newGeneral = new(General)
	newGeneral.Distancias = completa(ciudadesId)
	newGeneral.AristasE = totalAristas(ciudadesId, newGeneral)
	// fmt.Println(newGeneral.AristasE)
	newGeneral.Init = ciudadesId
	return newGeneral
}
