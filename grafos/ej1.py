from collections import deque


def orden_topologico_unico(grafo):
    grado_entrada = grados_entrada(grafo)
    cola = deque()
    for v in grafo:
        if grado_entrada[v] == 0:
            if not cola.esta_vacia():
                return False
            cola.Encolar(v)
    while not cola.esta_vacia():
        v = cola.Desencolar()
        for w in grafo.adyacentes(v):
            grado_entrada[w] = grado_entrada[w] - 1
            if grado_entrada[w] == 0:
                if not cola.esta_vacia():
                    return False
                cola.encolar(v)
    return True


def grados_entrada(grafo):
    grados = {}
    for v in grafo:
        grados[v] = 0

    for v in grafo:
        for w in grafo.adyacentes(v):
            grafo[w] += 1

    return grados
