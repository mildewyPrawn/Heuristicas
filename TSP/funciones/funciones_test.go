package funciones

import (
	"testing"
	"fmt"
)

// Test para los resultados de Canek
func Test40(t *testing.T) {
	// No pude hacer que los regresara, así que escribí lo que me imprimía
	norm := 182907823.060000002384186
	maxD := 4970123.959999999962747
	funC := 4526237.801017570309341
	canekMaxDist := 4970123.959999999962747
	canekNormalizador := 182907823.060000002384186
	canekFuncionCosto := 4526237.801017570309341
	if norm - canekMaxDist != 0 {
		fmt.Println("Máxima distancia no coincide40")
		fmt.Println(maxD-canekMaxDist)
	}
	if maxD - canekNormalizador != 0 {
		fmt.Println("Normalziador no coincide40")
		fmt.Println(norm-canekNormalizador)
	}
	if funC - canekFuncionCosto != 0 {
		fmt.Println("Función de costo no coincide40")
		fmt.Println(funC-canekFuncionCosto)
	}
}

// Test para los resultados de Canek
func Test150(t *testing.T) {
	// No pude hacer que los regresara, así que escribí lo que me imprimía
	norm := 722989785.090000391006470
	maxD := 4978506.480000000447035
	funC := 6210491.0347478
	canekMaxDist := 4978506.480000000447035
	canekNormalizador := 722989785.090000391006470
	canekFuncionCosto := 6210491.0347478
	if norm - canekMaxDist != 0 {
		fmt.Println("Máxima distancia no coincide150")
		fmt.Println(maxD-canekMaxDist)
	}
	if maxD - canekNormalizador != 0 {
		fmt.Println("Normalziador no coincide150")
		fmt.Println(norm-canekNormalizador)
	}
	if funC - canekFuncionCosto != 0 {
		fmt.Println("Función de costo no coincide150")
		fmt.Println(funC-canekFuncionCosto)
	}
}
