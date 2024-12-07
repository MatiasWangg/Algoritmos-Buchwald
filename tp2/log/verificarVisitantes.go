package log

import (
	servidor "tp2/TDAServidor"
)

func VerVisitantes(desde, hasta string, servidor *servidor.Servidor) error {
	visitantes := servidor.ObtenerVisitantes(desde, hasta)
	servidor.MostrarVisitantes(visitantes)
	return nil
}

func VerMasVisitados(n int, servidor *servidor.Servidor) error {
	sitios := servidor.CalcularMasVisitados(n)
	servidor.MostrarSitios(sitios)
	return nil
}
