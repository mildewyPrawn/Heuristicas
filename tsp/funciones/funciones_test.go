package funciones

import (
	"testing"
	"fmt"
	"math/rand"
)

func TestLlenaListaL(t *testing.T) {
	ciudadesPrueba := []int{1,3,5,7,9}
	lista := LlenaListaL(ciudadesPrueba)
	// son puros 0.0 porque no tengo las distancias, getDistancia regresa 0.0
	listaChida := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
	if lista == nil {
		t.Error("La lista es nil", listaChida)
	}
	if len(lista) != len(listaChida) {
		t.Error("No miden lo mismo :think:", len(listaChida), len(lista))
	}
	for i := range lista {
		if lista[i] != listaChida[i] {
			t.Error("No son la misma", listaChida)
		}
	}
	fmt.Println("Corregir con otras distancias")
	fmt.Println(listaChida)
}

func TestNormalizador(t *testing.T) {
	res := make([]float64, 30)
	suma := 0.0
	for i := 0; i < len(res)-1; i++ {
		r := rand.Float64()
		if (r == 0.0) {
			r = 54.5
		} else {
			r *= 1000
		}
		suma += r
		res[i] = r
	}
	res[len(res)-1] = 5.5
	var normalizador float64
	normalizador = Normalizador(res)
	if suma != normalizador {
		t.Error("No es la misma suma.", suma, normalizador)		
	}
	fmt.Printf("%f, %f\n", suma, normalizador)
}

func TestFuncionCostoSuma(t *testing.T) {
	ciudadesPrueba := []int{1,3,5,7,9}
	suma := FuncionCostoSuma(ciudadesPrueba)
	// son puros 0.0 porque no tengo las distancias, getDistancia regresa 0.0
	//corregir con otras distancias
	if suma != 0.0 {
		t.Error("No son 0 hehe.", suma)
	}
	fmt.Printf("Corregir con otras distancias\t%f\n", suma)
}

func TestFuncionCosto(t *testing.T) {
	fmt.Println("Error: falta poner pesos para poder probar")
}
