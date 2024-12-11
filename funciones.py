from TDAGRAFO.grafo import Grafo
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