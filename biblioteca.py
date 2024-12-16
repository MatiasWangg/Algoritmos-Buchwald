import heapq
from collections import deque
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

def dfs(grafo, v, origen, visitados, camino, n):
    visitados.add(v)
    camino.append(v)

    if len(camino) == n:
        if origen in grafo.adyacentes(v):  
            camino.append(origen)
            return camino
        camino.pop()
        visitados.remove(v)
        return None

    # Explorar vecinos no visitados
    for vecino in grafo.adyacentes(v):
        if vecino not in visitados:
            resultado = dfs(grafo, vecino, origen, visitados, camino, n)
            if resultado: 
                return resultado

    camino.pop()
    visitados.remove(v)
    return None


def buscar_ciclo(grafo, n, origen):
    if n < 3:
        return None
    if origen not in grafo.obtener_vertices():
        return None
    if len(grafo.obtener_vertices()) < n:
        return None


    visitados = set()
    camino = []

    return dfs(grafo, origen, origen, visitados, camino, n)



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


def buscar_rango(grafo, n, cancion):
    if n < 0:
        return 0
    
    visitados = set()
    distancias = {}
    cantidad = 0
    visitados.add(cancion)
    distancias[cancion] = 0
    cola = deque()
    cola.append(cancion)

    while cola:
        v = cola.popleft()
        
        # Se procesan los vértices adyacentes
        for w in grafo.adyacentes(v):  
            if w not in visitados:
                visitados.add(w)
                distancias[w] = distancias[v] + 1

                # Solo agregamos a la cola si estamos dentro del rango
                if distancias[w] < n:
                    cola.append(w)
                elif distancias[w] == n:
                    cantidad += 1  # Contamos solo los que están exactamente a n saltos

    return cantidad

    
def reconstruir_camino(padres,destino):
    camino = []
    actual = destino
    while actual is not None:
        camino.append(actual)
        actual = padres[actual]
    camino.reverse()
    return camino