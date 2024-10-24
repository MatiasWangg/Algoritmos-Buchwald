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

func (a *abb[K, V]) buscar(clave K, nodo *nodoAbb[K, V], padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return nodo, padre
	}
	comparacion := a.cmp(clave, nodo.clave)

	if comparacion == 0 {
		return nodo, padre
	}else if comparacion < 0 {
		return a.buscar(clave, nodo.izquierdo, nodo)
	}else {
		return a.buscar(clave, nodo.derecho, nodo)
	}
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := a.buscar(clave, a.raiz, nil)
	return nodo != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	nodo, _ := a.buscar(clave, a.raiz, nil)

	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}

	return nodo.dato
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	nodo, padre := a.buscar(clave, a.raiz, nil)

	if nodo != nil {
		nodo.dato = dato
	} else {
		nuevoNodo := crearNodoAbb[K, V]() 
		nuevoNodo.clave = clave
		nuevoNodo.dato = dato

		if padre != nil { 
			comparacion := a.cmp(clave, padre.clave)
			if comparacion < 0 {
				padre.izquierdo = nuevoNodo
			} else {
				padre.derecho = nuevoNodo
			}
		} else {
			a.raiz = nuevoNodo 
		}

		a.cantidad++
	}
}

func (a *abb[K, V]) Borrar(clave K) V {
	nodo, padre := a.buscar(clave, a.raiz, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}

	dato := nodo.dato 

	if nodo.izquierdo == nil && nodo.derecho == nil {
		if padre != nil {
			if a.cmp(clave, padre.clave) < 0 {
				padre.izquierdo = nil
			} else {
				padre.derecho = nil
			}
		} else {
			a.raiz = nil
		}
	} else if nodo.izquierdo == nil || nodo.derecho == nil {
		var hijo *nodoAbb[K, V]
		if nodo.izquierdo != nil {
			hijo = nodo.izquierdo
		} else {
			hijo = nodo.derecho
		}

		if padre != nil {
			if a.cmp(clave, padre.clave) < 0 {
				padre.izquierdo = hijo
			} else {
				padre.derecho = hijo
			}
		} else {
			a.raiz = hijo
		}
	} else {
		sustituto := a.encontrarMinimo(nodo.derecho)
		nodo.clave = sustituto.clave
		nodo.dato = sustituto.dato
		a.Borrar(sustituto.clave) 
	}

	a.cantidad--
	return dato
}

func (a *abb[K, V]) encontrarMinimo(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	for nodo.izquierdo != nil {
		nodo = nodo.izquierdo
	}
	return nodo
}
