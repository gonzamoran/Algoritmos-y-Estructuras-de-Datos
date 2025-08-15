package cola

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato      T
	siguiente *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(colaEnlazada[T])
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) nuevoNodo(nuevo_nodo T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = nuevo_nodo
	return nodo
}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.
func (cola *colaEnlazada[T]) Encolar(nuevo_nodo T) {
	nodo := cola.nuevoNodo(nuevo_nodo)

	if !cola.EstaVacia() {
		cola.ultimo.siguiente = nodo
	} else {
		cola.primero = nodo
	}

	cola.ultimo = nodo
}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	valor := cola.VerPrimero()
	cola.primero = cola.primero.siguiente
	return valor
}
