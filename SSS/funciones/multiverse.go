package funciones

import (
	"fmt"
	"math/rand"
)

type transactores struct {
	id int // id del que está involucrado en la transacción
	deuOacr bool // true si es acreedor, false eoc
	monto int // dinero en cuestión
	aQuien int // id del que le debe o al que le debe
	deQuienes []int // lista de deudores
}

// Quien le debe a quien
type Pair struct {
	d int
	a int
}

type universo struct {
	id int
	acreedores []transactores
	deudores []transactores
	aristas []Pair
	total int
	error int
}

type multiverso struct {
	init universo
	best universo
	mult []universo
}

// Función que crea el primer universo, no tiene aristas
// Recibe a los acreedores y deudores
func PrimerUniverso(a, d map[int]int) universo {
	var acre []transactores
	var deud []transactores
	tot := 0
	for k, v := range a {
		ti := transactores{k, true, v, -1, nil}
		acre = append(acre, ti)
		tot += v
	}
	for k, v := range d {
		ti := transactores{k, false, v, -1, nil}
		deud = append(deud, ti)
	}
	u := universo{0, acre, deud, nil, tot, 0}
	// fmt.Println(u)
	return u
}


// problemas porque usa una copia, que pedo
func calculaError(u universo) int {
	error := 0
	for _,a := range u.acreedores {
		for _,d := range a.deQuienes {
			fmt.Println(d)
			for _, td := range u.deudores {
				if td.id == d {
					fmt.Printf(">>>%d\n",a.monto)
					a.monto += td.monto
				}
				error += abs(a.monto)
				fmt.Printf("EEEEE%d\n", error)
			}
		}
		fmt.Println(a.deQuienes)
	}
	fmt.Println(u)
	fmt.Printf(">>>E%d\n",error)
	return error
}

// Función que crea aristas entre los deudores y acreedores
// Recibe un universo (El primer universo) y regresa un universo nuevo, con error
// y por lo tanto con aristas
func CreaAristas(u universo) universo {
	var aris []Pair
	
	for i := 0; i < len(u.deudores); i++ {
		var ia = rand.Intn(len(u.acreedores))
		u.deudores[i].aQuien = u.acreedores[ia].id
		u.acreedores[ia].deQuienes = append(u.acreedores[ia].deQuienes,u.deudores[i].id)
	}
	for i := 0; i < len(u.deudores); i++ {
		p := Pair{u.deudores[i].id, u.deudores[i].aQuien}
		aris = append(aris, p)
	}
	// TODO calcular el error
	v := universo {1, u.acreedores, u.deudores, aris, u.total, -1}
	fmt.Println(v)
	calculaError(v)
	return v
}
