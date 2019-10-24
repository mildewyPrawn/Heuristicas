package argumentos

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
)

// func Leer_grafica(path string) []int {
func Leer_grafica(path string) int {
	m := make(map[int]int)
	archivo, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error al leer el archivo: %s", path)
	}
	scanner := bufio.NewScanner(archivo)
	linea := ""
	for scanner.Scan() {
		linea += scanner.Text() + "\n"
	}
	linea = strings.TrimSpace(linea) // quitar \n
	clases := strings.Split(linea, "\n") // separar lineas
	for i := 0; i < len(clases); i++ {
		verts := strings.Split(clases[i], "       "); // separar \t
		key, _ := strconv.Atoi(verts[0])// deudor
		if _, ok := m[key]; ok {
			m[key] += 0
		} else {
			m[key] = 0
		}
		deudas := strings.Split(verts[1],",") // acredores
		for j := 0; j < len(deudas); j++ {
			acredores := strings.Split(deudas[j], ":")
			// otherKey, _ := strconv.Atoi(acredores[0]) // acredor i
			debt, _ := strconv.Atoi(acredores[1]) // deuda de key a otherKey
			val := m[key]
			m[key] = val - debt
			
			// fmt.Printf("[%d]\t[%s]--[%s]\n", key, acredores[0], acredores[1])
		}
		// fmt.Println(deudas)
		fmt.Println(m)
	}
	// fmt.Println(m)
	return 8
}
