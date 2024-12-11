
class Grafo_no_dirigido:
    def __init__(self):
        self.grafo={}

    def vertice_existe(self,vertice):
        return vertice in self.grafo

    def añadir_vertice(self,vertice):
        if not self.vertice_existe(vertice):
            self.grafo[vertice]={}
        

    def añadir_arista(self, v1, v2):
        if v1 == v2:
            return 
        if not self.vertice_existe(v1):
            self.añadir_vertice(v1)
        if not self.vertice_existe(v2):
            self.añadir_vertice(v2)
        self.grafo[v1][v2] = 1
        self.grafo[v2][v1] = 1
    
    def borrar_vertice(self, v):
        if not self.vertice_existe(v):
            return
        for adyacente in self.grafo[v]:
            del self.grafo[adyacente][v]
        del self.grafo[v]
        
    def borrar_arista(self,v1,v2):
        if not self.vertice_existe(v1) or not self.vertice_existe(v2):
            return
        del self.grafo[v1][v2]
        del self.grafo[v2][v1]

    def adyacentes(self,vertice):
        if not self.vertice_existe(vertice):
            return
        adya=[]
        for v in self.grafo[vertice]:
            adya.append(v)
        return adya
        
    def obtener_vertices(self):
        return list(self.grafo)
