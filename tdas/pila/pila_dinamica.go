package pila

const (
	TAMAÑO_INICIAL    = 10
	TAM_REDIMENSION   = 2
	MOMENTO_REDUCCION = 4
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, TAMAÑO_INICIAL)
	return pila
}

func (pila *pilaDinamica[T]) redimensionar(tamaño int) {
	nuevaPila := make([]T, tamaño)
	copy(nuevaPila, pila.datos)
	pila.datos = nuevaPila
}

func (pila *pilaDinamica[T]) Apilar(nuevo T) {
	pila.datos[pila.cantidad] = nuevo
	pila.cantidad++
	if pila.cantidad == len(pila.datos) {
		pila.redimensionar(pila.cantidad * TAM_REDIMENSION)
	}
}
func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	ultimo := pila.datos[pila.cantidad-1]
	pila.cantidad--
	if pila.cantidad <= len(pila.datos)/MOMENTO_REDUCCION && len(pila.datos) >= TAMAÑO_INICIAL {
		pila.redimensionar(len(pila.datos) / TAM_REDIMENSION)
	}
	return ultimo
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[t]) EstaVacia() bool {
	return pila.cantidad == 0
}
