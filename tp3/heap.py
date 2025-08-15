import heapq


class MinHeap:
    def __init__(self):
        self.heap = []

    def encolar(self, elem, prioridad):
        heapq.heappush(self.heap, (prioridad, elem))

    def esta_vacia(self):
        return len(self.heap) == 0

    def desencolar(self):
        if self.esta_vacia():
            raise IndexError("Heap vacio.")
        return heapq.heappop(self.heap)[1]
