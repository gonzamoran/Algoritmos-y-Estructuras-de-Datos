package lista

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}
type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type iteradorListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

// CrearListaEnlazada crea una nueva lista enlazada vacía.
// PRECONDICIÓN: Ninguna.
// POSTCONDICIÓN: Devuelve una lista vacía, lista para usar.
func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

// crearNodo crea un nuevo nodo que contiene un dato y apunta al siguiente nodo.
// PRECONDICIÓN: Ninguna.
// POSTCONDICIÓN: Devuelve un puntero a un nodo que contiene el dato dado y siguiente igual a nil.
func crearNodo[T any](nuevo_nodo T, proximo *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{nuevo_nodo, proximo}
}

// EstaVacia indica si la lista enlazada no contiene elementos.
// PRECONDICIÓN: La lista debe haber sido creada correctamente.
// POSTCONDICIÓN: Devuelve true si la lista no tiene elementos, false en caso contrario.
func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

// InsertarPrimero inserta un nuevo elemento al principio de la lista enlazada.
// PRECONDICIÓN: La lista debe haber sido creada correctamente.
// POSTCONDICIÓN: El nuevo elemento se ubica en la primera posición de la lista, y el largo aumenta en 1.
func (lista *listaEnlazada[T]) InsertarPrimero(nodo T) {
	nuevoNodo := crearNodo(nodo, lista.primero)

	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}
	lista.primero = nuevoNodo
	lista.largo++
}

// InsertarUltimo inserta un nuevo elemento al final de la lista enlazada.
// PRECONDICIÓN: La lista debe haber sido creada correctamente.
// POSTCONDICIÓN: El nuevo elemento se ubica en la última posición de la lista, y el largo aumenta en 1.
func (lista *listaEnlazada[T]) InsertarUltimo(nodo T) {
	nuevoNodo := crearNodo(nodo, nil)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

// BorrarPrimero elimina el primer elemento de la lista enlazada y devuelve su dato.
// PRECONDICIÓN: La lista no debe estar vacía.
// POSTCONDICIÓN: El primer elemento se elimina de la lista, el largo disminuye en 1, y se devuelve su dato.
func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	aux := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--
	if lista.primero == nil {
		lista.ultimo = nil
	}
	return aux
}

// VerPrimero devuelve el dato almacenado en el primer elemento de la lista enlazada.
// PRECONDICIÓN: La lista no debe estar vacía.
// POSTCONDICIÓN: Se devuelve el dato del primer elemento sin modificar la lista.
func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

// VerUltimo devuelve el dato almacenado en el último elemento de la lista enlazada.
// PRECONDICIÓN: La lista no debe estar vacía.
// POSTCONDICIÓN: Se devuelve el dato del último elemento sin modificar la lista.
func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

// Largo devuelve la cantidad de elementos en la lista enlazada.
// PRECONDICIÓN: La lista debe haber sido creada correctamente.
// POSTCONDICIÓN: No modifica la lista.
func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

// Iterar recorre la lista enlazada desde el primero al último elemento,
// aplicando la función visitar a cada dato.
// PRECONDICIÓN: La función visitar no debe ser nil.
// POSTCONDICIÓN: Se recorre la lista en orden, aplicando visitar a cada elemento hasta que visitar devuelva false o se termine la lista.
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	if lista.EstaVacia() {
		return
	}

	actual := lista.primero

	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}
}

// Iterador devuelve un iterador externo para recorrer la lista enlazada.
// PRECONDICION: Ninguna (puede invocarse sobre cualquier lista, vacía o no).
// POSCONDICION: Devuelve un iterador posicionado en el primer elemento de la lista (o en nil si la lista está vacía).
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{
		actual:   lista.primero,
		anterior: nil,
		lista:    lista,
	}
}

// HaySiguiente indica si el iterador todavía tiene un elemento siguiente para visitar.
// PRECONDICION: Ninguna (puede llamarse en cualquier momento).
// POSCONDICIONES: Devuelve true si el iterador aún no terminó de recorrer la lista (actual != nil), false en caso contrario.
func (iterador *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

// VerActual devuelve el dato del nodo actual del iterador.
// PRECONDICION: El iterador no terminó de iterar (actual != nil).
// POSCONDICION: Devuelve el dato almacenado en el nodo actual.
func (iterador *iteradorListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

// Siguiente avanza el iterador al siguiente elemento de la lista.
// PRECONDICION: El iterador no debe haber terminado de iterar (actual != nil).
// POSCONDICION: El iterador avanza al siguiente nodo.
//
//	'anterior' pasa a ser el nodo actual previo al avance,
//	y 'actual' apunta al siguiente nodo, o nil si terminó.
func (iterador *iteradorListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

// Insertar inserta un nuevo elemento en la posición actual del iterador.
// PRECONDICION: El iterador es válido (puede estar al principio, medio o final).
// POSCONDICION: Se inserta un nuevo nodo antes del actual. El iterador queda apuntando al nuevo nodo.
func (iterador *iteradorListaEnlazada[T]) Insertar(dato T) {
	nuevoNodo := crearNodo(dato, iterador.actual)

	if !iterador.HaySiguiente() {
		iterador.lista.ultimo = nuevoNodo
	}
	if iterador.anterior == nil {
		iterador.lista.primero = nuevoNodo
	} else {
		iterador.anterior.siguiente = nuevoNodo
	}
	iterador.actual = nuevoNodo
	iterador.lista.largo++
}

// Borrar elimina el nodo actual donde está posicionado el iterador y devuelve su dato.
// PRECONDICION: El iterador no debe haber terminado de iterar (actual != nil).
// POSCONDICION: Se elimina el nodo actual. El iterador avanza al siguiente nodo automáticamente.
func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	dato := iterador.actual.dato
	iterador.actual = iterador.actual.siguiente
	if !iterador.HaySiguiente() {
		iterador.lista.ultimo = iterador.anterior
	}
	if iterador.anterior == nil {
		iterador.lista.primero = iterador.actual
	} else {
		iterador.anterior.siguiente = iterador.actual
	}

	iterador.lista.largo--
	return dato
}
