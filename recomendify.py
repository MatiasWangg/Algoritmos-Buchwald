#!/usr/bin/python3
import sys
from TDAGRAFO.grafo import Grafo
import csv
import funciones as f

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


def procesar_entrada(entrada, grafo_bipartito, grafo_canciones_repetidas, usuarios_canciones, canciones_usuarios, generos_por_cancion, canciones_populares):
    """
    Procesa la entrada del usuario y ejecuta el comando correspondiente.
    """
    entrada = entrada.strip().split(" ")
    if not entrada:
        raise ValueError("Por favor, ingrese un comando válido.")
    
    comando, parametros = entrada[0], entrada[1:]
    comandos = {
        "camino": ejecutar_camino,
        "mas_importantes": ejecutar_mas_importantes,
        "recomendacion": ejecutar_recomendacion,
        "ciclo": ejecutar_ciclo,
        "rango": ejecutar_rango
    }
    
    if comando not in comandos:
        raise ValueError(f"Comando no reconocido: {comando}")
    
    # Ejecutar el comando correspondiente
    comandos[comando](parametros, grafo_bipartito, grafo_canciones_repetidas, usuarios_canciones, canciones_usuarios, generos_por_cancion, canciones_populares)


    
    def validar_parametros(parametros,minimo):
        if len(parametros)<minimo:
            raise Exception(f"Faltan parametros para el comando: {comando}")
        
    
    def ejecutar_camino(parametros, grafo_bipartito, *_):
        canciones = "".join(parametros).split(">>>>")
        validar_parametros(canciones, 2)
        origen, destino = canciones[0].strip(), canciones[1].strip()

        f.comando_camino(grafo_bipartito, origen, destino)

    
    def ejecutar_mas_importantes(parametros,grafo_bipartito,*_):
        validar_parametros(parametros, 1)
        try:
            n = int(parametros[0])
        except ValueError:
            raise ValueError("El parámetro debe ser un número válido.")
    
        f.comando_mas_importantes(grafo_bipartito, n)

    
    def ejecutar_recomendacion(parametros,grafo_bipartito,*_):
        validar_parametros(parametros,3)
        recomendado=parametros[0]
        try:
            n=int(parametros[0])
        except ValueError:
            raise ValueError("El parámetro debe ser un número válido.")
        canciones="".join(parametros[2:]).split(">>>>")
        f.comando_recomendacion(grafo_bipartito,recomendado,n,canciones)

    
    def ejecutar_ciclo(parametros, grafo_bipartito, grafo_canciones_repetidas, usuarios_canciones, canciones_usuarios, generos_por_cancion, canciones_populares):
        validar_parametros(parametros,2)
        try:
            n=int(parametros[0])
        except ValueError:
            raise ValueError("El parámetro debe ser un número válido.")
        cancion="".join(parametros[2:])
        f.completar_grafo_canciones_repetidas(usuarios_canciones, grafo_canciones_repetidas)
        f.comando_ciclo(grafo_canciones_repetidas, int(n), cancion)

    
    def ejecutar_rango(parametros, _, grafo_canciones_repetidas, usuarios_canciones):
        validar_parametros(parametros,2)
        try:
            n=int(parametros[0])
        except ValueError:
            raise ValueError("El parámetro debe ser un número válido.")
        cancion="".join(parametros[2:])
        f.completar_grafo_canciones_repetidas(usuarios_canciones, grafo_canciones_repetidas)
        f.comando_rango(grafo_canciones_repetidas, n, cancion)
    

def main():
    archivo = sys.argv[1]
    grafo_bipartito, usuarios_canciones, canciones_usuarios, generos_por_cancion, canciones_populares = procesar_archivo(archivo)
    grafo_canciones_repetidas = Grafo()

    for entrada in sys.stdin:
        procesar_entrada(entrada, grafo_bipartito, grafo_canciones_repetidas, usuarios_canciones, canciones_usuarios, generos_por_cancion, canciones_populares)


main()

