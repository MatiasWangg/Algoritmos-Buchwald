import sys
from TDAGRAFO.grafo import Grafo
import csv
from funciones import f

def procesar_archivo(ruta):
    with open(ruta) as archivo:
        tsv = csv.reader(archivo, delimiter="\t")
        lista_archivo = list(tsv)
        cabecera = lista_archivo[0]
        datos = lista_archivo[1:] 
    usuarios_canciones, canciones_usuarios = f.cargar_usuarios_canciones(datos)
    generos_por_cancion = f.cargar_generos_por_cancion(datos)
    canciones_populares = f.construir_lista_canciones_populares(datos)
    grafo_bipartito = f.construir_grafo_bipartito(usuarios_canciones, canciones_usuarios)

    return grafo_bipartito, usuarios_canciones, canciones_usuarios, generos_por_cancion, canciones_populares

def procesar_entrada(entrada, grafo_bipartito, usuarios_canciones, canciones_usuarios, generos_por_cancion, canciones_populares):
    entrada=entrada.split(" ")
    if not entrada:
        raise Exception("Porfavor, ingrese un comando valido")
    comando=entrada[0]
    parametros=entrada[1:]

    if comando == "camino":
        if len(parametros)<3:
            raise Exception(f"Faltan argumentos para el comando: {comando}")
        canciones="".join(parametros)
        canciones.split(">>>>")
        origen=parametros[0]
        destino=parametros[2]
   
    elif comando == "mas_importantes":
        try:
            n = int(parametros.strip()[0])
        except ValueError:
            print(f"Error en el parametro de {comando}")

    elif comando =="recomendacion":
        try:
            n = parametros[1]
        except ValueError:
            print(f"Error en los parametros de {comando}")

        canciones = "".join(parametros[2:])
        canciones = canciones.split(">>>>")

        if parametros[0]=="usuarios":
            pass
        elif parametros[0]=="canciones":
            pass
        else:
            raise Exception(f"error en parametro de {comando}")
    elif comando == "ciclo":
        try:
            n = parametros[0]
        except ValueError:
            print(f"Error en los parametros de {comando}")
        cancion="".join(parametros[1:])
        pass
    elif comando == "rango":
        try:
            n = parametros[0]
        except ValueError:
            print(f"Error en los parametros de {comando}")
        cancion="".join(parametros[1:])
        pass
    else:
        raise Exception("Comando no reconocido")
        
    

def main():
    archivo = sys.argv[1]
    grafo_bipartito, usuarios_canciones, canciones_usuarios, generos_por_cancion, canciones_populares = procesar_archivo(archivo)
    grafo_canciones_repetidas = Grafo()

    for entrada in sys.stdin:
        procesar_entrada(entrada, grafo_bipartito, usuarios_canciones, canciones_usuarios, generos_por_cancion, canciones_populares)


main()