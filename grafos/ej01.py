import queue


def orden_topo(grafo):
    grados = {}
    for v in grafo:
        grados[v] = 0

    for v in grafo:
        for w in grafo.adyacentes(v):
            grados[w] += 1

    cola = queue()
    for v in grafo:
        if grados[v] == 0:
            cola.add()

    while not cola.is_empty():
        if len(cola) > 1:
            return False
        vertice = cola.pop()
        for w in grafo.adyacentes(vertice):
            grados[w] -= 1
            if grados[w] == 0:
                cola.add(w)

    return True
