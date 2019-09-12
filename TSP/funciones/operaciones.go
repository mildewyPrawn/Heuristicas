package funciones

import (
	"math"
	"math/big"
	"crypto/rand"
)

// Constante del radio de la tierra aproximado
const radio = 6373000
// medidas que son las que debería ir ajustando
const L = 300
const PHI = .75
const EPSILON = .0001
const EPSILONP = .0001
const P = .90

// Convierte una coordenada a radianes
// Regresa la coordenada en radianes
func radianes(coordenada float64) float64 {
	return (coordenada*math.Pi)/180
}

// Saca la distancia NATURAL
// Regresa la distancia natural entre dos ciudades dadas por su ID
func distanciaNatural(u, v int) float64 {
	a := obtenerA(u, v)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return float64(radio) * c
}

// Regresa la formula 'A' del pdf
func obtenerA (u, v int) float64 {
	latV, lonV := obtenerLatLon(v)
	latU, lonU := obtenerLatLon(u)
	sin1 := math.Pow(math.Sin((latV-latU)/2), 2)
	sin2 := math.Pow(math.Sin((lonV-lonU)/2), 2)
	cos1 := math.Cos(latU)
	cos2 := math.Cos(latV)
	return sin1 + cos1 * cos2 * sin2
}

// Regresa el peso aumentado entre dos ciudades
// El peso aumentado es la distancia natural * maxima distancia
func pesoAumentado(i, j int, max float64) float64 {
	dist := distanciaNatural(i, j)
	return dist * max
}

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

// Copia un arreglo de ciudades, (O de enteros)
// Regresa una copia de actual
func copiarCiudades(actual []int) []int {
	nuevo := make([]int, len(actual))
	for i := 0; i < len(actual); i++ {
		nuevo[i] = actual[i]
	}
	return nuevo
}

// Swapea un arreglo
// Recibe dos indices para swapear
// Regresa el arreglo con los indices intercambiados
func swap(i, j int, ciudades []int) []int {
	nuevo := copiarCiudades(ciudades)
	nuevo[j] = ciudades[i]
	nuevo[i] = ciudades[j]
	return nuevo
}

// Obtiene un vecino en la gráfica
// Recibe un arreglo de ciudades
// Regresa un vecino en la grafica de actual
func vecino(actual []int) []int {
	i := randInt(len(actual))
	j := randInt(len(actual))
	for i == j {
		i = randInt(len(actual))
	}
	nuevo := swap(i,j, actual)
	return nuevo
}
