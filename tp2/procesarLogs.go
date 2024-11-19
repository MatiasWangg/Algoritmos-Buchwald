package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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