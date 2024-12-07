package log

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/diccionario"
	"time"
	servidor "tp2/TDAServidor"
)

const LAYOUT = "2006-01-02T15:04:05-07:00"

func AgregarArchivo(archivo string, servidor *servidor.Servidor) error {
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

		servidor.Mantenimiento(ip, sitio)
		registrarTiempo(ip, t, IpRequeridas)
	}

	DoSIPs := verificarDoS(IpRequeridas)

	DoSIPsOrdenadas := radixSort(DoSIPs)

	imprimirDoS(DoSIPsOrdenadas)

	return nil
}

func verificarDoS(IpRequeridas diccionario.Diccionario[string, []time.Time]) []string {
	DoSIPs := []string{}
	for iter := IpRequeridas.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, tiempos := iter.VerActual()
		if IpRequeridas.Pertenece(clave) && detectarDos(tiempos) {
			DoSIPs = append(DoSIPs, clave)
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

func imprimirDoS(DoSIPs []string) {
	for _, ip := range DoSIPs {
		fmt.Printf("DoS: %s\n", ip)
	}
}

func registrarTiempo(ip, t string, IpRequeridas diccionario.Diccionario[string, []time.Time]) {
	registroTiempo, _ := time.Parse(LAYOUT, t)

	if IpRequeridas.Pertenece(ip) {
		nuevaLista := IpRequeridas.Obtener(ip)
		IpRequeridas.Guardar(ip, append(nuevaLista, registroTiempo))
	} else {
		IpRequeridas.Guardar(ip, []time.Time{registroTiempo})
	}
}

func radixSort(ips []string) []string {
	ordenadoPorOcteto4 := counting(ips, func(ip string) int {
		octetos := strings.Split(ip, ".")
		valor, _ := strconv.Atoi(octetos[3])
		return valor
	})

	ordenadoPorOcteto3 := counting(ordenadoPorOcteto4, func(ip string) int {
		octetos := strings.Split(ip, ".")
		valor, _ := strconv.Atoi(octetos[2])
		return valor
	})

	ordenadoPorOcteto2 := counting(ordenadoPorOcteto3, func(ip string) int {
		octetos := strings.Split(ip, ".")
		valor, _ := strconv.Atoi(octetos[1])
		return valor
	})

	ordenadoPorOcteto1 := counting(ordenadoPorOcteto2, func(ip string) int {
		octetos := strings.Split(ip, ".")
		valor, _ := strconv.Atoi(octetos[0])
		return valor
	})

	return ordenadoPorOcteto1
}

func counting(arr []string, criterio func(string) int) []string {
	freq := make([]int, 256)
	for _, ip := range arr {
		freq[criterio(ip)] += 1
	}

	inicios := make([]int, 256)
	for i := 1; i < 256; i++ {
		inicios[i] = inicios[i-1] + freq[i-1]
	}

	ordenado := make([]string, len(arr))
	for _, ip := range arr {
		indice := criterio(ip)
		ordenado[inicios[indice]] = ip
		inicios[indice] += 1
	}
	return ordenado
}
