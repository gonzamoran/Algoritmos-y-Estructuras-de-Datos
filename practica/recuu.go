package main

import (
	Cola "tdas/cola"
	Pila "tdas/pila"
)

func Distribuir[T any](pilaOriginal Pila.Pila[T]) Cola.Cola[T] {
	aux := Pila.CrearPilaDinamica[T]()
	resultado := Cola.CrearColaEnlazada[T]()

	pos := 1
	for !pilaOriginal.EstaVacia() {
		if pos%2 == 1 {
			resultado.Encolar(pilaOriginal.Desapilar())
		} else {
			aux.Apilar(pilaOriginal.Desapilar())
		}
		pos++
	}
	for !aux.EstaVacia() {
		pilaOriginal.Apilar(aux.Desapilar())
	}

	return resultado
}

func esPiramidal(pila Pila.Pila[int]) bool {
	pilaAux := Pila.CrearPilaDinamica[int]()
	if pila.EstaVacia() {
		return true
	}
	valorAnterior := pila.Desapilar()
	pilaAux.Apilar(valorAnterior)

	for !pila.EstaVacia() {
		valorActual := pila.Desapilar()
		if valorActual >= valorAnterior {
			return false
		}
		valorAnterior = valorActual
		pilaAux.Apilar(valorActual)
	}
	return true
}

func Ordenar(pila Pila.Pila[int]) {
	pilaAux := Pila.CrearPilaDinamica[int]()
	for !pila.EstaVacia() {
		elem := pila.Desapilar()
		for !pilaAux.EstaVacia() && elem > pilaAux.VerTope() {
			pila.Apilar(pilaAux.Desapilar())
		}
		pilaAux.Apilar(elem)
	}
	for !pilaAux.EstaVacia() {
		pila.Apilar(pilaAux.Desapilar())
	}
}
