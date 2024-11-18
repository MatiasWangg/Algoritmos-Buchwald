package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/diccionario"
)

const LAYOUT = "2015-05-17T10:05:00+00:00"

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
