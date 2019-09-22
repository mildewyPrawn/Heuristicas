package funciones

import (
	"math"
	"math/rand"
	// "fmt"
)

const (
	// Constante del radio de la tierra aproximado
	radio = 6373000
	// medidas que son las que debería ir ajustando
	L = 1000
	// L = 500
	PHI = .9
	EPSILON = .0001
	EPSILONP = .0001
	P = .90
	// 525669.094937816378661 L = 1000; L*10 T = 21387....
	// 506144.423279115348123 L = 1000; L*L T = 9202....
	// 535995.873955731629394 L = 100; L*L T = 7454...
	// 521889.019054602307733 L = 1000; L*L T = 10032... SEED = 127
	// 493138.574513699917588 L = 1000; L*L T = 8695... SEED = 1193
	// 342184.580924343201332 L = 500; L*L T = 6636... SEED = 1195
	// 315208.333823074062821 L = 1000; L*L T = 9103... SEED = 1195
	// 335569.438928690564353                           S = 4
)


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

// Copia un arreglo de ciudades, (O de enteros)
// Regresa una copia de actual
func CopiarCiudades(actual []int) []int {
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
	nuevo := CopiarCiudades(ciudades)
	nuevo[j] = ciudades[i]
	nuevo[i] = ciudades[j]
	return nuevo
}

// Obtiene un vecino en la gráfica
// Recibe un arreglo de ciudades
// Regresa un vecino en la grafica de actual
func vecino(actual []int) ([]int, int, int) {
// func vecino(actual []int) []int {
	i := rand.Intn(len(actual))
	j := rand.Intn(len(actual))
	// fmt.Printf("I: %d\tJ: %d", i, j)
	for i == j {
		i = rand.Intn(len(actual))
	}
	nuevo := swap(i,j, actual)
	return nuevo, i, j
}
