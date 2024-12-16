from grafo import Grafo
"""
usuarios_canciones: Permite conocer qué canciones le gustan a un usuario.
canciones_usuarios: Permite conocer qué usuarios han marcado una canción como parte de sus listas de reproducción.
"""
def formatear_camino(camino, usuarios_playlists, playlists_nombres, playlists_canciones):
    salida = []
    
    # Recorre el camino en posiciones impares para saltar entre las canciones
    for i in range(1, len(camino) - 1, 2):
        cancion_anterior = camino[i - 1]  
        usuario = camino[i]
        cancion_siguiente = camino[i + 1]
        
        # Verifica si el usuario tiene playlists
        if usuario in usuarios_playlists:
            lista_de_playlists = list(usuarios_playlists[usuario])
        else:
            # Si no hay playlists para el usuario, continúa con el siguiente ciclo
            continue
        
        # Obtén las playlists donde aparece la canción
        playlist_anterior = playlist_de_cancion(lista_de_playlists, playlists_canciones, cancion_anterior)
        playlist_siguiente = playlist_de_cancion(lista_de_playlists, playlists_canciones, cancion_siguiente)
        
        # Asegúrate de que las playlists existen en playlists_nombres
        if playlist_anterior in playlists_nombres:
            nombre_playlist_anterior = playlists_nombres[playlist_anterior]
        else:
            nombre_playlist_anterior = "Desconocida"
        
        if playlist_siguiente in playlists_nombres:
            nombre_playlist_siguiente = playlists_nombres[playlist_siguiente]
        else:
            nombre_playlist_siguiente = "Desconocida"
        
        # Construir la parte del camino
        salida.append(f"{cancion_anterior} --> aparece en playlist --> {nombre_playlist_anterior} --> de --> {usuario} --> tiene una playlist --> {nombre_playlist_siguiente} --> donde aparece -->")
    
    # Añadir la última canción
    salida.append(f"{camino[-1]}")
    
    # Unir todos los fragmentos y retornar el resultado final
    return " ".join(salida)
    """cancion-->playlist[cancion1]-->usuario-->playlist[cancion2]-->"""

def playlist_de_cancion(lista_playlists,playlists_canciones,cancion):
    for playlist in lista_playlists:
        if cancion in playlists_canciones[playlist]:
            return playlist
    else:
        print("NO ESTA")

def guardar_datos(datos):
    usuarios_canciones = {}
    canciones_usuarios = {}
    usuarios_playlist={}
    playlist_canciones={}
    playlists_nombres={}
    for _, user_id, track_name, artist, playlist_id,playlist_name, _ in datos:
        cancion= f"{track_name} - {artist}"
        cargar_playlist_canciones(playlist_canciones,playlist_id,cancion)
        cargar_usuarios_canciones(usuarios_canciones,user_id,cancion)
        cargar_playlists_nombres(playlists_nombres,playlist_id,playlist_name)
        cargar_usuarios_playlists(usuarios_playlist,playlist_id,user_id)
        cargar_canciones_usuarios(canciones_usuarios,user_id,cancion)
    return usuarios_canciones, canciones_usuarios,usuarios_playlist,playlists_nombres,playlist_canciones

def cargar_usuarios_canciones(usuarios_canciones,user_id,cancion):
    if user_id not in usuarios_canciones:
        usuarios_canciones[user_id] = set()
    usuarios_canciones[user_id].add((cancion))

def cargar_canciones_usuarios(canciones_usuarios,user_id,cancion):
    if cancion not in canciones_usuarios:
        canciones_usuarios[cancion] = set()
    canciones_usuarios[cancion].add(user_id)

def cargar_playlist_canciones(playlist_canciones, playlist_id,cancion):
    if playlist_id not in playlist_canciones:
        playlist_canciones[playlist_id]=set()
    playlist_canciones[playlist_id].add(cancion)


def cargar_usuarios_playlists(usuarios_playlist, playlist_id, user_id):
    if user_id not in usuarios_playlist:
        usuarios_playlist[user_id]=set()
    usuarios_playlist[user_id].add(playlist_id)

def cargar_playlists_nombres(playlists_nombres,playlist_id,playlist_name):
    if playlist_id not in playlists_nombres:
        playlists_nombres[playlist_id]=playlist_name 

def cargar_generos_por_cancion(datos):
    generos_por_cancion = {}
    for _id, user_id, track_name, artist, playlist_id, playlist_name, genres in datos:
        cancion = f"{track_name} - {artist}"
        if genres:
            generos_por_cancion[cancion] = set(genres.split(", "))
        else:
            generos_por_cancion[cancion] = set()
    return generos_por_cancion

def construir_grafo_bipartito(usuarios_canciones, canciones_usuarios):
    grafo_bipartito = Grafo()
    
    for usuario in usuarios_canciones:
        grafo_bipartito.agregar_vertice(usuario)

    for cancion in canciones_usuarios:
        grafo_bipartito.agregar_vertice(cancion)
    for usuario, canciones in usuarios_canciones.items():
        for cancion in canciones:
            grafo_bipartito.agregar_arista(usuario, cancion)
    return grafo_bipartito

def construir_lista_canciones_populares(datos):
    popularidad_canciones = {}
    for _id, user_id, track_name, artist, playlist_id, playlist_name, genres in datos:
        cancion = f"{track_name} - {artist}"
        popularidad_canciones[cancion] = popularidad_canciones.get(cancion, 0) + 1
    
    canciones_populares = sorted(popularidad_canciones.items(), key=lambda x: x[1], reverse=True)
    return canciones_populares
