package main

import (
	"fmt"
	"math"
	"crypto/rand"
	"math/big"
	city "github.com/Heuristicas/TSP/funciones"
)

type ciudades = city.Ciudades
var ciudadeser interface{} = "hola"

const TAMLOTE = 100
const PHI = .75
const EPSILON = .0000001
const P = .90
const EPSILONP = .0000001

// Genera un número random
 // Regresa un número entre [0, i)
func randInt(i int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(i)))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return int(n)
}

// Copia un arreglo
// Regresa la copia del arreglo
func copiarCiudades(actual []int) []int {
	nuevo := make([]int, len(actual))
	for i := 0; i < len(actual); i++ {
		nuevo[i] = actual[i]
	}
	return nuevo
}

// Swapea un arreglo
// Recibe dos índices
// Regresa el arreglo con los índices swapeados
func swap(i, j int, actual []int) []int {
	nuevo := copiarCiudades(actual)
	nuevo[j] = actual[i]
	nuevo[i] = actual[j]
	return nuevo;
}

// Regresa índices distintos 
// func Vecino(actual []int) ([]int, int, int) {
func Vecino(actual []int) []int {
	i := randInt(len(actual))
	j := randInt(len(actual))
	for i == j {
		i = randInt(len(actual))
	}
	nuevo := swap(i, j, actual)
	return nuevo
	// return nuevo, i, j
}






// QUIZA SEA BUENO PASAR TODO A FUNCIONES (?)









func calculaLote(temperatura float64, ciudad *ciudades) (float64, []int) {
	c := 0
	r := 0.0
	i := 0
	norm := ciudad.GetNormalizador()
	dist := ciudad.GetDistancias()
	for c < TAMLOTE || i <= TAMLOTE*TAMLOTE {
		sPrima := Vecino(ciudad.Id)
		fsP := city.FunCostoSolucion(sPrima, dist, ciudad.GetAristasE())
		if fsP/norm <= ciudad.Costo/norm + temperatura {
			ciudad.SetId(sPrima)
			c++
			ciudad.FunCosto()
			r = r + ciudad.GetDistTotal()/norm
		}
		i++
	}
	return (r/TAMLOTE), ciudad.Id
}

// Creo que ya está lo de calcular LOTE.
// tuve que cambiar ciudades por Ciudades ***** quizá eso sea importante después

func aceptacionPorUmbrales(temperatura float64, ciudad *ciudades) []int{
	s := ciudad.GetId()
	p := 0.0
	for temperatura > EPSILON {
		q := math.MaxFloat64
		for p <= q {
			q = p
			//problemas con las firmas y lo que regreso
			p, s = calculaLote(temperatura, ciudad) //oh oh
		}
		temperatura = PHI*temperatura
	}
	// ¿Mejor solucion?
	return s
}

// Necesito recibir una perra ciudad
// func PorcentajeAceptados(s []int, t float64) float64 {
func PorcentajeAceptados(ciudad *ciudades, temperatura float64) float64 {
	norm := ciudad.GetNormalizador()
	dist := ciudad.GetDistancias()
	c := 0.0
	// ¿Qué es N?
	for i := 0; i < len(ciudad.Id); i++ {
		sP := Vecino(ciudad.Id)
		fsP := city.FunCostoSolucion(sP, dist, ciudad.GetAristasE())
		if (fsP/norm <= ciudad.Costo/norm + temperatura) {
			c++
			ciudad.SetId(sP)
		}
	}
	return c/float64(len(ciudad.Id))
}

// Mi pMayos (PDF) es P en cont
// func BusquedaBinaria(s []int, t1, t2 float64) float64 {
func BusquedaBinaria(ciudad *ciudades, t1, t2 float64) float64 {
	// s := ciudad.Id
	// return 0.0
	tm := (t1 + t2)/2.0
	if (t2 - t1 < EPSILONP) {
		return tm
	}
	p := PorcentajeAceptados(ciudad, tm)
	if (math.Abs(P - p) < EPSILONP) {
		return tm
	}
	if (p > P) {
		return BusquedaBinaria(ciudad, t1, tm)
	} else {
		return BusquedaBinaria(ciudad, tm, t2)
	}
}

func TemperaturaInicial(ciudad *ciudades, t float64) float64 {
	p := PorcentajeAceptados(ciudad, t)
	t1 := 0.0
	t2 := 0.0
	if math.Abs(P - p) <= EPSILONP {
		return t
	}
	if p < P {
		for p < P {
			t = 2.0*t
			p = PorcentajeAceptados(ciudad, t)
		}
		t1 = t*2.0
		t2 = t
	} else {
		for p > P {
			t = t/2.0
			p = PorcentajeAceptados(ciudad, t)
		}
		t1 = t
		t2 = 2.0*t
	}
	return BusquedaBinaria(ciudad, t1, t2)
}

func main() {
	var ciudades40 = []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,
		332,333,489,490,491,492,493,496,652,653,654,656,657,792,815,816,
		817,820,978,979,980,981,982,984}

	c := city.NewCiudades(ciudades40)
	c.FunCosto()
	// c.PrintCiudad()

	// copyCiudades := copiarCiudades(ciudades40)
	// fmt.Println(copyCiudades)
	
	// randArray := swap(0,1, ciudades40)
	// fmt.Println(randArray)

	
	
	sol := aceptacionPorUmbrales(13, c)

	// fmt.Println(sol)

	// vecinoC40, i, j := Vecino(ciudades40)
	vecinoC40 := Vecino(ciudades40)
	// fmt.Printf("%d----%d\n", i, j)
	fmt.Println(vecinoC40)

	v := city.NewCiudades(vecinoC40)
	v.FunCosto()
	// v.PrintCiudad()
	// costoV := v.GetDistTotal()
	// fmt.Println("------------------------%2.15f---------------", costoV)
	// fmt.Println(v.costo/v.normalizador)
	
}
