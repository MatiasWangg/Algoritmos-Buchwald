import biblioteca as b
import pagerank as pr
import heapq
from grafo import Grafo

def comando_camino(grafo_bipartito, origen, destino):
    padres, _ = b.camino_minimo_bfs(grafo_bipartito, origen)
    if destino not in padres:
        print(f"No se encontro recorrido")
        return
    
    camino = b.reconstruir_camino(padres,destino)
    return camino

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
    return salida


def comando_recomendacion(grafo_bipartito, tipo, n, canciones_favoritas):

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
    return salida



def completar_grafo_canciones_repetidas(usuarios_canciones):
    grafo_canciones_repetidas=Grafo()
    for usuario, canciones in usuarios_canciones.items():
        canciones = list(canciones)

        for cancion in canciones:
            if not grafo_canciones_repetidas.vertice_existe(cancion):
                grafo_canciones_repetidas.agregar_vertice(cancion)
        for i in range(len(canciones)):
            for j in range(i + 1, len(canciones)):
                cancion1 = canciones[i]
                cancion2 = canciones[j]

                if not grafo_canciones_repetidas.estan_unidos(cancion1, cancion2):
                    grafo_canciones_repetidas.agregar_arista(cancion1, cancion2, 1)
                else:
                    peso_actual = grafo_canciones_repetidas.peso_arista(cancion1, cancion2)
                    grafo_canciones_repetidas.agregar_arista(cancion1, cancion2, peso_actual + 1)
    return grafo_canciones_repetidas

def comando_ciclo(grafo, n, cancion):
    ciclo = b.buscar_ciclo(grafo, n, cancion)
    
    if ciclo is None:
        return(f"No se encontró un ciclo de longitud {n} desde {cancion}.")
    else:
        ciclo_str = " --> ".join(ciclo)
        return ciclo_str


def comando_rango(grafo, n, cancion):
    cantidad = b.buscar_rango(grafo, n, cancion)
    return cantidad