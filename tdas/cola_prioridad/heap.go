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
	heap.cantidad = TAM_INICIAL
	heap.cmp = funcion_cmp
	heap.arreglo = make([]T, heap.cantidad)
	return heap
}

//func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T]

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(elem T) {
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
