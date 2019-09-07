var ID []int
var Distancias [][]float64
var AristasE []float64
var MaximaDist float64


// Calcula la funcion de costo, recibe los ID's de las ciudades, las distancias y
// la máxima distancia.
// func FunCosto(ciudadesID []int) (float64, float64) {
func FunCosto(ciudadesID []int) float64 {
	suma := 0.0
	for i := 1; i < len(ciudadesID); i++ {
		if (Distancias[i][i-1]) == 0 && Distancias[i-1][i] == 0 {
			suma += PesoAumentado(ciudadesID[i], ciudadesID[i-1])
		} else {
			suma += Distancias[i][i-1] + Distancias[i-1][i]
		}
	}
	// return suma/Normalizador(), suma
	return suma/Normalizador()
}

// Regresa la suma de las últimas k aristas
func Normalizador() float64 {
	suma := 0.0
	end := len(AristasE)-len(ID)
	for i := len(AristasE)-1; i > end; i-- {
		suma += AristasE[i]
	}
	return suma
}

// Para cada par no ordenado si la arista está en distancias(en tsp.sql) la
// agregamos a una lista, la cual está ordenada de menor a mayor.
func totalAristas(distancias [][]float64, ciudades []int) []float64{
	var totalAristasE []float64
	for i := 0; i < len(ciudades); i++ {
		for j := 0; j < len(ciudades); j++ {
			if distancias[i][j] != 0 {
				totalAristasE = append(totalAristasE,
					distancias[i][j])
			}
		}
	}
	sort.Float64s(totalAristasE) // ordenado de menor a mayor
	return totalAristasE
}

func Init(ciudadesID []int) {
	ID = ciudadesID
	Distancias = completa(ciudadesID)
	AristasE = totalAristas(Distancias, ID)
	MaximaDist = AristasE[len(AristasE)-1]
}
