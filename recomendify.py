#!/usr/bin/python3
import sys,csv
import comandos as f
import utils as u


# Diccionario global para almacenar los grafos
grafos = {
    "grafo_bipartito": None,
    "grafo_canciones_repetidas": None
}

def procesar_archivo(ruta):
    with open(ruta) as archivo:
        tsv = csv.reader(archivo, delimiter="\t")
        lista_archivo = list(tsv)
        datos = lista_archivo[1:] 

    usuarios_canciones, canciones_usuarios, usuarios_playlist, playlists_nombres, playlists_canciones = u.guardar_datos(datos)
    
    # Construir el grafo bipartito
    grafo_bipartito = u.construir_grafo_bipartito(usuarios_canciones, canciones_usuarios)
    grafos["grafo_bipartito"] = grafo_bipartito  # Guardamos el grafo bipartito en el diccionario global
    
    return usuarios_canciones, canciones_usuarios, usuarios_playlist, playlists_nombres, playlists_canciones

def procesar_entrada(entrada, usuarios_canciones, canciones_usuarios, usuarios_playlist, playlists_nombres, playlists_canciones):
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
    }
    
    comandos_ciclo_rango = {
        "ciclo": ejecutar_ciclo,
        "rango": ejecutar_rango,
    }
    
    if comando not in comandos and comando not in comandos_ciclo_rango:
        raise ValueError(f"Comando no reconocido: {comando}")

    # Si el grafo_canciones_repetidas no ha sido creado, lo creamos solo si se requiere
    if comando in comandos_ciclo_rango and grafos["grafo_canciones_repetidas"] is None:
        grafos["grafo_canciones_repetidas"] = f.completar_grafo_canciones_repetidas(usuarios_canciones)
    
    # Ejecutar los primeros 3 comandos sin necesidad de grafo_canciones_repetidas
    if comando in comandos:
        comandos[comando](parametros, grafos["grafo_bipartito"], grafos["grafo_canciones_repetidas"], usuarios_canciones, canciones_usuarios, usuarios_playlist, playlists_nombres, playlists_canciones)

    # Ejecutar los últimos dos comandos con la condición de que grafo_canciones_repetidas sea creado
    elif comando in comandos_ciclo_rango:
        comandos_ciclo_rango[comando](parametros, grafos["grafo_bipartito"], grafos["grafo_canciones_repetidas"], usuarios_canciones, canciones_usuarios, usuarios_playlist, playlists_nombres, playlists_canciones)

    
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
    usuarios_canciones, canciones_usuarios, usuarios_playlist,playlists_nombres,playlists_canciones  = procesar_archivo(archivo)

    for entrada in sys.stdin:
        procesar_entrada(entrada, usuarios_canciones, canciones_usuarios, usuarios_playlist,playlists_nombres,playlists_canciones)


main()