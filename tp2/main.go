package tp2

import(
	"bufio"
	"os"
	"fmt"
	"strings"
)

//Estructura inicial, nada definitivo 
func main(){
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		comando := scanner.Text()

		err := procesarComando(comando, ....)
		if err != nil{
			fmt.Fprintf(os.Stderr,"Error en el comando ingresado")
		}
	}
}

//Pensaba hacer  una funcion con un switch para procesar cada comando recibido e ir
//llamando a las diferentes funciones que estaran en otros archivos
func procesarComando(comando string, ....){
	partes := strings.Fields(comando)

	switch partes[0] {
	case "agregar_archivo":
		archivo := partes[1]
		//funcion que trabaja con los log (parametros provisorios)
		return AgregarArchivo(comando, .......)
	case "ver_visitantes":
		desde, hasta := partes[1], partes[2]
		//Lo mismo aca con los parametros
		VerVisitantes(desde, hasta, ........)
	case "ver_mas_visitados":
		n, err := strconv.Atoi(partes[1])
		//Lo mismo aca con los parametros
		VerMasVisitados(.....)
	default:
		return fmt.Errorf("Comando no reconocido")
	}
	fmt.Println("OK")
	return nil
}