import random
pagerank_global = None
def pagerank_general(grafo_bipartito, max_iter=10, damping=0.85):
    global pagerank_global
    
    # Si ya existe un cálculo previo, reutilízalo
    if pagerank_global is not None:
        return pagerank_global
    
    # Inicialización del PageRank
    vertices = grafo_bipartito.obtener_vertices()
    num_vertices = len(vertices)
    pagerank_actual = {v: 1.0 / num_vertices for v in vertices}  # Distribución uniforme inicial
    
    # Iteraciones del algoritmo de PageRank
    for _ in range(max_iter):
        nuevo_pr = {}
        for vertice in vertices:
            suma = 0.0
            for vecino in grafo_bipartito.adyacentes(vertice):
                suma += pagerank_actual[vecino] / len(grafo_bipartito.adyacentes(vecino))
            # Aplicar el factor de amortiguación
            nuevo_pr[vertice] = (1 - damping) / num_vertices + damping * suma
        
        # Actualizar el PageRank actual
        pagerank_actual = nuevo_pr
    
    # Guardar el resultado globalmente
    pagerank_global = pagerank_actual
    return pagerank_global


def pagerank_recomendacion(grafo_bipartito, canciones_favoritas, max_iter=10,camino=10):
    # Inicialización: Asignar 1/len(canciones_favoritas) para las canciones favoritas, 0 para las demás
    pagerank = {v: 0.0 for v in grafo_bipartito.obtener_vertices()}
    total_favoritas = len(canciones_favoritas)
    # Asignar un valor de 1/total_favoritas a las canciones favoritas
    for cancion in canciones_favoritas:
        pagerank[cancion] = 1.0/total_favoritas

    # Realizar múltiples iteraciones de Random Walks
    for _ in range(max_iter):
        cancion_aleatoria=random.choice(canciones_favoritas)
        for _ in range(camino):
            vecino_aleatorio=random.choice(grafo_bipartito.adyacentes(cancion_aleatoria))
            pagerank[vecino_aleatorio]+=pagerank[cancion_aleatoria]/len(grafo_bipartito.adyacentes(cancion_aleatoria))
            cancion_aleatoria=vecino_aleatorio

    return pagerank


def separar_pr_por_canciones(pr):
    pr_canciones={}
    for v in pr:
         if " - " in v:
            pr_canciones[v]=pr[v]
    return pr_canciones