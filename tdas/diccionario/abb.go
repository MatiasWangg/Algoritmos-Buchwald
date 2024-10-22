package diccionario

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

func crearNodoAbb[K comparable, V any]() *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	return nodo
}

func crearAbb[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	var raiz *nodoAbb[K, V]
	abb.raiz = raiz
	abb.cmp = funcion_cmp
	return abb
}

func (a *abb[K, V]) buscar(clave K, nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo == nil {
		return nodo
	}
	comparacion := a.cmp(clave, nodo.clave)

	if comparacion == 0 {
		return nodo
	}else if comparacion < 0 {
		return a.buscar(clave, nodo.izquierdo)
	}else {
		return a.buscar(clave, nodo.derecho)
	}
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	return a.buscar(clave, a.raiz) != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	nodo := a.buscar(clave, a.raiz)

	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}

	return nodo.dato
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

