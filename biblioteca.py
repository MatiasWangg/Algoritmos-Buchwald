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

def camino_minimo_bfs(grafo, origen):
    padres = {}
    orden = {}
    visitados = set()
    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)
    bfs(grafo, origen, padres, visitados, orden)
    return padres, orden

def bfs(grafo, vertice, padres, visitados, orden):
    cola = deque()
    cola.append(vertice)
    while cola:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                padres[w] = v
                orden[w] = orden[v]+1
                cola.append(w)

def orden_topologico_dfs(grafo):
    visitados = set()
    pila = deque()
    for v in grafo.obtener_vertices():
        if v not in visitados:
            visitados.add(v)
            dfs(grafo, v, visitados, pila)
        resultado = []
        while pila:
            resultado.append(pila.pop())
        return resultado

def dfs(grafo, vertice, visitados, pila):
    for w in grafo.adyacentes(vertice):
        if w not in visitados:
            visitados.add(w)
            dfs(grafo, w, visitados, pila)
    pila.append(vertice)

def centralidad(grafo):
    centralidad = {}
    for v in grafo.obtener_vertices():
        centralidad[v] = 0
    for v in grafo.obtener_vertices():
        padre, distancias = camino_minimo_dijkstra(grafo, v, None)
        centralidadAux = {}
        for w in grafo.obtener_vertices():
            centralidadAux[w] = 0
        verticesOrdenados = ordenar_vertices(distancias)
        for w in verticesOrdenados:
            if padre[w] == None:
                continue
            centralidadAux[padre[w]] += 1 + centralidadAux[w]
        for w in grafo.obtener_vertices():
            if v == w:
                continue
            centralidad[w] += centralidadAux[w]
    return centralidad

def ordenar_vertices(distancias):
    vertices_ordenadas = sorted(distancias.keys(), key=lambda v: distancias[v], reverse=True)

    vertices_filtradas =[]
    for vertice in vertices_ordenadas:
        if distancias[vertice] != float("inf"):
            vertices_filtradas.append(vertice)
    return vertices_filtradas

