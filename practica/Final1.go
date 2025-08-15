package main

import Heap "tdas/cola_prioridad"

//Final del 21-07-2025

// Ej 1)
//Implementar un algoritmo que dado un arreglo de dígitos (0-9) determine cuál es el número más grande que se puede formar con dichos dígitos. Indicar y justificar la complejidad del algoritmo implementado.

//	func numMasGrande(digitos []int) int {
//		frecuencias := make([]int, 10)
//		for _, elem := range digitos {
//			frecuencias[elem]++
//		}
//		inicios := make([]int, len(digitos))
//		for i := 1; i < len(digitos)-1; i++ {
//			inicios[i] = inicios[i-1] + frecuencias[i-1]
//		}
//
//		var resultado string
//		for i := 9; i >= 0; i-- {
//			cantidad := frecuencias[i]
//			for j := 0; j < cantidad; j++ {
//				resultado.WriteString(strconv.Itoa(i))
//			}
//		}
//	}
func comparar(a, b int) int {
	return b - a
}

func topK(numeros []int, k int) []int {

	resultado := make([]int, k)
	heap := Heap.CrearHeapArr(numeros, comparar)

	for i := 0; i <= k; i++ {
		resultado[i] = heap.Desencolar()
	}

	return resultado
}
