package cola_prioridad_test

import (
	"fmt"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}
var Ints = []int{20, 100, 40, 1, 5, 30, 60, 15, 19, 10, 5, 132, 5, 42, 26, 2, 626, 8, 626, 2}
var Strings = []string{"E", "O", "A", "I", "S", "A", "E", "T", "Y", "D", "G", "OK", "HOLA"}

func TestCrearColaVacia(t *testing.T) {
	t.Log("Crea una cola nueva y debe funcionar")
	heap := TDAHeap.CrearHeap(CompararStrings)
	require.True(t, heap.EstaVacia(), "La cola deberia estar vacia")
	require.EqualValues(t, 0, heap.Cantidad(), "La cola no deberia tener elementos")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() }, "Desencolar deberia causar panico en una cola vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() }, "VerMax deberia causar panico en una cola vacia")
	heap.Encolar("A")
	require.False(t, heap.EstaVacia(), "La cola no deberia estar vacia")
}

func TestCrearColaConArray(t *testing.T) {
	t.Log("Crea una cola a partir de un arreglo y debe actuar como una cola con elementos")
	datos := make([]int, 10)
	for e := range 10 {
		datos[e] = e
	}
	heap := TDAHeap.CrearHeapArr(datos, CompararInts)
	require.False(t, heap.EstaVacia(), "La cola no deberia estar vacia")
	require.EqualValues(t, 10, heap.Cantidad(), "La cola deberia tener 10 elementos")
	require.EqualValues(t, 9, heap.VerMax(), "El valor maximo debe ser 9")
	require.EqualValues(t, 9, heap.Desencolar(), "Se deberia desencolar el maximo")
	require.EqualValues(t, 8, heap.VerMax(), "El nuevo maximo debe ser 8")
}

func TestMismoElemento(t *testing.T) {
	t.Log("Aqui se pureba si una cola puede tener elementos repetidos")
	heap := TDAHeap.CrearHeap(CompararInts)
	i := 0
	for i != 5 {
		heap.Encolar(1)
		i++
	}
	for i != 0 {
		require.EqualValues(t, 1, heap.VerMax())
		require.EqualValues(t, 1, heap.Desencolar())
		i--
	}
}

func TestEncolarSinOrden(t *testing.T) {
	t.Log("Si encolo elementos sin ningun orden la cola deberia estar ordenada por prioridad")
	heap := TDAHeap.CrearHeap(CompararInts)
	heap.Encolar(20)
	heap.Encolar(100)
	heap.Encolar(40)
	heap.Encolar(5)
	heap.Encolar(1)
	require.EqualValues(t, 100, heap.VerMax())
	require.EqualValues(t, 100, heap.Desencolar())
	require.EqualValues(t, 40, heap.VerMax())
	require.EqualValues(t, 40, heap.Desencolar())
	require.EqualValues(t, 20, heap.VerMax())
	require.EqualValues(t, 20, heap.Desencolar())
	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 1, heap.VerMax())
	require.EqualValues(t, 1, heap.Desencolar())

}

func TestHeapSortInt(t *testing.T) {
	t.Log("Aqui probamos la funcion de heapsort con ints, creamos un arreglo desordenado, lo pasamos a la funcion y comprobamos si esta ordenado")
	datos := Ints
	TDAHeap.HeapSort(datos, CompararInts)
	for i := 1; i < len(datos); i++ {
		require.True(t, datos[i] >= datos[i-1])
	}
}

func TestColaGenerica(t *testing.T) {
	t.Log("Aqui probamos si la cola puede ser de distintos tipos de datos")
	heapInts := TDAHeap.CrearHeapArr(Ints, CompararInts)
	datosI := Ints
	TDAHeap.HeapSort(datosI, CompararInts)
	for i := len(datosI); i > 0; i-- {
		require.EqualValues(t, datosI[i-1], heapInts.Desencolar())
	}
	heapStr := TDAHeap.CrearHeapArr(Strings, CompararStrings)
	datosS := Strings
	TDAHeap.HeapSort(datosS, CompararStrings)
	for i := len(datosS); i > 0; i-- {
		require.EqualValues(t, datosS[i-1], heapStr.Desencolar())
	}

}

func ejecutarPruebaVolumen(t *testing.T, n int) {
	heap := TDAHeap.CrearHeap(CompararInts)
	datos := make([]int, n)

	ok := true
	for i := 0; i < n; i++ {
		datos[i] = i
		heap.Encolar(i)
		require.EqualValues(t, i, heap.VerMax(), "No esta encolando correctamente")
	}
	for i := n; i > 0; i-- {
		require.EqualValues(t, i-1, heap.Desencolar(), "No esta desencolando correctamente")
	}
	require.True(t, ok, "La cola no funciona bien con muchos elementos")
}
func TestVolumen(t *testing.T) {
	t.Log("Probamos como se comporta la cola con una gran cantidad de elementos, los encolamos y desencolamos y vemos si siguen el orden de prioridad")
	for _, n := range TAMS_VOLUMEN {
		t.Run(fmt.Sprintf("Prueba %d elementos", n), func(t *testing.T) {
			for i := 0; i < 6; i++ {
				ejecutarPruebaVolumen(t, n)
			}
		})
	}
}

func TestColaVaciada(t *testing.T) {
	t.Log("Aqui probamos si una cola con datos se vacia correctamente y funciona igual que una cola vacia recien creada")
	datos := []string{"E", "D", "C", "B", "A"}
	heap := TDAHeap.CrearHeapArr(datos, CompararStrings)
	for i := range len(datos) {
		require.EqualValues(t, datos[i], heap.VerMax())
		require.EqualValues(t, datos[i], heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() }, "Desencolar deberia causar panico en una cola vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() }, "VerMax deberia causar panico en una cola vacia")
}

func CompararStrings(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func CompararInts(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
