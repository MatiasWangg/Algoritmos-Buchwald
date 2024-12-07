package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	servidor "tp2/TDAServidor"
	LOG "tp2/log"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	server := servidor.CrearServidor()

	for scanner.Scan() {
		comando := scanner.Text()

		resultado := procesarComando(comando, server)
		imprimirResultado(resultado)
	}
}

func procesarComando(comando string, server *servidor.Servidor) error {
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
		return LOG.AgregarArchivo(archivo, server)

	case "ver_visitantes":
		if len(partes) != 3 {
			return fmt.Errorf("ver_visitantes")
		}
		desde, hasta := partes[1], partes[2]
		return LOG.VerVisitantes(desde, hasta, server)
	case "ver_mas_visitados":
		if len(partes) != 2 {
			return fmt.Errorf("ver_mas_visitados")
		}
		n, err := strconv.Atoi(partes[1])
		if err != nil {
			return fmt.Errorf("ver_mas_visitados")
		}
		return LOG.VerMasVisitados(n, server)
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
