package funciones

import (
	"math"
	"math/rand"
	"fmt"
)

//Estructura para las ciudades datos de tsp.sql
type Ciudad struct {
	pos int // posicion
	id int  // identificador
	name string // nombre
	country string // pais
	population int // poblacion
	latitude float64 // latitud
	longitude float64 // longitud
}

// Datos generales
type General struct {
	aristasE []float64 // aristas en la gráfica
	normalizador float64 // normalizador
	completa [][] float64 // grafica completa
}

// solucion
type Solucion struct {
	ciudades []Ciudad // configuracion de las ciudades
	eval float64 // evaluacion de las ciudades
	i int // indides por intercambiar
	j int // indides por intercambiar
}

type TSP struct {
	init Solucion // solucion inicial
	actt Solucion // solucion actual
	best Solucion // solucion mejor
	temperatura float64 // temperatura usada
	datos General // datos. General struct
}

// Regresa la suma de las últimas k aristas
// Recibe todas las aristas en E y todas las ciudades
func GetNormalizador(aristasE []float64, ciudadesId []int) float64{
	suma := 0.0
	end := 0
	if len(aristasE) < len(ciudadesId)-1 {
		end = len(aristasE)
	} else {
		end = len(aristasE)-len(ciudadesId)
	}
	for i := len(aristasE)-1; i > end; i-- {
		suma += aristasE[i]
	}
	return suma
}

// Funcion que calcula el costo de una solucion
func FuncionCosto(cis []Ciudad, g *General) float64 {
	suma := 0.0
	for i := 1; i < len(cis); i++ {
		suma += g.completa[cis[i-1].pos][cis[i].pos]
	}
	return suma/g.normalizador
}

// Funcion que saca un vecino.
// Tal vez pueda hacerlo que regrese una solucion vecina
func vecino(cis []Ciudad) []Ciudad {
	copia := copiarCiudades(cis)
	var i = rand.Intn(len(cis))
	var j = rand.Intn(len(cis))
	for i == j {
		i = rand.Intn(len(cis))
	}
	copia[i] = cis[j]
	copia[j] = cis[i]
	return copia
}

// Funcion que regresa el porcentaje de aceptados en una solucion* uso el arreglo,
// no la solucion
func porcentajeAceptados(cis []Ciudad, t float64, g *General) float64{
	c := 0.0
	i := 1 
	s := copiarCiudades(cis)
	fs := FuncionCosto(s, g)
	for (i < 1000){
		sP := vecino(s)
		fsP := FuncionCosto(sP, g)
		if (fsP < fs + t){
			c++
			s = copiarCiudades(sP)
			fs = fsP
		}       
		i++
	}
	return c/1000.0
}

// Funcion que implementa busqueda binaria para obtener una temperatura
func busquedaBinaria(cis []Ciudad, t1, t2 float64, g *General) float64 {
	tm := (t1 + t2)/2.0
	if t2 - t1 < EPSILON {
		return tm
	}
	p := porcentajeAceptados(cis, tm, g)
	if math.Abs(P - p) < EPSILONP {
		return tm
	}
	if p > P {
		return busquedaBinaria(cis, t1, tm, g)
	} else {
		return busquedaBinaria(cis, tm, t2, g)
	}
}

// Funcion que calcula la temperatura inicial, dada la temperatura inicial de 8
func TemperaturaInicial(t float64, cis []Ciudad, g *General) float64 {	
	p := porcentajeAceptados(cis, t, g)
	var t1, t2 float64
	if math.Abs(P - p) <= EPSILONP {
		return t
	}
	if p < P {
		for p < P {
			t = 2*t
			p = porcentajeAceptados(cis, t, g)
		}
		t1 = t/2
		t2 = t
	} else {
		for p > P {
			t = t/2
			p = porcentajeAceptados(cis, t, g)
		}
		t1 = t
		t2 = 2*t        
	}
	return busquedaBinaria(cis, t1, t2, g)
}

