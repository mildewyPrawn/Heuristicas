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
	deudor int
	acreedor int
}

type universo struct {
	id int
	acreedores []transactores
	deudores []transactores
	aristas []Pair
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
	for k, v := range a {
		ti := transactores{k, true, v, -1, nil}
		acre = append(acre, ti)
	}
	for k, v := range d {
		ti := transactores{k, false, v, -1, nil}
		deud = append(deud, ti)
	}
	u := universo{0, acre, deud, nil, 0}
	fmt.Println(u)
	return u
}

func calculaError() {
	
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
	v := universo {1, u.acreedores, u.deudores, aris, -1}
	fmt.Println(v)
	return v
}
