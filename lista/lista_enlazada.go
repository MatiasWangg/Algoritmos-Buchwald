package lista

type lista_enlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
}
type nodoLista[T any] struct {
	dato T
	sig  *nodoLista[T]
}

func (lista *lista_enlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *lista_enlazada[T]) InsertarPrimero(dato T) {
	nodo := nodoCrear(dato)

	nodo.sig = lista.primero
	lista.primero = nodo

}

func (lista *lista_enlazada[T]) InsertarUltimo(dato T) {
	nodo := nodoCrear(dato)

	lista.ultimo.sig = nodo
	lista.ultimo = nodo
}

func (lista *lista_enlazada[T]) BorrarPrimero() T {

	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	borrado := lista.primero
	lista.primero = lista.primero.sig
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

func nodoCrear[T any](dato T) *nodoLista[T] {
	nodoLista := new(nodoLista[T])

	nodoLista.dato = dato
	nodoLista.sig = nil

	return nodoLista
}
