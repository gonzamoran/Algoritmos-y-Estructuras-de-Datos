package main

import (
	"fmt"
	//Heap "tdas/cola_prioridad"
	Hash "tdas/diccionario"
	Lista "tdas/lista"
)

func sonAnagramas(s1, s2 string) bool {
	letras1 := Hash.CrearHash[string, int]()
	letras2 := Hash.CrearHash[string, int]()

	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		letra1 := string(s1[i])
		apariciones1 := 1
		if letras1.Pertenece(letra1) {
			apariciones1 += letras1.Obtener(letra1)
		}
		letras1.Guardar(letra1, apariciones1)

		letra := string(s2[i])
		apariciones := 1
		if letras2.Pertenece(letra) {
			apariciones += letras2.Obtener(letra)
		}
		letras2.Guardar(letra, apariciones)
	}
	iter := letras1.Iterador()
	for iter.HaySiguiente() {
		letra, _ := iter.VerActual()
		if !(letras2.Pertenece(letra)) {
			return false
		}
		if letras1.Obtener(letra) != letras2.Obtener(letra) {
			return false
		}
		iter.Siguiente()
	}
	return true
}

func ListaNoRepetidos(lista Lista.Lista[int]) (Lista.Lista[int], Hash.Diccionario[int, int]) {
	dic := Hash.CrearHash[int, int]()
	lista.Iterar(func(elem int) bool {
		if !dic.Pertenece(elem) {
			dic.Guardar(elem, 1)
		} else {
			dic.Guardar(elem, dic.Obtener(elem)+1)
		}
		return true
	})
	resultado := Lista.CrearListaEnlazada[int]()
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if dic.Obtener(iter.VerActual()) == 1 {
			resultado.InsertarUltimo(iter.VerActual())
		}
	}
	return resultado, dic
}
func repetidos(arreglo []int) bool {
	contador := Hash.CrearHash[int, int]()
	for _, elem := range arreglo {
		veces := 1
		if contador.Pertenece(elem) {
			veces = contador.Obtener(elem) + 1
		}
		contador.Guardar(elem, veces)
		if veces > len(arreglo)/2 {
			fmt.Printf("Elemento mayoritario: %d (aparece %d veces)\n", elem, veces)
			return true
		}
	}
	return false
}

func sumarPares(lista Lista.Lista[*int]) int {
	suma := 0
	ptrSuma := &suma
	lista.Iterar(func(elem *int) bool {
		if *elem%2 == 0 {
			*ptrSuma += *elem
		}
		return true
	})
	return suma
}

type Personaje string
type Episodio string

func primeraAparicion(hashQuebrandoLoMalo, hashMejorLlamaASaul Hash.Diccionario[Personaje, []Episodio]) Hash.Diccionario[Personaje, []Episodio] {
	resultado := Hash.CrearHash[Personaje, []Episodio]()
	iter := hashMejorLlamaASaul.Iterador()
	for iter.HaySiguiente() {
		personaje, episodiosSaul := iter.VerActual()

		if hashQuebrandoLoMalo.Pertenece(personaje) {
			episodiosBB := hashQuebrandoLoMalo.Obtener(personaje)

			resultado.Guardar(personaje, []Episodio{episodiosSaul[0], episodiosBB[0]})

		}
		iter.Siguiente()
	}
	return resultado
}
