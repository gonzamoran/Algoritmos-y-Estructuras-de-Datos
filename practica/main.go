package main

import (
	"fmt"
	Lista "tdas/lista"
	Pila "tdas/pila"
)

func main() {
	pila := Pila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	fmt.Println("Pila original")
	for !pila.EstaVacia() {
		fmt.Println(pila.Desapilar())
	}
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)

	insertarEnPos(pila, 70, 2)
	fmt.Println("Pila despues de insertar 70 en la posicion 2:")
	for !pila.EstaVacia() {
		fmt.Println(pila.Desapilar())
	}
	// Ejemplo 1: Cadenas anagramas
	s1 := "amor"
	s2 := "roma"
	resultado := sonAnagramas(s1, s2)
	fmt.Printf("¿'%s' y '%s' son anagramas? %v\n", s1, s2, resultado)

	// Ejemplo 2: Cadenas no anagramas
	s1 = "hola"
	s2 = "mundo"
	resultado = sonAnagramas(s1, s2)
	fmt.Printf("¿'%s' y '%s' son anagramas? %v\n", s1, s2, resultado)

	// Ejemplo 3: Cadenas con diferentes longitudes
	s1 = "abc"
	s2 = "abcd"
	resultado = sonAnagramas(s1, s2)
	fmt.Printf("¿'%s' y '%s' son anagramas? %v\n", s1, s2, resultado)

	miLista := Lista.CrearListaEnlazada[int]()
	miLista.InsertarUltimo(4)
	miLista.InsertarUltimo(2)
	miLista.InsertarUltimo(4)
	miLista.InsertarUltimo(3)
	miLista.InsertarUltimo(2)
	miLista.InsertarUltimo(5)

	// Filtrar elementos no repetidos
	listaSinRepetidos, dic := ListaNoRepetidos(miLista)

	// Imprimir resultado
	iter := listaSinRepetidos.Iterador()
	for iter.HaySiguiente() {
		fmt.Println(iter.VerActual())
		iter.Siguiente()
	}
	dic.Iterar(func(clave int, valor int) bool {
		fmt.Printf("Clave: %d, Valor: %d\n", clave, valor)
		return true
	})

	arreglo := []int{1, 2, 3, 2, 2, 2}
	if repetidos(arreglo) {
		fmt.Println("Hay un elemento que aparece más de la mitad de las veces.")
	} else {
		fmt.Println("Ningún elemento es mayoritario.")
	}

	pilaEj := Pila.CrearPilaDinamica[int]()
	pilaEj.Apilar(1)
	pilaEj.Apilar(2)
	pilaEj.Apilar(4)
	pilaEj.Apilar(3)
	pilaEj.Apilar(15)
	pilaEj.Apilar(6)
	pilaEj.Apilar(0)
	pilaEj.Apilar(8)
	pilaEj.Apilar(9)
	pilaEj.Apilar(20)

	cola := Distribuir(pilaEj)
	var elemCola []int
	for !cola.EstaVacia() {
		elemCola = append(elemCola, cola.Desencolar())
	}
	fmt.Println(elemCola)

	var elemPila []int
	for !pilaEj.EstaVacia() {
		elemPila = append(elemPila, pilaEj.Desapilar())
	}
	fmt.Println(elemPila)

	pila1 := Pila.CrearPilaDinamica[int]()
	pila1.Apilar(5)
	pila1.Apilar(3)
	pila1.Apilar(2)
	pila1.Apilar(1)

	// Probar si la pila es piramidal
	fmt.Println("Pila 1 es piramidal?", esPiramidal(pila1)) // Debe ser true

	// Crear otra pila que no sea piramidal
	pila2 := Pila.CrearPilaDinamica[int]()
	pila2.Apilar(5)
	pila2.Apilar(6)
	pila2.Apilar(7)

	// Probar si la pila es piramidal
	fmt.Println("Pila 2 es piramidal?", esPiramidal(pila2))

	pila3 := Pila.CrearPilaDinamica[int]()
	pila3.Apilar(5)

	// Probar si la pila es piramidal
	fmt.Println("Pila 2 es piramidal?", esPiramidal(pila3))

	numeros := []int{16, 25, 26, 36, 100, 123, 1, 0}

	for _, numero := range numeros {
		if esCuadradoPerfecto(numero) {
			fmt.Printf("%d es un cuadrado perfecto.\n", numero)
		} else {
			fmt.Printf("%d no es un cuadrado perfecto.\n", numero)
		}
	}

	numeros2 := []int{16, 25, 26, 36, 100, 123, 1, 0, 10}
	for _, numeros := range numeros2 {
		raiz := raizEntera(numeros)
		fmt.Printf("La raiz cuadrada entera de %d es %d\n", numeros, raiz)
	}

	A := []int{2, 4, 6, 8, 9, 10, 12}
	B := []int{2, 4, 6, 8, 10, 12}

	faltante := elemDistinto(A, B, 0, len(B)-1)

	fmt.Println("El elemento faltante es:", faltante)

}
