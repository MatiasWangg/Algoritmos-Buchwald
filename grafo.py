import random
class Grafo:
    def __init__(self, es_dirigido = False, vertices_init = []):
        self.dirigido = es_dirigido
        self.vertices = {}
        for vertice in vertices_init:
            self.agregar_vertice(vertice)
    
    def agregar_vertice(self, v):
        if not v in self.vertices:
            self.vertices[v] = {}
    
    def borrar_vertice(self, v):
        if v not in self.vertices:
            raise ValueError(f"No se encuentra el vertice {v} en el grafo")
        self.vertices.pop(v)
        for _, dato in  self.vertices:
            if v in dato:
                dato.pop(v)

    def agregar_arista(self, v, w, peso = 1):
        if v not in self.vertices:
            raise ValueError(f"No se encuentra el vertice {v} en el grafo")
        elif w not in self.vertices:
            raise ValueError(f"No se encuentra el vertice {w} en el grafo")

        self.vertices[v][w] = peso
        if not self.dirigido:
            self.vertices[w][v] = peso
    
    def borrar_arista(self, v, w):
        if v not in self.vertices:
            raise ValueError(f"No se encuentra el vertice {v} en el grafo")
        elif w not in self.vertices:
            raise ValueError(f"No se encuentra el vertice {w} en el grafo")

        self.vertices[v].pop(w)
        if not self.dirigido:
            self.vertices[w].pop(v)
    
    def estan_unidos(self, v, w):
        return v in self.vertices and w in self.vertices
    
    def peso_arista(self, v, w):
        if self. estan_unidos(v, w):
            return self.vertices[v][w]
        return None
    
    def obtener_vertices(self):
        resultado = []
        for vertice in self.vertices:
            resultado.append(vertice)
        return resultado
    
    def adyacentes(self, v):
        adyacentes = []
        if v in self.vertices:
            for w in self.vertices[v]:
                adyacentes.append(w)
        return adyacentes

    def vertice_aleatorio(self):
        vertices = self.obtener_vertices()
        return random.choice(vertices)

