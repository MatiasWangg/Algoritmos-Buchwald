package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"tdas/diccionario"
)

const LAYOUT = "2006-01-02T15:04:05-07:00" //Layout dado por catedra

/*
Se procesaria cada linea del .log y tambien se detectaria si hay DoS
*/

func AgregarArchivo(archivo string, visitantes diccionario.DiccionarioOrdenado[int, string], recursos diccionario.Diccionario[string, int]) error {
	contenido, err := os.Open(archivo)
	if err != nil {
		return fmt.Errorf("error al leer el archivo")
	}
	defer contenido.Close()
	scanner := bufio.NewScanner(contenido)
	//El hash IpRequeridas almacena las IPs como claves y 
	//las listas de registros de tiempo como valores.
	IpRequeridas := diccionario.CrearHash[string, []time.Time]()
	for scanner.Scan() {
		log := strings.Fields(scanner.Text())
		if recursos.Pertenece(log[3]) {
			n := recursos.Obtener(log[3])
			recursos.Guardar(log[3], n+1)
		} else {
			recursos.Guardar(log[3], 1)
		}

		ip := log[0]
		ipNumerica := conversionIP(ip)
		if !visitantes.Pertenece(ipNumerica){
			visitantes.Guardar(ipNumerica, ip)
		}
		
		registroTiempo, err := time.Parse(LAYOUT, log[1])
		if err != nil {
			return fmt.Errorf("error al .Parse la fecha: %v", err)
		}

		if IpRequeridas.Pertenece(ip) {
			nuevaLista := IpRequeridas.Obtener(ip)
			IpRequeridas.Guardar(ip, append(nuevaLista, registroTiempo))
		} else {
			IpRequeridas.Guardar(ip, []time.Time{registroTiempo})
		}
		detectarDos(IpRequeridas.Obtener(ip), ip)
	}
	return nil
}

func conversionIP(ip string) int {
	valores := strings.Split(ip, ".")
	res := 0
	for i, e := range valores {
		n, err := strconv.Atoi(e)
		if err != nil {
			return 0
		}
		res += n << (8 * (3 - i))
	}
	return res
}

func detectarDos(tiemposSolicitud []time.Time, ip string) {
    if len(tiemposSolicitud) < 10 {
        return
    }

    primerSolicitud := tiemposSolicitud[len(tiemposSolicitud)-10]  
    ultimoSolicitud := tiemposSolicitud[len(tiemposSolicitud)-1]  

    if ultimoSolicitud.Sub(primerSolicitud) <= 10*time.Second {
        fmt.Printf("DoS: %s\n", ip)
    }
}