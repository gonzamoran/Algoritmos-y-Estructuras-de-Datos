package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	linea := bufio.NewScanner(os.Stdin)
	for linea.Scan() {

		expresion := linea.Text()
		token := Tokenizar(expresion)
		postfija := ConversionNotacion(token)

		var resultado []string
		for !postfija.EstaVacia() {
			resultado = append(resultado, postfija.Desencolar())
		}
		fmt.Println(strings.Join(resultado, " "))
	}
}
