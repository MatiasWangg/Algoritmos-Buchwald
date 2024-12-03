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

const LAYOUT = "2006-01-02T15:04:05-07:00"

func AgregarArchivo(archivo string, visitantes diccionario.DiccionarioOrdenado[int, string], recursos diccionario.Diccionario[string, int]) error {
	contenido, err := os.Open(archivo)
	if err != nil {
		return fmt.Errorf("agregar_archivo")
	}
	defer contenido.Close()

	scanner := bufio.NewScanner(contenido)

	IpRequeridas := diccionario.CrearHash[string, []time.Time]()

	for scanner.Scan() {
		log := strings.Fields(scanner.Text())

		ip := log[0]
		sitio := log[3]
		t := log[1]

		mantenimiento(ip, sitio, visitantes, recursos)

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

	DoSIPs := verificarDoS(visitantes, IpRequeridas)

	for _, ip := range DoSIPs {
		fmt.Printf("DoS: %s\n", ip)
	}

	return nil
}

func verificarDoS(visitantes diccionario.DiccionarioOrdenado[int, string], IpRequeridas diccionario.Diccionario[string, []time.Time]) []string {
	DoSIPs := []string{}
	for iter := visitantes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, dato := iter.VerActual()
		if IpRequeridas.Pertenece(dato) && detectarDos(IpRequeridas.Obtener(dato)) {
			DoSIPs = append(DoSIPs, dato)
		}
	}
	return DoSIPs
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
