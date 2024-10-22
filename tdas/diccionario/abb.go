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

