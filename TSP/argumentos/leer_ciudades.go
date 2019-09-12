package argumentos

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// const PATH = "../ejemplos/"
const CIUDADES = "ciudades"

// Leemos el archivo
// Regresa las ciudades en un slice
// Regresa el nombre "ciudadesXX" 
func Leer(path string) ([]int, string){
	archivo, err := os.Open(path)

	if (err != nil) {
		fmt.Println("TSP: Error al leer el archivo.")
	}
	scanner := bufio.NewScanner(archivo)
	
	linea := ""
	for scanner.Scan() {
		linea += scanner.Text() + "\n"
	}

	linea = strings.TrimSpace(linea) //quitar \n
	ciudades := strings.Split(linea, ",") //separar por ,
	num := len(ciudades)

	var id []int
	for i := 0; i < len(ciudades); i++ {
		s, _ := strconv.Atoi(ciudades[i])
		id = append(id, s)
	}
	return id, CIUDADES+strconv.Itoa(num)
}
