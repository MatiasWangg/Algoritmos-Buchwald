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
	sitios := make([]string, 0, recursos.Cantidad())
	resultado := make([]string, 0, n)
	recursos.Iterar(func(clave string, valor int) bool {
		sitios = append(sitios, clave)
		return true
	})
	heap := TDAHeap.CrearHeapArr(sitios, func(a, b string) int {
		valorA := recursos.Obtener(a)
		valorB := recursos.Obtener(b)
		return valorB - valorA
	})

	for i := 0; i < n && !heap.EstaVacia(); i++ {
		resultado = append(resultado, heap.Desencolar())
	}
	return resultado
}
