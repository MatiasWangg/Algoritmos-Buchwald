from grafo import Grafo
"""
usuarios_canciones: Permite conocer qué canciones le gustan a un usuario.
canciones_usuarios: Permite conocer qué usuarios han marcado una canción como parte de sus listas de reproducción.
usuarios_play: Permite conocer las playlists pertenecientes a un usuario
canciones_play: Permite conocer en que playlists esta cada cancion
play_nombres: Permite saber el nombre de cada playlist por su id
"""
def guardar_datos(datos):
    usuarios_canciones = {}
    canciones_usuarios = {}
    usuarios_play={}
    canciones_play={}
    play_nombres={}
    for _, user_id, track_name, artist, playlist_id,playlist_name, _ in datos:
        cancion= f"{track_name} - {artist}"
        cargar_canciones_play(canciones_play,playlist_id,cancion)
        cargar_usuarios_canciones(usuarios_canciones,user_id,cancion)
        cargar_play_nombres(play_nombres,playlist_id,playlist_name)

        cargar_usuarios_play(usuarios_play,playlist_id,user_id)
        cargar_canciones_usuarios(canciones_usuarios,user_id,cancion)

    return usuarios_canciones, canciones_usuarios,usuarios_play,play_nombres,canciones_play

def cargar_usuarios_canciones(usuarios_canciones,user_id,cancion):
    if user_id not in usuarios_canciones:
        usuarios_canciones[user_id] = set()
    usuarios_canciones[user_id].add((cancion))

def cargar_canciones_usuarios(canciones_usuarios,user_id,cancion):
    if cancion not in canciones_usuarios:
        canciones_usuarios[cancion] = set()
    canciones_usuarios[cancion].add(user_id)

def cargar_canciones_play(canciones_play, playlist_id,cancion):
    if cancion not in canciones_play:
        canciones_play[cancion]=set()
    canciones_play[cancion].add(playlist_id)


def cargar_usuarios_play(usuarios_play, playlist_id, user_id):
    if user_id not in usuarios_play:
        usuarios_play[user_id]=set()
    usuarios_play[user_id].add(playlist_id)

def cargar_play_nombres(play_nombres,playlist_id,playlist_name):
    if playlist_id not in play_nombres:
        play_nombres[playlist_id]=playlist_name 

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

def formatear_camino(camino, usuarios_playlists, play_nombres, canciones_play):
    salida = []
    
    # Recorre el camino en posiciones impares para saltar entre las canciones
    for i in range(1, len(camino) - 1, 2):
        cancion_anterior = camino[i - 1]  
        usuario = camino[i]
        cancion_siguiente = camino[i + 1]
        
        # Verifica si el usuario tiene playlists
        if usuario in usuarios_playlists:
            playlists_del_usuario = usuarios_playlists[usuario]

        # Obtén la playlist donde aparece la canción y los nombres
        playlist_anterior = buscar_playlist(playlists_del_usuario, canciones_play[cancion_anterior])
        playlist_siguiente = buscar_playlist(playlists_del_usuario, canciones_play[cancion_siguiente])
        
        # Asegúrate de que las playlists son válidas
        nombre_playlist_anterior = play_nombres[playlist_anterior] if playlist_anterior else "Desconocida"
        nombre_playlist_siguiente = play_nombres[playlist_siguiente] if playlist_siguiente else "Desconocida"

        # Construir la parte del camino
        salida.append(f"{cancion_anterior} --> aparece en playlist --> {nombre_playlist_anterior} --> de --> {usuario} --> tiene una playlist --> {nombre_playlist_siguiente} --> donde aparece -->")
    
    # Añadir la última canción
    salida.append(f"{camino[-1]}")
    
    # Unir todos los fragmentos y retornar el resultado final
    return " ".join(salida)

def buscar_playlist(playlists_del_usuario, playlists_de_cancion):
    # Iterar directamente sobre las playlists del usuario
    for playlist in playlists_del_usuario:
        if playlist in playlists_de_cancion:
            return playlist  # Retorna la primera playlist válida
    return None  # Aunque este caso se ha eliminado, sigue presente por seguridad
