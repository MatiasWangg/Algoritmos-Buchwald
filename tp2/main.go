package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/diccionario"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//Estructuras para guardar los datos
	visitantes := diccionario.CrearABB[int, string](compararInts)
	//visitantes debe ser un arbol para iterarlo con desde y hasta
	//la clave es un int que sera la representacion numerica
	//y el valor deberia guardar la ip
	recursos := diccionario.CrearHash[string, int]()

	for scanner.Scan() {
		comando := scanner.Text()

		err := procesarComando(comando, visitantes, recursos)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error en el comando ingresado")
		}
	}
}

// Pensaba hacer  una funcion con un switch para procesar cada comando recibido e ir
// llamando a las diferentes funciones que estaran en otros archivos
func procesarComando(comando string, visitantes diccionario.DiccionarioOrdenado[int, string], recursos diccionario.Diccionario[string, int]) error {
	partes := strings.Fields(comando)

	switch partes[0] {
	case "agregar_archivo":
		archivo := partes[1]
		//funcion que trabaja con los log (parametros provisorios)
		return AgregarArchivo(archivo, visitantes, recursos)
	case "ver_visitantes":
		desde, hasta := partes[1], partes[2]
		//Lo mismo aca con los parametros
		VerVisitantes(desde, hasta, visitantes)
	case "ver_mas_visitados":
		n, err := strconv.Atoi(partes[1])
		//Lo mismo aca con los parametros
		if err != nil {
			return fmt.Errorf("error en ver_mas_visitados")
		}
		VerMasVisitados(n, recursos)
	default:
		return fmt.Errorf("comando no reconocido")
	}
	fmt.Println("OK")
	return nil
}

func compararInts(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
