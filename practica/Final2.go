package main

/*
Implementar una primitiva para el árbol binario EsABB(func(T, T) int) bool que reciba una función de comparación y determine
si el árbol cumple con la propiedad de ABB para dicha función de comparación. Indicar y justificar la complejidad del algoritmo
implementado.
*/

//Un arbol binario es ABB cuando es mayor a sus hijos izquierdos y menor a sus hijos derechos.
/*
				9
		4			15
	3		7	10		20

Y verificamos si es un ABB usando la función cmp para enteros, donde cmp(a, b) retorna:
	Negativo si a < b
	Cero si a == b
	Positivo si a > b
*/

type arbol[T any] struct {
	izq   *arbol[T]
	der   *arbol[T]
	clave T
}

func (a *arbol[T]) esAbb(cmp func(T, T) int) bool {

	return esAbbRec(a, nil, nil, cmp)
}

func esAbbRec[T any](nodo *arbol[T], min, max *T, cmp func(T, T) int) bool {
	// 1. Caso base
	if nodo == nil {
		return true
	}

	if min != nil && cmp(nodo.clave, *min) <= 0 {
		return false
	}

	if max != nil && cmp(nodo.clave, *max) >= 0 {
		return false
	}

	return esAbbRec(nodo.izq, min, &nodo.clave, cmp) && esAbbRec(nodo.der, &nodo.clave, max, cmp)
}
