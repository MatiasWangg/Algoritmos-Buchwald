package tp2

import (
	"fmt"
	"tdas/cola_prioridad"
	"tdas/diccionario"
)

func VerVisitantes(desde, hasta string, visitantes diccionario.DiccionarioOrdenado[int, string]) {
	ipDesde := conversionIP(desde)
	ipHasta := conversionIP(hasta)

	fmt.Println("Visitantes:")
	iter := visitantes.IteradorRango(&ipDesde, &ipHasta)
	for iter.HaySiguiente() {
		_, ip := iter.VerActual()
		fmt.Printf("\t%s\n", ip)
	}
	iter.Siguiente()
}

func VerMasVisitados(n int, recursos diccionario.Diccionario[string, int]) {

	heap := cola_prioridad.CrearHeap(func(a, b string) int {
		valorA := recursos.Obtener(a)
		valorB := recursos.Obtener(b)
		if valorA > valorB {
			return -1
		} else if valorA < valorB {
			return 1
		}
		return 0
	})

	recursos.Iterar(func(k string, v int) bool {
		heap.Encolar(k)
		return true
	})

	fmt.Println("Sitios más visitados:")
	for i := 0; i < n && !heap.EstaVacia(); i++ {
		clave := heap.Desencolar()
		valor := recursos.Obtener(clave)
		fmt.Printf("\t%s - %d\n", clave, valor)
	}
}
