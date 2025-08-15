from grafo import Grafo
from collections import deque, defaultdict
from heap import MinHeap


# Camino minimo Dijkstra. Complejidad O(E log V). Siendo V los vertices y E las aristas.
def camino_minimo_dijkstra(grafo, origen, destino, peso):
    dist, padre = {}, {}
    for v in grafo.obtener_vertices():
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    heap = MinHeap()
    heap.encolar(origen, dist[origen])
    while not heap.esta_vacia():
        v = heap.desencolar()
        if destino is not None and v == destino:
            return padre, dist

        for w in grafo.adyacentes(v):
            peso_arista = grafo.peso_arista(v, w)
            nueva_distancia = dist[v] + peso_arista[peso]
            if nueva_distancia < dist[w]:
                dist[w] = nueva_distancia
                padre[w] = v
                heap.encolar(w, dist[w])
    return padre, dist


def camino_minimo_bfs(grafo, origen, destino, peso):
    dist, padres = {}, {}
    visitados = set()
    cola = deque()

    visitados.add(origen)
    padres[origen] = None
    dist[origen] = 0

    cola.append(origen)
    while cola:
        v = cola.popleft()
        if v == destino:
            return padres, dist
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                padres[w] = v
                dist[w] = dist[v] + 1
                cola.append(w)
    return padres, dist


# Orden topologico DFS. Complejidad O(V+E)
def orden_topologico_dfs(grafo):
    visitados = set()
    pila = deque()
    for v in grafo.obtener_vertices():
        if v not in visitados:
            orden_topologico_rec(grafo, v, pila, visitados)
    lista = []
    while pila:
        lista.append(pila.pop())
    return lista


def orden_topologico_rec(grafo, v, pila, visitados):
    visitados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            orden_topologico_rec(grafo, w, pila, visitados)
    pila.append(v)


# Arbol de Tendido Minimo. Algoritmo de Prim. Complejidad O(E log V). Siendo V los vertices y E las aristas.
def mst_prim(grafo, peso):
    origen = grafo.vertice_aleatorio()
    if not origen:
        return Grafo(False)
    visitados = set()
    visitados.add(origen)
    heap = MinHeap()
    arbol = Grafo(False, grafo.obtener_vertices())

    for w in grafo.adyacentes(origen):
        heap.encolar((origen, w), grafo.peso_arista(origen, w)[peso])

    while not heap.esta_vacia():
        v, w = heap.desencolar()
        if w in visitados:
            continue
        arbol.agregar_arista(v, w, grafo.peso_arista(v, w))
        visitados.add(w)
        for i in grafo.adyacentes(w):
            if i not in visitados:
                heap.encolar((w, i), grafo.peso_arista(w, i)[peso])

    return arbol


# Centralidad.
def algoritmo_centralidad(grafo, peso):
    cent = {}

    for v in grafo.obtener_vertices():
        cent[v] = 0

    for v in grafo.obtener_vertices():

        if peso is None:
            padre, distancia = camino_minimo_bfs(grafo, v)
        else:
            padre, distancia = camino_minimo_dijkstra(grafo, v, None, peso)

        cent_aux = {}
        for w in grafo.obtener_vertices():
            cent_aux[w] = 0
        vertices = ordenar_vertices(grafo, distancia)

        for w in vertices:
            if w == v or padre[w] is None:
                continue
            cent_aux[padre[w]] += 1 + cent_aux[w]
        for w in grafo.obtener_vertices():
            if w != v:
                cent[w] += cent_aux[w]
    cent = dict(sorted(cent.items(), key=lambda x: (-x[1], x[0])))
    return cent


def ordenar_vertices(grafo, distancia):
    vertices = grafo.obtener_vertices()
    vertices_ordenados = sorted(vertices, key=lambda v: distancia[v], reverse=True)

    return vertices_ordenados
