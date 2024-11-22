package log

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/diccionario"
	"time"
)

const LAYOUT = "2006-01-02T15:04:05-07:00" //Layout dado por catedra
const RUTA = "pruebas_analog/"

/*
Se procesaria cada linea del .log y tambien se detectaria si hay DoS
*/

func AgregarArchivo(archivo string, visitantes diccionario.DiccionarioOrdenado[int, string], recursos diccionario.Diccionario[string, int]) error {
	ruta := fmt.Sprintf("%s%s", RUTA, archivo)
	contenido, err := os.Open(ruta)
	if err != nil {
		return fmt.Errorf("error al leer el archivo")
	}
	defer contenido.Close()

	scanner := bufio.NewScanner(contenido)

	//ipRequeridas tiene como dato las ip y como valor una lista de sus tiempos
	IpRequeridas := diccionario.CrearHash[string, []time.Time]()

	for scanner.Scan() {
		log := strings.Fields(scanner.Text())

		ip := log[0]
		sitio := log[3]
		t := log[1]

		//mantenimiento de visitantes y recursos
		mantenimiento(ip, sitio, visitantes, recursos)

		//Guardar los tiempos para cada ip
		registroTiempo, err := time.Parse(LAYOUT, t)
		if err != nil {
			return fmt.Errorf("error al .Parse la fecha: %v", err)
		}

		if IpRequeridas.Pertenece(ip) {
			nuevaLista := IpRequeridas.Obtener(ip)
			IpRequeridas.Guardar(ip, append(nuevaLista, registroTiempo))
		} else {
			IpRequeridas.Guardar(ip, []time.Time{registroTiempo})
		}

	}
	//Deteccion e Impresion de las DoS
	for iter := visitantes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, dato := iter.VerActual()
		if IpRequeridas.Pertenece(dato) && detectarDos(IpRequeridas.Obtener(dato)) {
			fmt.Printf("DoS: %s\n", dato)

		}
	}
	return nil
}

func conversionIP(ip string) int {
	valores := strings.Split(ip, ".")
	res := 0
	for i, e := range valores {
		n, err := strconv.Atoi(e)
		if err != nil {
			return -1
		}
		res += n << (8 * (3 - i))
	}
	return res
}

func detectarDos(tiemposSolicitud []time.Time) bool {
	// Verificar si ya se detectó DoS para esta IP
	cant := len(tiemposSolicitud)
	if cant < 5 {
		return false
	}
	inicio := 0
	for fin := 4; fin < cant; fin++ {
		if tiemposSolicitud[fin].Sub(tiemposSolicitud[inicio]) < 2*time.Second {
			return true
		}
		inicio++
	}
	return false
}

func mantenimiento(ip, sitio string, visitantes diccionario.DiccionarioOrdenado[int, string], recursos diccionario.Diccionario[string, int]) {
	if recursos.Pertenece(sitio) {
		n := recursos.Obtener(sitio)
		recursos.Guardar(sitio, n+1)
	} else {
		recursos.Guardar(sitio, 1)
	}

	ipNumerica := conversionIP(ip)
	if ipNumerica == -1 {
		return
	}
	if !visitantes.Pertenece(ipNumerica) {
		visitantes.Guardar(ipNumerica, ip)
	}
}
