# def bipartito(grafo):
#    colores = {}
#    cola = cola_crear()
#    vertice_inicial = grafo.vertice_aleatorio()
#    cola.encolar(vertice_inicial)
#    colores[vertice_inicial] = 0
#
#    while not cola.esta_vacia():
#        v = cola.desencolar()
#        for w in grafo.adyacentes(v):
#            if w in colores:
#                if colores[w] == colores[v]:
#                    return False
#                else:
#                    colores[w] = 1 - colores[v]
#                    cola.encolar(w)
#    return True
