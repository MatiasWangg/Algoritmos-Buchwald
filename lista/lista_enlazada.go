package lista

type nodoLista[T any] struct {
	dato T
	sig  *nodoLista[T]
}

func nodoCrear[T any](dato T) *nodoLista[T] {
	nodoLista := new(nodoLista[T])
	
	nodoLista.dato = dato
	nodoLista.sig = nil
	
	return nodoLista
}

type lista_enlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo int
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(lista_enlazada[T])

	return lista
}

func (lista *lista_enlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *lista_enlazada[T]) InsertarPrimero(dato T) {
	nodo := nodoCrear(dato)

	nodo.sig = lista.primero
	lista.primero = nodo
	lista.largo++
}

func (lista *lista_enlazada[T]) InsertarUltimo(dato T) {
	nodo := nodoCrear(dato)

	lista.ultimo.sig = nodo
	lista.ultimo = nodo
	lista.largo++
}

func (lista *lista_enlazada[T]) BorrarPrimero() T {

	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	borrado := lista.primero
	lista.primero = lista.primero.sig
	lista.largo--
	return borrado.dato
}

func (lista *lista_enlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *lista_enlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *lista_enlazada[T]) Largo() int {
	return lista.largo
}
