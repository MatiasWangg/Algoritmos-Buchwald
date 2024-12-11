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
    pass

def main():
    archivo = sys.argv[1]
    grafo_bipartito, usuarios_canciones, canciones_usuarios, generos_por_cancion, canciones_populares = procesar_archivo(archivo)
    grafo_canciones_repetidas = Grafo()

    entrada = sys.stdin.readline()
    while entrada != "":
        procesar_entrada(entrada, grafo_bipartito, usuarios_canciones, canciones_usuarios, generos_por_cancion, canciones_populares)
        entrada = sys.stdin.readline()


main()