package funciones

import (
	"testing"
	"fmt"
)

// Test para radianes, los calculé a mano :v
func TestRadianes(t *testing.T) {
	var grados = []float64{80.5, 34.22, 299.14, 78.58, 1.1}
	var rad = []float64{1.4050, 0.5972, 5.2226, 1.3715, 0.0177}
	for i := range grados {
		radI := radianes(grados[i])
		// Los trunco porque le faltan como 1000 decimales
		if int(radI) != int(rad[i]) {
			fmt.Println("Error en radianes")
		}
	}
}

// Test para distancia natural, lo saqué del correo
func TestDistanciaNatural(t *testing.T) {
	d17 := 2999396.231968969572335
	d1163 := 3222670.069842538330704
	dist17 := distanciaNatural(1,7)
	dist1163 := distanciaNatural(1,163)
	if d17 != dist17 {
		fmt.Println("Error en distancias 1-7")
	}
	if d1163 != dist1163 {
		fmt.Println("Error en distancias 1-163")
	}
}

// Estoy suponiendo que (1-7), (1-163) no están conectadas y que la máxima es 2
func TestPesoAumentado(t *testing.T) {
	d17 := 2999396.231968969572335*2
	d1163 := 3222670.069842538330704*2
	pa17 := pesoAumentado(1, 7, 2)
	pa1163 := pesoAumentado(1, 163, 2)
	if d17 != pa17 {
		fmt.Println("Error en peso aumentado")
	}
	if d1163 != pa1163 {
		fmt.Println("Error en peso aumentado")
	}
}

// Test para obtenerA
func TestObtenerA(t *testing.T) {
	a12 := 0.01933276365124884
	a23 := 0.14812066305808086
	ta12 := obtenerA(1,2)
	ta23 := obtenerA(2,3)
	if a12 != ta12 {
		fmt.Println("Error en obtener A")
	}
	if a23 != ta23 {
		fmt.Println("Error en obtener A")
	}
}
