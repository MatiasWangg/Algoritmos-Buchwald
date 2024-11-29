package tp2

import (
	"fmt"
	"tdas/cola_prioridad"
	"tdas/diccionario"
)

func VerVisitantes(desde, hasta string, visitantes diccionario.DiccionarioOrdenado[int, string]) error {
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

func VerMasVisitados(n int, recursos diccionario.Diccionario[string, int]) {

	sitios := calcularMasVisitados(n, recursos)
	fmt.Println("Sitios más visitados:")
	for _, e := range sitios {
		fmt.Printf("\t%s - %d\n", e, recursos.Obtener(e))
	}
}

func calcularMasVisitados(n int, recursos diccionario.Diccionario[string, int]) []string {

	sitios := make([]string, 0, recursos.Cantidad())
	recursos.Iterar(func(k string, v int) bool {
		sitios = append(sitios, k)
		return true
	})

	heap := cola_prioridad.CrearHeapArr(sitios, func(a, b string) int {
		valorA := recursos.Obtener(a)
		valorB := recursos.Obtener(b)
		if valorA > valorB {
			return 1
		} else if valorA < valorB {
			return -1
		}
		return 0
	})
	for i := 0; i < n && !heap.EstaVacia(); i++ {
		sitio := heap.Desencolar()
		sitios[i] = sitio
	}
	return sitios
}
