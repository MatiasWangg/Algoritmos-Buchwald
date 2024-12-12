from TDAGRAFO.grafo import Grafo
from biblioteca import b
"""
usuarios_canciones: Permite conocer qué canciones le gustan a un usuario.
canciones_usuarios: Permite conocer qué usuarios han marcado una canción como parte de sus listas de reproducción.
"""
def cargar_usuarios_canciones(datos):
    usuarios_canciones = {}
    canciones_usuarios = {}
    for id, user_id, track_name, artist, playlist_id, playlist_name, genres in datos:
        cancion = f"{track_name} - {artist}"
        if user_id not in usuarios_canciones:
            usuarios_canciones[user_id] = set()
        usuarios_canciones[user_id].add(cancion)
        
        if cancion not in canciones_usuarios:
            canciones_usuarios[cancion] = set()
        canciones_usuarios[cancion].add(user_id)
    return usuarios_canciones, canciones_usuarios

def cargar_generos_por_cancion(datos):
    generos_por_cancion = {}
    for _id, user_id, track_name, artist, playlist_id, playlist_name, genres in datos:
        cancion = f"{track_name} - {artist}"
        if genres:
            generos_por_cancion[cancion] = set(genres.split(", "))
        else:
            generos_por_cancion[cancion] = set()
    return generos_por_cancion

def construir_lista_canciones_populares(datos):
    popularidad_canciones = {}
    for _id, user_id, track_name, artist, playlist_id, playlist_name, genres in datos:
        cancion = f"{track_name} - {artist}"
        popularidad_canciones[cancion] = popularidad_canciones.get(cancion, 0) + 1
    
    canciones_populares = sorted(popularidad_canciones.items(), key=lambda x: x[1], reverse=True)
    return canciones_populares

def construir_grafo_bipartito(usuarios_canciones, canciones_usuarios):
    grafo_bipartito = Grafo(es_dirigido=False)
    
    for usuario in usuarios_canciones:
        grafo_bipartito.agregar_vertice(usuario)

    for cancion in canciones_usuarios:
        grafo_bipartito.agregar_vertice(cancion)
   
    for usuario, canciones in usuarios_canciones.items():
        for cancion in canciones:
            grafo_bipartito.agregar_arista(usuario, cancion)
    
    return grafo_bipartito



def comando_camino(grafo_bipartito, origen, destino):
    if origen not in grafo_bipartito.obtener_vertices() or destino not in grafo_bipartito.obtener_vertices():
        print(f"Error: Uno o ambos vértices ('{origen}', '{destino}') no existen en el grafo.")
        return
    
    padres, _ = b.camino_minimo_bfs(grafo_bipartito, origen)

    if destino not in padres:
        print(f"No existe un camino entre '{origen}' y '{destino}'.")
        return

    camino = []
    actual = destino
    while actual is not None:
        camino.append(actual)
        actual = padres[actual]

    camino.reverse()
    imprimir(" >>>> ".join(camino))

def comando_mas_importantes(grafo_bipartito, n):
    centralidades = b.centralidad(grafo_bipartito)

    canciones_ordenadas = sorted(centralidades.items(), key=lambda x: x[1], reverse=True)

    canciones_mas_importantes = []
    for i in range(min(n, len(canciones_ordenadas))):
        cancion = canciones_ordenadas[i][0]
        canciones_mas_importantes.append(cancion)

    salida = "; ".join(canciones_mas_importantes)
    imprimir(salida)

def comando_recomendacion():
    pass

def completar_grafo_canciones_repetidas(usuarios_canciones, grafo_canciones_repetidas):
    for canciones_usuario in usuarios_canciones.values():
        for i, cancion1 in enumerate(canciones_usuario):
            for cancion2 in canciones_usuario[i + 1:]:
                if not grafo_canciones_repetidas.estan_unidos(cancion1, cancion2):
                    grafo_canciones_repetidas.agregar_arista(cancion1, cancion2)

def comando_ciclo(grafo, n, cancion):
    ciclo = b.buscar_ciclo(grafo, cancion, n)
    
    if ciclo is None:
        print(f"No se encontró un ciclo de longitud {n} desde {cancion}.")
    else:
        ciclo_str = " --> ".join(ciclo) + f" --> {ciclo[0]}"
        imprimir(ciclo_str)


def comando_rango(grafo, n, cancion):
    cantidad = b.buscar_rango(grafo, cancion, n)
    imprimir(str(cantidad))

def imprimir(mensaje):
    print(mensaje)
