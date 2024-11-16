package diccionario

import (
	p "tdas/pila"
)

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

func crearNodoABB[K comparable, V any]() *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	return nodo
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.raiz = nil
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
	} else if comparacion < 0 {
		return a.buscar(clave, nodo.izquierdo, nodo)
	} else {
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
		nuevoNodo := crearNodoABB[K, V]()
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
		padreSustituto := nodo
		sustituto := nodo.derecho

		for sustituto.izquierdo != nil {
			padreSustituto = sustituto
			sustituto = sustituto.izquierdo
		}

		nodo.clave = sustituto.clave
		nodo.dato = sustituto.dato

		if padreSustituto.izquierdo == sustituto {
			padreSustituto.izquierdo = sustituto.derecho
		} else {
			padreSustituto.derecho = sustituto.derecho
		}
	}

	a.cantidad--
	return dato
}

//----------------------------------Iteradores-----------------------------------------

//interno

func (a *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	a.IterarRango(nil, nil, visitar)
}

//Interno por rangos

func (nodo *nodoAbb[K, V]) iterarRangos(desde *K, hasta *K, cmp func(K, K) int, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	cmpDesde := 0
	if desde != nil {
		cmpDesde = cmp(nodo.clave, *desde)
	}

	cmpHasta := 0
	if hasta != nil {
		cmpHasta = cmp(nodo.clave, *hasta)
	}

	if cmpDesde >= 0 {
		if !nodo.izquierdo.iterarRangos(desde, hasta, cmp, visitar) {
			return false
		}
	}
	if cmpDesde >= 0 && cmpHasta <= 0 {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}

	if cmpHasta <= 0 {
		if !nodo.derecho.iterarRangos(desde, hasta, cmp, visitar) {
			return false
		}
	}

	return true
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	a.raiz.iterarRangos(desde, hasta, a.cmp, visitar)
}

// externo---------------------------------------------------------------------------------------
type iteradorArbol[K comparable, V any] struct {
	arbol *abb[K, V]
	pila  p.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   func(K, K) int
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	abbiter := new(iteradorArbol[K, V])
	abbiter.pila = p.CrearPilaDinamica[*nodoAbb[K, V]]()
	abbiter.desde = desde
	abbiter.hasta = hasta
	abbiter.arbol = a
	abbiter.cmp = a.cmp
	nodo := a.raiz

	for nodo != nil {

		cmpDesde := 0
		if desde != nil {
			cmpDesde = a.cmp(nodo.clave, *desde)
		}
		cmpHasta := 0
		if hasta != nil {
			cmpHasta = a.cmp(nodo.clave, *hasta)
		}

		if cmpDesde < 0 {
			nodo = nodo.derecho
		} else {
			if cmpHasta <= 0 {
				abbiter.pila.Apilar(nodo)
			}
			nodo = nodo.izquierdo
		}
	}
	return abbiter
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

func (abbiter *iteradorArbol[K, V]) HaySiguiente() bool {
	if abbiter.pila.EstaVacia() {
		return false
	}
	if abbiter.hasta != nil {
		actual := abbiter.pila.VerTope()
		if abbiter.arbol.cmp(actual.clave, *abbiter.hasta) > 0 {
			return false
		}
	}
	return true
}

func (abbiter *iteradorArbol[K, V]) VerActual() (K, V) {
	if !abbiter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	actual := abbiter.pila.VerTope()
	return actual.clave, actual.dato
}

func (abbiter *iteradorArbol[K, V]) Siguiente() {
	if !abbiter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := abbiter.pila.Desapilar()
	if nodo.derecho != nil {
		abbiter.apiloIzq(nodo.derecho)
	}
}

func (abbiter *iteradorArbol[K, V]) apiloIzq(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if abbiter.desde != nil && abbiter.arbol.cmp(nodo.clave, *abbiter.desde) < 0 {
			nodo = nodo.derecho
		} else {
			abbiter.pila.Apilar(nodo)
			nodo = nodo.izquierdo
		}
	}
}