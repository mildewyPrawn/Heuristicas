package funciones

import (
	"testing"
	// "fmt"
	// "math/rand"
)

var ciudades40 = []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,332,333,
	489,490,491,492,493,496,652,653,654,656,657,792,815,816,817,820,978,979,
	980,981,982,984}

var ciudades150 = []int{1,2,3,4,5,6,7,8,9,11,12,14,16,17,19,20,22,23,25,26,27,74,
	75,77,163,164,165,166,167,168,169,171,172,173,174,176,179,181,182,183,
	184,185,186,187,297,326,327,328,329,330,331,332,333,334,336,339,340,343,
	344,345,346,347,349,350,351,352,353,444,483,489,490,491,492,493,494,495,
	496,499,500,501,502,504,505,507,508,509,510,511,512,520,652,653,654,655,
	656,657,658,660,661,662,663,665,666,667,668,670,671,673,674,675,676,678,
	792,815,816,817,818,819,820,821,822,823,825,826,828,829,832,837,839,840,
	978,979,980,981,982,984,985,986,988,990,991,995,999,1001,1003,1004,1037,
	1038,1073,1075}

func TestNormalizador(t *testing.T) {
	norm40 := 182907823.060000002384186
	norm150 := 722989785.090000391006470
	Init(ciudades40)
	Init(ciudades150)
	norm40Test := Normalizador()
	norm150Test := Normalizador()
	if norm40 != norm40Test {
		t.Error("No es lo mismo en Normalizador 40\t", norm40Test)		
	}
	if norm150 != norm150Test {
		t.Error("No es lo mismo en Normalizador 150\t", norm150Test)
	}
}

func TestFunCosto(t *testing.T) {
	funCosto40 := 4526237.801017570309341
	funCosto150 := 6210491.034747813828290
	Init(ciudades40)
	Init(ciudades150)
	funCosto40Test := FunCosto(ciudades40, completa(ciudades40))
	funCosto150Test := FunCosto(ciudades150, completa(ciudades150))
	if funCosto40 != funCosto40Test {
		t.Error("No es lo mismo en FunCost 40\t", funCosto40Test)		
	}
	if funCosto150 != funCosto150Test {
		t.Error("No es lo mismo en FunCost 150\t", funCosto150Test)
	}
}

func TestGetMaxDist(t *testing.T) {
	max40 := 4970123.959999999962747
	max150 := 4978506.480000000447035
	Init(ciudades40)
	Init(ciudades150)
	max40Test := MaximaDist
	max150Test := MaximaDist
	if max40 != max40Test {
		t.Error("No es lo mismo en MaxDist 40\t", max40Test)
	}
	if max150 != max150Test {
		t.Error("No es lo mismo en MaxDist 150\t", max150Test)
	}
}

func TestLatLon(t *testing.T) {
	
}
