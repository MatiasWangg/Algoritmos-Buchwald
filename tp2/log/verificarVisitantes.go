package log

import (
	"fmt"
	servidor "tp2/TDAServidor"
)

func VerVisitantes(desde, hasta string, servidor *servidor.Servidor) error {
	visitantes := servidor.ObtenerVisitantes(desde, hasta)
	fmt.Println("Visitantes:")
	for _, ip := range visitantes {
		fmt.Printf("\t%s\n", ip)
	}
	return nil
}

func VerMasVisitados(n int, servidor *servidor.Servidor) error {
	sitios := servidor.CalcularMasVisitados(n)
	fmt.Println("Sitios más visitados:")
	for _, i := range sitios {
		fmt.Printf("\t%s - %d\n", i, servidor.ObtenerCantidadVisitas(i))
	}
	return nil
}
