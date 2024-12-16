#!/usr/bin/python3
import sys,csv
from grafo import Grafo
import comandos as f
import utils as u


def procesar_archivo(ruta):
    with open(ruta) as archivo:
        tsv = csv.reader(archivo, delimiter="\t")
        lista_archivo = list(tsv)
        datos = lista_archivo[1:] 

    usuarios_canciones, canciones_usuarios,usuarios_playlist,playlists_nombres,playlists_canciones = u.guardar_datos(datos)
    grafo_bipartito = u.construir_grafo_bipartito(usuarios_canciones, canciones_usuarios)
    grafo_canciones_repetidas=f.completar_grafo_canciones_repetidas(usuarios_canciones)

    return grafo_bipartito,grafo_canciones_repetidas, usuarios_canciones, canciones_usuarios, usuarios_playlist,playlists_nombres,playlists_canciones


def procesar_entrada(entrada,grafo_bipartito,grafo_canciones_repetidas, usuarios_canciones, canciones_usuarios, usuarios_playlist,playlists_nombres,playlists_canciones):
    """
    Procesa la entrada del usuario y ejecuta el comando correspondiente.
    """
    entrada = entrada.strip().split(maxsplit=1)
    if not entrada:
        raise ValueError("Por favor, ingrese un comando válido.")
    
    comando, parametros = entrada[0], entrada[1]
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
    comandos[comando](parametros,grafo_bipartito,grafo_canciones_repetidas, usuarios_canciones, canciones_usuarios, usuarios_playlist,playlists_nombres,playlists_canciones)


    
def validar_parametros(parametros,minimo):
    if len(parametros)<minimo:
        raise Exception("Faltan parametros")
        
    
def ejecutar_camino(parametros,grafo_bipartito,grafo_canciones_repetidas, usuarios_canciones, canciones_usuarios, usuarios_playlist,playlists_nombres,playlists_canciones):
    parametros= parametros.split(">>>>")
    validar_parametros(parametros, 2)
    origen, destino = parametros[0].strip(), parametros[1].strip()
    if origen not in canciones_usuarios or destino not in canciones_usuarios:
        print("Tanto el origen como el destino deben ser canciones")
        return
    camino=f.comando_camino(grafo_bipartito, origen, destino)
    if not camino:
        return
    salida=u.formatear_camino(camino,usuarios_playlist,playlists_nombres,playlists_canciones)
    imprimir(salida)
    
def ejecutar_mas_importantes(parametros, grafo_bipartito,*_):
    validar_parametros(parametros, 1)
    try:
        n = int(parametros)
    except ValueError:
        raise ValueError("El parámetro debe ser un número válido.")
    salida=f.comando_mas_importantes(grafo_bipartito, n)
    imprimir(salida)


def ejecutar_recomendacion(parametros,grafo_bipartito,*_):
    parametros=parametros.split(maxsplit=2)
    validar_parametros(parametros,3)
    tipo=parametros[0]
    try:
        n=int(parametros[1])
    except ValueError:
        raise ValueError("El parámetro debe ser un número válido.")
    canciones = [c.strip() for c in parametros[2].split(">>>>")]
    salida=f.comando_recomendacion(grafo_bipartito,tipo,n,canciones)
    imprimir(salida)
def ejecutar_ciclo(parametros,__,grafo_canciones_repetidas,*_):
    parametros=parametros.split(maxsplit=1)
    validar_parametros(parametros, 2)
    try:
        n = int(parametros[0])
    except ValueError:
        raise ValueError("El parámetro debe ser un número válido.")
    cancion = "".join(parametros[1].strip())
    cancion=cancion.strip()
    salida=f.comando_ciclo(grafo_canciones_repetidas, n, cancion)
    imprimir(salida)


def ejecutar_rango(parametros,__,grafo_canciones_repetidas,*_):
    parametros=parametros.split(maxsplit=1)
    validar_parametros(parametros,2)
    try:
        n=int(parametros[0])
    except ValueError:
        raise ValueError("El parámetro debe ser un número válido.")
    cancion="".join(parametros[1])
    salida=f.comando_rango(grafo_canciones_repetidas, n, cancion)
    imprimir(salida)

def imprimir(mensaje):
    print(mensaje)

def main():
    archivo = sys.argv[1]
    grafo_bipartito,grafo_canciones_repetidas, usuarios_canciones, canciones_usuarios, usuarios_playlist,playlists_nombres,playlists_canciones  = procesar_archivo(archivo)

    for entrada in sys.stdin:
        procesar_entrada(entrada,grafo_bipartito,grafo_canciones_repetidas, usuarios_canciones, canciones_usuarios, usuarios_playlist,playlists_nombres,playlists_canciones)


main()