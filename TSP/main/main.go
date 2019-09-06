package main

import (
	"fmt"
	"github.com/Heuristicas/TSP/funciones"
)


func main() {
	var ciudades40 = []int{1,2,3,4,5,6,7,75,163,164,165,168,172,327,329,331,
		332,333,489,490,491,492,493,496,652,653,654,656,657,792,815,816,
		817,820,978,979,980,981,982,984}
	
	funciones.Init(ciudades40)
	fmt.Println(funciones.ID)
	// fmt.Println(funciones.Distancias)
	// fmt.Println(funciones.AristasE)
	fmt.Println("40 CIUDADES")
	fmt.Printf("MAXIMA DISTANCIA\t%2.15f\n",funciones.MaximaDist)
	fmt.Printf("NORMALIZADOR\t%2.15f\n",funciones.Normalizador())
	fmt.Printf("FUNCION DE COSTO\t%2.15f\n",funciones.FunCosto(ciudades40))
	// costo40, suma40 := funciones.FunCosto(ciudades40)
	// fmt.Printf("FUNCION DE COSTO\t%2.15f\t%2.15f\n",costo40, suma40)

	var ciudades150 = []int{1,2,3,4,5,6,7,8,9,11,12,14,16,17,19,20,22,23,25,
		26,27,74,75,77,163,164,165,166,167,168,169,171,172,173,174,176,
		179,181,182,183,184,185,186,187,297,326,327,328,329,330,331,332,
		333,334,336,339,340,343,344,345,346,347,349,350,351,352,353,444,
		483,489,490,491,492,493,494,495,496,499,500,501,502,504,505,507,
		508,509,510,511,512,520,652,653,654,655,656,657,658,660,661,662,
		663,665,666,667,668,670,671,673,674,675,676,678,792,815,816,817,
		818,819,820,821,822,823,825,826,828,829,832,837,839,840,978,979,
		980,981,982,984,985,986,988,990,991,995,999,1001,1003,1004,1037,
		1038,1073,1075}
	funciones.Init(ciudades150)
	fmt.Println(funciones.ID)
	// fmt.Println(funciones.Distancias)
	// fmt.Println(funciones.AristasE)
	fmt.Println("150 CIUDADES")
	fmt.Printf("MAXIMA DISTANCIA\t%2.15f\n",funciones.MaximaDist)
	fmt.Printf("NORMALIZADOR\t%2.15f\n",funciones.Normalizador())
	fmt.Printf("FUNCION DE COSTO\t%2.15f\n",funciones.FunCosto(ciudades150))
	// costo150, suma150 := funciones.FunCosto(ciudades40)
	// fmt.Printf("FUNCION DE COSTO\t%2.15f\t%2.15f\n",costo150, suma150)

}