// Funcion que calcula el porcentaje de aceptados en un lote
func calculaLote(tsp *TSP) (float64, *TSP) {
	c := 0.0
	r := 0.0
	i := 0
	s := copiarCiudades(tsp.actt.ciudades)
	fs := tsp.actt.eval
	mejor := copiarCiudades(tsp.best.ciudades)
	n := len(tsp.init.ciudades)
	fMejor:= tsp.best.eval
	for c < L && i < L*2*n {
		sP := vecino(s)
		fsP := FuncionCosto(sP, &tsp.datos)
		if fsP < fs + tsp.temperatura {
			i = 0
			s = copiarCiudades(sP)
			fs = fsP
			c++
			r += fs
			if fs < fMejor {
				mejor = copiarCiudades(s)
				fMejor = fs
			}
		}
		i++
	}
	solAct := Solucion{s, fs, -1, -1}
	solBest := Solucion{mejor, fMejor, -1, -1}
	nuevoTSP := TSP{tsp.init, solAct, solBest, tsp.temperatura, tsp.datos}
	return r/L, &nuevoTSP
}

// Funcion que implementa aceptacion por umbrales
func AceptacionPorUmbrales(tsp *TSP) ([]Ciudad, float64){
	mejor := copiarCiudades(tsp.best.ciudades)
	p := 0.0
	for tsp.temperatura > EPSILON {
		var q = math.MaxFloat64	
		for p <= q {
			q = p
			rl, newTSP := calculaLote(tsp)
			p = rl
			mejor = copiarCiudades(newTSP.best.ciudades)
			tsp.best.ciudades = mejor
			tsp.best.eval = newTSP.best.eval
			fmt.Printf("E:%2.15f\n", tsp.best.eval)
		}
		tsp.temperatura = tsp.temperatura * PHI
	}
	fMejor := FuncionCosto(mejor, &tsp.datos)
	return mejor, fMejor
}

func checkVecindario(cis []Ciudad, costo float64, g *General) ([]Ciudad, float64) {
	for i := 0; i < len(cis); i++ {
		for j := 0; j < len(cis); j++ {
			aux := cis[i]
			cis[i] = cis[j]
			cis[j] = aux
			newCosto := FuncionCosto(cis, g)
			if newCosto < costo {
				fmt.Printf("i--%d\nj--%d\n",i,j)
				costo = newCosto
				i = 0
				j = 0
			}
		}
	}
	return cis, costo
}

func NewTSP(ids []int) float64 {
	ciudades := ciudades(ids) // arreglo de todas las ciudades
	aristasE := TotalAristas(ciudades) // las aristas de la gráfica
	max := aristasE[len(aristasE)-1] // máxima distancia
	fmt.Printf("MAX DIST:\t%2.15f\n",max) 
	normalizador := GetNormalizador(aristasE, ids) // normalizador
	fmt.Printf("NORM:\t%2.15f\n",normalizador)
	completa := Completa(ciudades, max) // grafica completa de todas las aristas
	gen := General{aristasE, normalizador, completa}
	costo := FuncionCosto(ciudades, &gen) // funcion de costo
	fmt.Printf("COSTO:\t%2.15f\n", costo)
	ti := TemperaturaInicial(8, ciudades, &gen) // temperatura inicial, inicial
	ini := Solucion{ciudades, costo, -1, -1} // solucion inicial
	act := Solucion{ciudades, costo, -1, -1} // solucion actual
	mej := Solucion{ciudades, costo, -1, -1} // solucion mejor
	tsp := TSP{ini, act, mej, ti, gen}
	fmt.Printf("TEMP INICIAL:\t%2.15f\n", tsp.temperatura)
	mejorCiu, mejorCos := AceptacionPorUmbrales(&tsp) // mejor sol y su eval
	PrettyPrint(mejorCiu)
	mejorCiu, mejorCos = checkVecindario(mejorCiu, mejorCos, &gen)
	fmt.Printf("EVAL MEJOR:\t%2.15f\n", mejorCos)
	return mejorCos
}
