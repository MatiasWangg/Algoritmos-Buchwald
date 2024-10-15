package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const TAM_INICIAL = 5

type parClaveValor[K comparable, V any] struct {
	clave K
	valor V
}
type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// Suma los valores de los bytes y calcula el módulo del tamaño de la tabla
func hashingFuncion[K comparable](clave K, tam int) int {
	bytes := convertirABytes(clave)
	suma := 0
	for _, b := range bytes {
		suma += int(b)
	}
	return suma % tam
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashAbierto[K, V])
	hash.tabla = make([]TDALista.Lista[parClaveValor[K, V]], TAM_INICIAL)
	for i := 0; i < TAM_INICIAL; i++ {
		hash.tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	hash.tam = TAM_INICIAL
	return hash
}

func (h *hashAbierto[K, V]) redimensionar(nuevoTam int) {
	nuevaTabla := make([]TDALista.Lista[parClaveValor[K, V]], nuevoTam)
	for i := 0; i < nuevoTam; i++ {
		nuevaTabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}

	for i := 0; i < h.tam; i++ {
		lista := h.tabla[i]
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			claveValor := iter.VerActual()
			indice := hashingFuncion(claveValor.clave, nuevoTam)
			nuevaTabla[indice].InsertarUltimo(parClaveValor[K, V]{claveValor.clave, claveValor.valor})
			iter.Siguiente()
		}
	}
	h.tabla = nuevaTabla
	h.tam = nuevoTam
}

func (h *hashAbierto[K, V]) buscar(clave K) TDALista.IteradorLista[parClaveValor[K, V]] {
	indice := hashingFuncion(clave, h.tam)
	listaActual := h.tabla[indice]
	iter := listaActual.Iterador()
	for iter.HaySiguiente() {
		claveValorActual := iter.VerActual()
		if clave == claveValorActual.clave {
			return iter
		}
		iter.Siguiente()
	}
	return nil
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	factorCarga := h.cantidad / h.tam
	if factorCarga > 3 {
		h.redimensionar(h.tam * 2)
	}
	indice := hashingFuncion(clave, h.tam)
	iter := h.buscar(clave)

	if iter == nil {
		h.tabla[indice].InsertarUltimo(parClaveValor[K, V]{clave: clave, valor: dato})
		h.cantidad++
	} else {
		iter.Borrar()
		h.tabla[indice].InsertarUltimo(parClaveValor[K, V]{clave: clave, valor: dato})
	}
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	iter := h.buscar(clave)
	return iter != nil
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	iter := h.buscar(clave)

	if iter == nil {
		panic("La clave no pertenece al diccionario")
	}

	claveValor := iter.VerActual()
	return claveValor.valor
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	factorCarga := h.cantidad / h.tam
	if factorCarga < 2 && h.tam > TAM_INICIAL {
		h.redimensionar(h.tam / 2)
	}

	iter := h.buscar(clave)

	if iter == nil {
		panic("La clave no pertenece al diccionario")
	}

	claveValor := iter.VerActual()
	valor := claveValor.valor
	iter.Borrar()
	h.cantidad--
	return valor
}

func (h *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := 0; i < h.tam; i++ {
		lista := h.tabla[i]

		if lista.EstaVacia() {
			continue
		}
		iter := lista.Iterador()

		for iter.HaySiguiente() {
			claveValor := iter.VerActual()
			if !visitar(claveValor.clave, claveValor.valor) {
				return
			}
			iter.Siguiente()
		}
	}
}

// Estructura y Primitivas del iterador Externo

type iteradorDiccionario[K comparable, V any] struct {
	hash   *hashAbierto[K, V]
	indice int
	iter   TDALista.IteradorLista[parClaveValor[K, V]]
	cant   int
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	diter := new(iteradorDiccionario[K, V])
	diter.hash = h
	diter.indice = 0
	diter.cant = 0

	if h.Cantidad() == 0 {
		diter.indice = h.tam
		return diter
	}

	for diter.indice < h.tam && h.tabla[diter.indice].Largo() == 0 {
		diter.indice++
	}

	if diter.indice < h.tam {
		diter.iter = h.tabla[diter.indice].Iterador()
	}

	return diter
}

func (diter *iteradorDiccionario[K, V]) HaySiguiente() bool {
	if diter.indice >= diter.hash.tam {
		return false
	}

	if diter.iter.HaySiguiente() {
		return true
	}

	return diter.avanzarAProximaListaConElementos()
}

func (diter *iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !diter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return diter.iter.VerActual().clave, diter.iter.VerActual().valor
}

func (diter *iteradorDiccionario[K, V]) Siguiente() {
	if !diter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	diter.iter.Siguiente()
	diter.cant++
}

func (diter *iteradorDiccionario[K, V]) avanzarAProximaListaConElementos() bool {
	for diter.indice++; diter.indice < diter.hash.tam; diter.indice++ {
		if diter.hash.tabla[diter.indice].Largo() > 0 {
			diter.iter = diter.hash.tabla[diter.indice].Iterador()
			return true
		}
	}
	return false
}
