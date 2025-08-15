package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const ruta1 = "archivo1.in"
const ruta2 = "archivo2.in"

func ordenarVector(vector []int) {
	ejercicios.Seleccion(vector)
	for i := 0; i < len(vector); i++ {
		fmt.Println(vector[i])
	}
}

func main() {

	var numeros_arch1 []int
	var numeros_arch2 []int

	archivo1, err := os.Open(ruta1)
	archivo2, err1 := os.Open(ruta2)

	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", ruta1, err)
	}
	if err1 != nil {
		fmt.Printf("Error %v al abrir el archivo %s", ruta2, err1)
	}

	defer archivo1.Close()
	defer archivo2.Close()

	s1 := bufio.NewScanner(archivo1)
	s2 := bufio.NewScanner(archivo2)

	for s1.Scan() {
		numeros1, _ := strconv.Atoi(s1.Text())
		numeros_arch1 = append(numeros_arch1, numeros1)
	}
	for s2.Scan() {
		numeros2, _ := strconv.Atoi(s2.Text())
		numeros_arch2 = append(numeros_arch2, numeros2)
	}

	resultado := ejercicios.Comparar(numeros_arch1, numeros_arch2)

	if resultado == 1 {
		ordenarVector(numeros_arch1)
	} else if resultado == -1 || resultado == 0 {
		ordenarVector(numeros_arch2)
	}
}
