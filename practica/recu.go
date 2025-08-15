package main

import (
	Pila "tdas/pila"
)

func insertarEnPos[T any](pila Pila.Pila[T], elemento T, n int) {

	pilaAux := Pila.CrearPilaDinamica[T]()
	contador := 0

	for !pila.EstaVacia() && contador < n {
		pilaAux.Apilar(pila.Desapilar())
		contador++
	}
	if contador < n {
		panic("La posicion es mayor a la cantidad de elementos de la pila")
	}

	pila.Apilar(elemento)
	for !pilaAux.EstaVacia() {
		pila.Apilar(pilaAux.Desapilar())
	}
}

func _suma_total(arreglo []float64, ini int, fin int) float64 {
	if ini > fin {
		return 0
	}
	if ini == fin {
		return arreglo[ini]
	}
	medio := (ini + fin) / 2
	suma_izq := _suma_total(arreglo, ini, medio)
	suma_der := _suma_total(arreglo, medio+1, fin)
	return float64(suma_der) + float64(suma_izq)
}

func suma_total(arreglo []float64) float64 {
	if len(arreglo) == 0 {
		return 0
	}
	return _suma_total(arreglo, 0, len(arreglo)-1)
}

func PartirPila[T any](p Pila.Pila[T], n int) Pila.Pila[T] {
	aux := Pila.CrearPilaDinamica[T]()
	contador := 0
	for !p.EstaVacia() && contador < n/2 {
		aux.Apilar(p.Desapilar())
		contador++
	}
	if contador < n/2 {
		panic("La pila tiene menos elementos que la mitad especificada")
	}
	// Crear una nueva pila para devolver la segunda mitad
	pilaSegundaMitad := Pila.CrearPilaDinamica[T]()
	// Desapilar los elementos de la pila auxiliar y apilarlos en la nueva pila
	for !aux.EstaVacia() {
		pilaSegundaMitad.Apilar(aux.Desapilar())
	}
	// Devolver la nueva pila que contiene la segunda mitad
	return pilaSegundaMitad
}

func esCuadradoPerfecto(n int) bool {
	if n <= 1 {
		return true
	}
	return _esCuadradoPerfecto(0, n, n)
}

func _esCuadradoPerfecto(inicio, fin, numero int) bool {
	if inicio > fin {
		return false
	}
	medio := (inicio + fin) / 2
	cuadrado := medio * medio

	if cuadrado == numero {
		return true
	} else if cuadrado < numero {
		return _esCuadradoPerfecto(medio+1, fin, numero)
	} else {
		return _esCuadradoPerfecto(inicio, medio-1, numero)
	}
}

func _raizEntera(n, inicio, fin int) int {
	medio := (inicio + fin) / 2
	cuadrado := medio * medio
	if cuadrado <= n && (medio+1)*(medio+1) > n {
		return medio
	}
	if cuadrado > n {
		return _raizEntera(n, inicio, medio)
	}
	return _raizEntera(n, medio+1, fin)
}

func raizEntera(n int) int {
	if n <= 0 {
		return 0
	}
	return _raizEntera(n, 0, n)
}

type Palo int

const (
	PICAS Palo = iota
	CORAZONES
	DIAMANTES
	TREBOLES
)

type Cartas struct {
	Palo   Palo
	numero int
}

/*
func buscar(arr []Cartas, elem Cartas) int {
	return _buscar(arr,elem, 0,len(arr)-1)
}

func _buscar(arr []Cartas, elem Cartas, inicio int, fin int) int {
	if inicio > fin {
		return -1
	}

	medio := (inicio + fin)/2
	cartaMedio := arr[medio]

	if elem.Palo == cartaMedio.Palo && elem.numero == cartaMedio.numero{
		return medio
	}

	if
}*/

/*
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

	insertarEnPos(pila, 70, 0)
	fmt.Println("Pila despues de insertar 70 en la posicion 2:")
	for !pila.EstaVacia() {
		fmt.Println(pila.Desapilar())
	}

	// Ejemplo de arreglo de números flotantes
	arreglo := []float64{1.5, 2.3, 3.7, 4.1, 5.6}

	// Llamamos a la función suma_total para obtener la suma
	resultado := suma_total(arreglo)

	// Imprimimos el resultado
	fmt.Printf("La suma total del arreglo es: %.2f\n", resultado)

	p := Pila.CrearPilaDinamica[int]()

	// Cargar la pila con 6 elementos (del 1 al 6)
	for i := 1; i <= 6; i++ {
		p.Apilar(i)
	}
	fmt.Println("Pila original (de arriba hacia abajo):")
	imprimirPila(p)

	// Partimos la pila
	segundaMitad := PartirPila(p, 6)

	fmt.Println("\nPrimera mitad (pila original):")
	imprimirPila(p)

	fmt.Println("\nSegunda mitad (pila devuelta):")
	imprimirPila(segundaMitad)
}

func imprimirPila[T any](p Pila.Pila[T]) {
	aux := Pila.CrearPilaDinamica[T]()
	// Invertimos para imprimir de arriba hacia abajo
	for !p.EstaVacia() {
		elem := p.Desapilar()
		fmt.Println(elem)
		aux.Apilar(elem)
	}
	// Restauramos la pila original
	for !aux.EstaVacia() {
		p.Apilar(aux.Desapilar())
	}
}
*/
