package tp2

import (
	"fmt"
	"tdas/diccionario"
)

func VerVisitantes(desde, hasta string, visitantes diccionario.DiccionarioOrdenado[int, string]) {
	ipDesde := conversionIP(desde)
	ipHasta := conversionIP(hasta)

	fmt.Println("Visitantes:")
	iter := visitantes.Iterador()
	for iter.HaySiguiente(){
		ipNumerica, ip := iter.VerActual()
		if ipNumerica >= ipDesde && ipNumerica <= ipHasta {
			fmt.Printf("\t%s\n",ip)
		}
		iter.Siguiente()
	}
	fmt.Println("OK")
}

func VerMasVisitados(n int, recursos diccionario.Diccionario[string, int]) {
Ver
}