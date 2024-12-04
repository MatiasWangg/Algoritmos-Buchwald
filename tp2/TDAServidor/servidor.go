package servidor

import (
	"fmt"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"tp2/Ip"
)

type Servidor struct {
	recursos   diccionario.Diccionario[string, int]
	visitantes diccionario.DiccionarioOrdenado[*Ip.IP, string]
}

func CrearServidor() *Servidor {
	server := new(Servidor)
	server.recursos = diccionario.CrearHash[string, int]()
	server.visitantes = diccionario.CrearABB[*Ip.IP, string](func(a, b *Ip.IP) int {
		return Ip.CompararIps(a, b)
	})
	return server
}

func (server *Servidor) ObtenerVisitantes(desde, hasta string) []string {
	visitantes := make([]string, 0, server.visitantes.Cantidad())
	ipDeesde := Ip.CrearIp(desde)
	ipHasta := Ip.CrearIp(hasta)
	server.visitantes.IterarRango(&ipDeesde, &ipHasta, func(clave *Ip.IP, dato string) bool {
		visitantes = append(visitantes, dato)
		return true
	})
	return visitantes
}

func (server *Servidor) CalcularMasVisitados(n int) []string {
	sitios := make([]string, 0, server.recursos.Cantidad())

	server.recursos.Iterar(func(clave string, valor int) bool {
		sitios = append(sitios, clave)
		return true
	})

	heap := cola_prioridad.CrearHeapArr(sitios, func(a, b string) int {
		return server.recursos.Obtener(a) - server.recursos.Obtener(b)
	})
	resultado := make([]string, 0, n)
	for i := 0; i < n && !heap.EstaVacia(); i++ {
		resultado = append(resultado, heap.Desencolar())
	}
	return resultado
}

func (server *Servidor) Mantenimiento(ip, sitio string) {
	if server.recursos.Pertenece(sitio) {
		n := server.recursos.Obtener(sitio)
		server.recursos.Guardar(sitio, n+1)
	} else {
		server.recursos.Guardar(sitio, 1)
	}

	Ip := Ip.CrearIp(ip)
	if !server.visitantes.Pertenece(Ip) {
		server.visitantes.Guardar(Ip, ip)
	}

}

func (server *Servidor) MostrarSitios(sitios []string) {
	fmt.Println("Sitios más visitados:")
	for _, i := range sitios {
		fmt.Printf("\t%s - %d\n", i, server.recursos.Obtener(i))
	}
}

func (server *Servidor) MostrarVisitantes(visitantes []string) {
	fmt.Println("Visitantes:")
	for _, ip := range visitantes {
		fmt.Printf("\t%s\n", ip)
	}
}

//fmt.Printf("\t%s\n", ip)
