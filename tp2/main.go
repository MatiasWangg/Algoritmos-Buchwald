package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDADicc "tdas/diccionario"
	LOG "tp2/log"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	recursos := TDADicc.CrearHash[string, int]()

	visitantes := TDADicc.CrearABB[int, string](func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	for scanner.Scan() {
		comando := scanner.Text()

		resultado := procesarComando(comando, visitantes, recursos)
		imprimirResultado(resultado)
	}
}

func procesarComando(comando string, visitantes TDADicc.DiccionarioOrdenado[int, string], recursos TDADicc.Diccionario[string, int]) error {
	partes := strings.Fields(comando)

	if len(partes) == 0 {
		return fmt.Errorf("comando no reconocido")
	}

	switch partes[0] {
	case "agregar_archivo":
		if len(partes) != 2 {
			return fmt.Errorf("agregar_archivo")
		}
		archivo := partes[1]
		return LOG.AgregarArchivo(archivo, visitantes, recursos)

	case "ver_visitantes":
		if len(partes) != 3 {
			return fmt.Errorf("ver_visitantes")
		}
		desde, hasta := partes[1], partes[2]
		return LOG.VerVisitantes(desde, hasta, visitantes)
	case "ver_mas_visitados":
		if len(partes) != 2 {
			return fmt.Errorf("ver_mas_visitados")
		}
		n, err := strconv.Atoi(partes[1])
		if err != nil {
			return fmt.Errorf("ver_mas_visitados")
		}
		return LOG.VerMasVisitados(n, recursos)
	default:
		return fmt.Errorf("comando no reconocido")
	}
}

func imprimirResultado(resultado error) {
	if resultado != nil {
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", resultado)
	} else {
		fmt.Println("OK")
	}
}
