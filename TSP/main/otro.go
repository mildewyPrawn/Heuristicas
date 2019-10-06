package main

import (
	"fmt"
	city "github.com/Heuristicas/TSP/funciones"
	arg "github.com/Heuristicas/TSP/argumentos"
	"os"
	"math/rand"
)

func  main() {  

	var ciudades40 = []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,332,333,489,490,491,492,493,496,652,653,654,656,657,792,815,816,817,820,978,979,980,981,982,984}
	// var ciudades40 = []int{1,2,3,4,5,6,7,8,9,11,12,14,16,17,19,20,22,23,25,26,27,28,74,75,151,163,164,165,166,167,168,169,171,172,173,174,176,179,181,182,183,184,185,186,187,297,326,327,328,329,330,331,332,333,334,336,339,340,343,344,345,346,347,349,350,351,352,444,483,489,490,491,492,493,494,495,496,499,500,501,502,504,505,507,508,509,510,511,512,520,652,653,654,655,656,657,658,660,661,662,663,665,666,667,668,670,671,673,674,675,676,678,814,815,816,817,818,819,820,821,822,823,825,826,828,829,832,837,839,840,978,979,980,981,982,984,985,986,988,990,991,995,999,1001,1003,1004,1037,1038,1073,1075}
	ciudades := ciudades(ciudades40)

	totAristas := totalAristas(ciudades) // aristasE
	maxDist := totAristas[len(totAristas)-1] // maxDist
	fmt.Printf("MAXD:\t%2.15f\n", maxDist)

	normalizador := GetNormalizador(totAristas, ciudades40) // normalizador
	fmt.Printf("NORM:\t%2.15f\n", normalizador)

	completa := completa(ciudades, maxDist) // grafica completa
	
	gen := General{totAristas, normalizador, completa}

	costo := funcionCosto(ciudades, &gen) // funcion costo
	fmt.Printf("COSTO:\t%2.15f\n", costo)
	for i := 0; i < 1000; i++ {
		seed := i
		fmt.Printf("SEED:\t%d\n", seed)
		// seed := 15 // semilla
		// 15 -> 0.39
		// 51 -> 0.42
		// 127 -> 0.43
		// 89 -> 0.38
		// 123456789 -> 0.45
		// 149 -> 
		rand.Seed(int64(seed))
		
		tempInit := temperaturaInicial(8, ciudades, &gen)
		
		ini := Solucion{ciudades, costo, -1, -1}
		act := Solucion{ciudades, costo, -1, -1}
		mej := Solucion{ciudades, costo, -1, -1}
		
		tsp := TSP{ini, act, mej, tempInit, gen}
		fmt.Printf("TI:\t%2.15f\n", tsp.temperatura)
		
		mejorCiudades, costoMejor := aceptacionPorUmbrales(&tsp)
		prettyPrint(mejorCiudades)
		// fmt.Println(mejorCiudades)
		fmt.Printf("MEJOR:\t%2.15f\n", costoMejor)
	}
}
