import heapq
from collections import deque
import grafo
"""
Implementaciones iniciales porque pueden cambiar (Parametros)
"""
def camino_minimo_dijkstra(grafo, origen, destino):
    distancias = {}
    padres = {}
    for v in grafo.obtener_vertices():
        distancias[v] = float("inf")
    distancias[origen] = 0
    padres[origen] = None
    heap = []
    heapq.heappush(heap, (0, origen))
    while len(heap) != 0:
        _, v = heapq.heappop(heap)
        if v == destino:
            return padres, distancias
        for w in grafo.adyacentes(v):
            nueva_distancia = distancias[v] + grafo.peso_arista(v, w)
            if (nueva_distancia < distancias[w]):
                distancias[w] = nueva_distancia
                padres[w] = v
                heapq.heappush(heap, (distancias[w], w))
    return padres, distancias

