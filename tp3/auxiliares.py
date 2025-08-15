import csv
from grafo import Grafo
from biblioteca import *


def cargar_archivos(entrada):
    csv_aeropuerto = entrada[1]
    csv_vuelos = entrada[2]

    grafo = Grafo()
    aeropuertos, ciudades = {}, {}

    with open(csv_aeropuerto, "r") as f_aeropuertos, open(csv_vuelos, "r") as f_vuelos:
        reader_a = csv.reader(f_aeropuertos)
        reader_v = csv.reader(f_vuelos)

        for fila in reader_a:

            codigo = fila[1]
            grafo.agregar_vertices(codigo)
            aeropuertos[codigo] = fila[2:]

            ciudad = fila[0]
            ciudades[ciudad] = ciudades.get(ciudad, [])
            ciudades[ciudad].append(codigo)

        for fila in reader_v:
            aero_i, aero_j = fila[0], fila[1]
            tiempo, precio, cant_vuelos = int(fila[2]), int(fila[3]), int(fila[4])

            grafo.agregar_arista(
                aero_i, aero_j, [precio, tiempo, cant_vuelos, 1 / cant_vuelos]
            )

    return grafo, ciudades, aeropuertos


def procesar_input(entrada):
    stdin = entrada.split(" ", 1)
    comando = stdin[0]
    argumentos = stdin[1]
    if (
        comando == "centralidad"
        or comando == "itinerario"
        or comando == "nueva_aerolinea"
        or comando == "exportar_kml"
    ):
        return comando, argumentos
    argumentos = argumentos.split(",")

    return comando, argumentos


# ---------------------------------------------------------------
# -----------------Funciones auxiliares comandos-----------------
def camino_menos_costoso(caminos, grafo):
    precios_camino = {}
    precios_finales = []

    for camino_actual in caminos:
        precio_camino_actual = 0
        for i in range(len(camino_actual) - 1):
            v = camino_actual[i]
            w = camino_actual[i + 1]
            peso = grafo.peso_arista(v, w)
            if peso is None:
                precio_camino_actual = float("inf")
            else:
                precio_conex = peso[0]
            precio_camino_actual += precio_conex
        precios_camino[precio_camino_actual] = camino_actual
        precios_finales.append(precio_camino_actual)
    return precios_camino[min(precios_finales)]


def encontrar_camino(ciudad_origen, ciudad_destino, grafo, ciudades, criterio):
    caminos_minimos = []
    if criterio == "rapido":
        peso = 1
    elif criterio == "barato":
        peso = 0
    for aeropuerto_origen in ciudades[ciudad_origen]:
        for aeropuerto_destino in ciudades[ciudad_destino]:
            padres, _ = camino_minimo_dijkstra(grafo, aeropuerto_origen, None, peso)
            camino = reconstruir_camino(padres, aeropuerto_origen, aeropuerto_destino)
            caminos_minimos.append(camino)

    if criterio == "barato":
        return camino_menos_costoso(caminos_minimos, grafo)
    else:
        return min(caminos_minimos, key=len)


def obtener_camino_menos_escalas(origen, destino, grafo, ciudades):
    caminos_validos = []

    for aeropuerto_origen in ciudades[origen]:
        for aeropuerto_destino in ciudades[destino]:
            padres, _ = camino_minimo_bfs(
                grafo, aeropuerto_origen, aeropuerto_destino, peso=None
            )
            if aeropuerto_destino not in padres:
                continue

            camino = reconstruir_camino(padres, aeropuerto_origen, aeropuerto_destino)
            caminos_validos.append(camino)
    if not caminos_validos:
        return []
    return min(caminos_validos, key=len)


def k_mas_importantes(grafo, k, peso):
    centr = algoritmo_centralidad(grafo, peso)
    aeropuertos_importantes = [codigo for codigo, _ in list(centr.items())[:k]]
    print(", ".join(aeropuertos_importantes))
    return aeropuertos_importantes


def escribir_rutas_mst(grafo, arbol, ruta_salida):
    vuelos = []
    for v, w, peso in arbol.obtener_aristas():
        vuelos.append(
            [
                v,
                w,
                str(grafo.peso_arista(v, w)[1]),
                str(peso[0]),
                str(grafo.peso_arista(v, w)[2]),
            ]
        )
    with open(ruta_salida, "w", newline="") as archivo:
        writer = csv.writer(archivo)
        writer.writerows(vuelos)


def grafo_ciudades(archivo):
    grafo = Grafo(True)
    with open(archivo, "r") as f_ciudades:
        reader = csv.reader(f_ciudades)
        ciudades = next(reader)
        for ciudad in ciudades:
            grafo.agregar_vertices(ciudad)
        for linea in reader:
            grafo.agregar_arista(linea[0], linea[1])
    return grafo


def ejecutar_itinerario(archivo_itinerario, ciudades, grafo):
    grafo_de_ciudades = grafo_ciudades(archivo_itinerario)
    ruta = orden_topologico_dfs(grafo_de_ciudades)

    for ciudad in ruta:
        if ciudad not in ciudades:
            print(f"Ciudad {ciudad} no existe en el grafo original.")

    print(", ".join(ruta))
    for i in range(len(ruta) - 1):
        origen = ruta[i]
        destino = ruta[i + 1]
        camino = obtener_camino_menos_escalas(origen, destino, grafo, ciudades)
        print(" -> ".join(camino))


def escribir_kml(archivo, camino, aeropuertos):

    with open(archivo, "w") as f:
        f.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        f.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
        f.write("   <Document>\n")
        for ciudad in camino:
            lat, lon = aeropuertos[ciudad]

            f.write("       <Placemark>\n")
            f.write(f"          <name>{ciudad}</name>\n")
            f.write("           <Point>\n")
            f.write(f"              <coordinates>{lon}, {lat} </coordinates>\n")
            f.write("           </Point>\n")
            f.write("       </Placemark>\n")
        f.write("\n")
        for i in range(len(camino) - 1):
            origen = camino[i]
            destino = camino[i + 1]
            lat1, lon1 = aeropuertos[origen]
            lat2, lon2 = aeropuertos[destino]
            f.write("       <Placemark>\n")
            f.write("           <LineString>\n")
            f.write(
                f"                  <coordinates>{lon1}, {lat1} {lon2}, {lat2} </coordinates>\n"
            )
            f.write("           </LineString>\n")
            f.write("       </Placemark>\n")

        f.write("   </Document>\n")
        f.write("</kml>\n")


# ----------------------------------------------------
# -----------Mas Funciones Auxiliares-----------------
# -----------------------------------------------------


def reconstruir_camino(padres, origen, destino):
    vertice = destino
    camino = []
    while vertice != origen:
        camino.append(vertice)
        vertice = padres[vertice]
    camino.append(origen)
    camino.reverse()
    return camino
