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