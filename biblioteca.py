import heapq
from collections import deque
"""
Implementaciones iniciales porque pueden cambiar (Parametros)
"""
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
    # Marca el vértice actual como visitado y lo agrega al camino
    visitados.add(v)
    camino.append(v)

    # Si el camino tiene exactamente n vértices
    if len(camino) == n:
        # Verifica si el último vértice está conectado al origen para cerrar el ciclo
        if origen in grafo.adyacentes(v):  
            camino.append(origen)  # Completa el ciclo agregando el origen
            return camino
        # Si no se puede cerrar el ciclo, retrocede
        camino.pop()
        visitados.remove(v)
        return None

    # Explorar vecinos no visitados
    for vecino in grafo.adyacentes(v):
        # Permite la visita a un vecino si no ha sido visitado o si es el origen
        if vecino not in visitados or (len(camino) == n - 1 and vecino == origen):
            resultado = dfs(grafo, vecino, origen, visitados, camino, n)
            if resultado: 
                return resultado

    # Si no se encontró el ciclo, retrocede
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