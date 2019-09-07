package funciones

import (
	"math"
)

// Constante del radio de la tierra aproximado
const radio = 6373000

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
