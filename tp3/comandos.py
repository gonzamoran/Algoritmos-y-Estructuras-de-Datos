import csv
from auxiliares import *
from biblioteca import *


def camino_mas(argumentos, grafo, ciudades):
    modo, origen, destino = argumentos[0], argumentos[1], argumentos[2]
    camino = encontrar_camino(origen, destino, grafo, ciudades, modo)
    print(" -> ".join(camino))
    return camino


def camino_escalas(argumentos, grafo, ciudades):
    origen, destino = argumentos[0], argumentos[1]
    camino = obtener_camino_menos_escalas(origen, destino, grafo, ciudades)
    print(" -> ".join(camino))
    return camino


def centralidad(argumentos, grafo):
    k = int(argumentos)
    return k_mas_importantes(grafo, k, 3)


def nueva_aerolinea(argumentos, grafo):
    arbol = mst_prim(grafo, 0)
    escribir_rutas_mst(grafo, arbol, argumentos)
    print("OK")


def itinerario(argumentos, ciudades, grafo):
    ejecutar_itinerario(argumentos, ciudades, grafo)


def exportar_klm(argumentos, ruta, aeropuertos):

    escribir_kml(argumentos, ruta, aeropuertos)
    print("OK")
