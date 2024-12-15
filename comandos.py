import biblioteca as b
import pagerank as pr
import heapq

def comando_camino(grafo_bipartito, origen, destino):
    if origen not in grafo_bipartito.obtener_vertices() or destino not in grafo_bipartito.obtener_vertices():
        print("Tanto el origen como el destino deben ser canciones")
        return
    
    padres, _ = b.camino_minimo_bfs(grafo_bipartito, origen)

    if destino not in padres:
        print(f"No se encontro recorrido")
        return
    
    camino = b.reconstruir_camino(padres,destino)
    print(camino)
    imprimir(" >>>> ".join(camino))

def separar_nodos(grafo_bipartito):
    canciones=[]
    usuarios=[]
    for v in grafo_bipartito.obtener_vertices():
        if " " not in v:
            usuarios.append[v]
        else:
            canciones.append[v]

    return usuarios,canciones


def comando_mas_importantes(grafo_bipartito, C):
    # Obtener el PageRank
    pagerank = pr.pagerank_general(grafo_bipartito)
    pagerank_canciones = pr.separar_pr_por_canciones(pagerank)

    # Crear una lista de tuplas (puntaje, canción)
    canciones = [(puntaje, cancion) for cancion, puntaje in pagerank_canciones.items()]

    # Usar heapq.nlargest para obtener las C canciones con los mayores puntajes
    canciones_mas_importantes = heapq.nlargest(C, canciones, key=lambda x: x[0])

    # Formatear la salida como una cadena separada por punto y coma
    salida = "; ".join(cancion for _, cancion in canciones_mas_importantes)
    imprimir(salida)

    return canciones_mas_importantes


def comando_recomendacion(grafo_bipartito, tipo, n, canciones_favoritas):
    """
    Recomendación personalizada de canciones o usuarios.
    
    tipo: 'canciones' o 'usuarios' (especifica si se recomienda canciones o usuarios)
    n: número de recomendaciones
    canciones_favoritas: canciones favoritas para personalizar la recomendación
    """
    # Obtener el Personalized PageRank
    pagerank = pr.pagerank_recomendacion(grafo_bipartito, canciones_favoritas)
    
    if tipo == 'canciones':
        # Filtrar solo canciones y que no esten en las favoritas
        recomendaciones = [(cancion, puntaje) for cancion, puntaje in pagerank.items() if ' - ' in cancion and cancion not in canciones_favoritas]
    else:  # tipo == 'usuarios'
        # Filtrar usuarios
        recomendaciones = [(usuario, puntaje) for usuario, puntaje in pagerank.items() if ' - ' not in usuario]
    
    # Ordenar las recomendaciones por puntaje
    recomendaciones_ordenadas = sorted(recomendaciones, key=lambda x: x[1], reverse=True)
    
    # Formatear la salida
    salida = "; ".join([recomendacion[0] for recomendacion in recomendaciones_ordenadas[:n]])
    imprimir(salida)
























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
