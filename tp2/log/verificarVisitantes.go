package log

import (
	"fmt"
	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
)

func VerVisitantes(desde, hasta string, visitantes TDADicc.DiccionarioOrdenado[int, string]) error {
	ipDesde := conversionIP(desde)
	ipHasta := conversionIP(hasta)
	if ipDesde == -1 || ipHasta == -1 {
		return fmt.Errorf("ip no valida")
	}
	fmt.Println("Visitantes:")
	iter := visitantes.IteradorRango(&ipDesde, &ipHasta)
	for iter.HaySiguiente() {
		_, ip := iter.VerActual()
		fmt.Printf("\t%s\n", ip)
		iter.Siguiente()
	}
	return nil
}

func VerMasVisitados(n int, recursos TDADicc.Diccionario[string, int]) error {
	sitios := calcularMasVisitados(n, recursos)
	fmt.Println("Sitios más visitados:")
	for _, i := range sitios {
		fmt.Printf("\t%s - %d\n", i, recursos.Obtener(i))
	}
	return nil
}

func calcularMasVisitados(n int, recursos TDADicc.Diccionario[string, int]) []string {
	heap := TDAHeap.CrearHeap(func(a, b string) int {
		valorA := recursos.Obtener(a)
		valorB := recursos.Obtener(b)
		return valorB - valorA
	})

	recursos.Iterar(func(clave string, valor int) bool {
		if heap.Cantidad() < n {
			heap.Encolar(clave)
		} else if valor > recursos.Obtener(heap.VerMax()) {
			heap.Desencolar()
			heap.Encolar(clave)
		}
		return true
	})

	resultado := make([]string, 0, heap.Cantidad())
	for !heap.EstaVacia() {
		resultado = append(resultado, heap.Desencolar())
	}

	for i, j := 0, len(resultado)-1; i < j; i, j = i+1, j-1 {
		resultado[i], resultado[j] = resultado[j], resultado[i]
	}

	return resultado
}
