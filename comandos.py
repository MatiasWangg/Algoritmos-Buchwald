import biblioteca as b


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



def comando_mas_importantes(grafo_bipartito, n):
    pagerank={}
    for v in grafo_bipartito.obtener_vertices():
        pagerank[v]=1.0
    centralidades = aplicar_pagerank(grafo_bipartito,pagerank)

    canciones_ordenadas = sorted(centralidades.items(), key=lambda x: x[1], reverse=True)

    canciones_mas_importantes = []
    for i in range(min(n,len(canciones_ordenadas))):
        cancion = canciones_ordenadas[i][0]
        canciones_mas_importantes.append(cancion)
    salida = "; ".join(canciones_mas_importantes)
    imprimir(salida)




def comando_recomendacion(grafo_bipartito, tipo, n, canciones):
    if tipo == "canciones":
        canciones_recomendadas = recomendar_canciones(grafo_bipartito, canciones, n)
        salida = "; ".join(canciones_recomendadas)
        imprimir(salida)
    
    elif tipo == "usuarios":
        usuarios_recomendados = recomendar_usuarios(grafo_bipartito, canciones, n)
        salida = "; ".join(usuarios_recomendados)
        imprimir(salida)

def recomendar_canciones(grafo_bipartito, canciones_conocidas, n):
    pagerank = {}
    for cancion in canciones_conocidas:
        pagerank[cancion] = 1.0  
    
    pagerank = aplicar_pagerank(grafo_bipartito, pagerank)
    

    canciones_ordenadas = sorted(pagerank.items(), key=lambda x: x[1], reverse=True)

    canciones_recomendadas = []
    for cancion, _ in canciones_ordenadas:
        if cancion not in canciones_conocidas:
            canciones_recomendadas.append(cancion)

    return canciones_recomendadas[:n]

def recomendar_usuarios(grafo_bipartito, canciones_conocidas, n):
    pagerank = {}
    for cancion in canciones_conocidas:
        for usuario in grafo_bipartito.obtener_adyacentes(cancion):
            if usuario not in pagerank:
                pagerank[usuario] = 1.0  

    pagerank = aplicar_pagerank(grafo_bipartito, pagerank)
    usuarios_ordenados = sorted(pagerank.items(), key=lambda x: x[1], reverse=True)
    usuarios_recomendados = []
    for usuario, _ in usuarios_ordenados[:n]:
        usuarios_recomendados.append(usuario)
    return usuarios_recomendados


def aplicar_pagerank(grafo_bipartito,pagerank_inicial,max_iter=10):
    for _ in range(max_iter): 
        nuevo_pagerank = {}
        for vertice in grafo_bipartito.obtener_vertices():
            suma = 0.0
            for vecino in grafo_bipartito.adyacentes(vertice):
                suma += pagerank_inicial.get(vecino, 0) / len(grafo_bipartito.adyacentes(vecino))
            nuevo_pagerank[vertice] = suma
        
        pagerank_inicial = nuevo_pagerank
    
    return pagerank_inicial




























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
