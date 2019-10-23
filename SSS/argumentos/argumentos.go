package argumentos

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

// func Leer_grafica(path string) []int {
func Leer_grafica(path string) int {
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
		fmt.Println(clases[i])
	}
	return 8
}
