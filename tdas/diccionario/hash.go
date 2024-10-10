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
	tabla []TDALista.Lista[parClaveValor[K, V]]
	tam int
	cantidad int
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

//Suma los valores de los bytes y calcula el módulo del tamaño de la tabla
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
    hash.tam = TAM_INICIAL
    return hash
	}
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
        guardado := iter.VerActual()
        guardado.valor = dato
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
		panic("La clave no pertecene al diccionario")
	}

	claveValor := iter.VerActual()
	valor := claveValor.valor
	iter.Borrar()
	h.cantidad--
	return valor
}