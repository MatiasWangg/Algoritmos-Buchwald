package tp2

import(
	"bufio"
	"os"
	"fmt"
)

//Estructura inicial, nada definitivo 
func main(){
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		comando := scanner.Text()


		err := procesarComando(......)
		if err != nil{
			fmt.Println(os.Stderr,"Error en el comando ingresado")
		}
	}
}



//Pensaba hacer  una funcion con un switch para procesar cada comando recibido e ir
//llamando a las diferentes funciones que estaran en otros archivos