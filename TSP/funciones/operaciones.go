package funciones

import (
	"math"
	"strconv"
	"fmt"
)

const (
	// Constante del radio de la tierra aproximado
	radio = 6373000
	// medidas que son las que debería ir ajustando
	P = 0.90
	EPSILON = 0.0001
	EPSILONP = 0.0001
	L = 1000
	PHI = .9
)

// Funcion que saca la gráfica completa de las ciudades
func Completa(cis []Ciudad, max float64) [][]float64 {
	var matriz = [][]float64{}
	for i := 0; i < len(cis); i++ {
		adyacentes := make([]float64, len(cis))
		for j := 0; j < len(cis); j++ {
			adyacentes[j] = pesoAumentado(cis[i], cis[j], max)
		}
		matriz = append(matriz, adyacentes)
	}
	return matriz
}

// Funcion que transforma una corrdenada a radianes
func radianes(f float64) float64 {
	return (f*math.Pi)/180
}

// Funcion que calcula la formula A del PDF
func obtenerA (u, v Ciudad) float64 {
	latV := radianes(v.latitude)
	lonV := radianes(v.longitude)
	latU := radianes(u.latitude)
	lonU := radianes(u.longitude)
	sin1 := math.Pow(math.Sin((latV-latU)/2), 2)
	sin2 := math.Pow(math.Sin((lonV-lonU)/2), 2)
	cos1 := math.Cos(latU)
	cos2 := math.Cos(latV)
	return sin1 + cos1 * cos2 * sin2
}

//Funcion para obtener la distancia natural acorde al PDF
func distanciaNatural(u, v Ciudad) float64 {
	a := obtenerA(u, v)
	c := 2 * math.Atan2(math.Sqrt(a),math.Sqrt(1-a))
	return radio * c

}

// Funcion que copia un arreglo de ciudades.
func copiarCiudades(cis []Ciudad) []Ciudad {
	nueva := make([]Ciudad, len(cis))
	for i:= 0; i < len(cis); i++ {
		nueva[i] = cis[i]
	}
	return nueva
}

// Funcion que solo imprime los indices de la solucion
func PrettyPrint(cis []Ciudad) {
	s := ""
	for i := 0; i < len(cis); i++ {
		s += strconv.Itoa(cis[i].id) + " "
	}
	fmt.Println(s)
}
