package cola_prioridad

const (
	TAM_INICIAL        = 10
	FACTOR_REDIMENSION = 2
	FACTOR_REDUCCION   = 4
)

//hijo izquierdo 2*i+1
//hijo derecho 2*i+2
//padre (i-1)/2

type heap[T any] struct {
	arreglo  []T
	cmp      func(T, T) int
	cantidad int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {

	heap := new(heap[T])
	heap.cantidad = 0
	heap.cmp = funcion_cmp
	heap.arreglo = make([]T, TAM_INICIAL)
	return heap
}
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.cmp = funcion_cmp
	heap.cantidad = len(arreglo)
	if heap.cantidad == 0 {
		heap.arreglo = make([]T, TAM_INICIAL)
	} else {
		heap.arreglo = make([]T, heap.cantidad)
		copy(heap.arreglo, arreglo)

		heapify(heap.arreglo, heap.cantidad, heap.cmp)
	}

	return heap
}

func (heap *heap[T]) redimensionar(nuevoTam int) {
	nuevoArreglo := make([]T, nuevoTam)
	copy(nuevoArreglo, heap.arreglo)
	heap.arreglo = nuevoArreglo
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.arreglo[0]
}

func (heap *heap[T]) Encolar(elem T) {
	if heap.cantidad == len(heap.arreglo) {
		heap.redimensionar(FACTOR_REDIMENSION * len(heap.arreglo))
	}
	heap.arreglo[heap.cantidad] = elem
	upheap(heap.arreglo, heap.cantidad, heap.cmp)
	heap.cantidad++
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	eliminado := heap.arreglo[0]
	heap.arreglo[0], heap.arreglo[heap.cantidad-1] = heap.arreglo[heap.cantidad-1], heap.arreglo[0]
	heap.cantidad--

	if heap.cantidad > 0 && heap.cantidad == len(heap.arreglo)/FACTOR_REDUCCION {
		heap.redimensionar(len(heap.arreglo) / FACTOR_REDIMENSION)
	}

	downheap(heap.arreglo, heap.cantidad, 0, heap.cmp)

	return eliminado
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

//----------------------------------------------------------//

func upheap[T any](arreglo []T, i int, cmp func(T, T) int) {
	for i > 0 {
		padre := padre(i)
		if cmp(arreglo[i], arreglo[padre]) <= 0 {
			return
		}
		arreglo[i], arreglo[padre] = arreglo[padre], arreglo[i]
		i = padre
	}
}
func downheap[T any](arreglo []T, n, i int, cmp func(T, T) int) {
	mayor := i
	hijoIzquierdo := hijoIzq(i)
	hijoDerecho := hijoDer(i)

	if hijoIzquierdo < n && cmp(arreglo[hijoIzquierdo], arreglo[mayor]) > 0 {
		mayor = hijoIzquierdo
	}
	if hijoDerecho < n && cmp(arreglo[hijoDerecho], arreglo[mayor]) > 0 {
		mayor = hijoDerecho
	}
	if mayor != i {
		arreglo[i], arreglo[mayor] = arreglo[mayor], arreglo[i]
		downheap(arreglo, n, mayor, cmp)
	}
}

func hijoIzq(i int) int {
	return 2*i + 1
}
func hijoDer(i int) int {
	return 2*i + 2
}
func padre(i int) int {
	return (i - 1) / 2
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	n := len(elementos)

	heapify(elementos, n, funcion_cmp)

	for i := n - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downheap(elementos, i, 0, funcion_cmp)
	}
}

func heapify[T any](elementos []T, n int, cmp func(T, T) int) {
	for i := n/2 - 1; i >= 0; i-- {
		downheap(elementos, n, i, cmp)
	}

}
