package cola_prioridad

const TAM_INICIAL = 10

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

func (heap *heap[T]) redimensionar(nuevoTam int) {
	nuevoArreglo := make([]T, nuevoTam)
	copy(nuevoArreglo, heap.arreglo)
	heap.arreglo = nuevoArreglo
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.cmp = funcion_cmp
	heap.cantidad = len(arreglo)
	heap.arreglo = make([]T, heap.cantidad)

	copy(heap.arreglo, arreglo)

	for i := heap.cantidad/2 - 1; i >= 0; i-- {
		heap.heapify(i)
	}

	return heap
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(elem T) {
	if heap.cantidad == len(heap.arreglo) {
		heap.redimensionar(2 * len(heap.arreglo))
	}
	heap.arreglo[heap.cantidad] = elem
	heap.upheap(heap.cantidad)
	heap.cantidad++
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.arreglo[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	eliminado := heap.arreglo[0]
	heap.arreglo[0], heap.arreglo[heap.cantidad-1] = heap.arreglo[heap.cantidad-1], heap.arreglo[0]
	heap.cantidad--

	if heap.cantidad > 0 && heap.cantidad == len(heap.arreglo)/4 {
		heap.redimensionar(len(heap.arreglo) / 2)
	}

	heap.downheap(0)

	return eliminado
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *heap[T]) downheap(i int) {

	for i < heap.cantidad {
		hijoIzquierdo := hijoIzq(i)
		if hijoIzquierdo >= heap.cantidad {
			return
		}
		hijoDerecho := hijoDer(i)
		hijoMayor := heap.hijoMaximo(hijoIzquierdo, hijoDerecho)

		if heap.cmp(heap.arreglo[i], heap.arreglo[hijoMayor]) >= 0 {
			return
		}
		heap.arreglo[i], heap.arreglo[hijoMayor] = heap.arreglo[hijoMayor], heap.arreglo[i]
		i = hijoMayor
	}
}
func (heap *heap[T]) upheap(i int) {
	for i >= 0 {
		padre := padre(i)
		if padre < 0 || heap.cmp(heap.arreglo[i], heap.arreglo[padre]) <= 0 {
			return
		}
		heap.arreglo[i], heap.arreglo[padre] = heap.arreglo[padre], heap.arreglo[i]
		i = padre
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

func (heap *heap[T]) hijoMaximo(i, j int) int {
	if j < heap.cantidad && heap.cmp(heap.arreglo[j], heap.arreglo[i]) > 0 {
		return j
	} else {
		return i
	}
}

func (heap *heap[T]) heapify(i int) {
	hijoIzquierdo := hijoIzq(i)
	hijoDerecho := hijoDer(i)
	mayor := i

	if hijoIzquierdo < heap.cantidad && heap.cmp(heap.arreglo[hijoIzquierdo], heap.arreglo[mayor]) > 0 {
		mayor = hijoIzquierdo
	}

	if hijoDerecho < heap.cantidad && heap.cmp(heap.arreglo[hijoDerecho], heap.arreglo[mayor]) > 0 {
		mayor = hijoDerecho
	}

	if mayor != i {
		heap.arreglo[i], heap.arreglo[mayor] = heap.arreglo[mayor], heap.arreglo[i]
		heap.heapify(mayor)
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heap := CrearHeapArr(elementos, funcion_cmp)

	for i := len(elementos) - 1; i >= 0; i-- {
		elementos[i] = heap.Desencolar()
	}
}
