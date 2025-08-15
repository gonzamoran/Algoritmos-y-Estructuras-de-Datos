#!/usr/bin/python3

import sys, auxiliares, comandos
from grafo import Grafo


def main():
    entrada = sys.argv
    grafo, ciudades, aeropuertos = auxiliares.cargar_archivos(entrada)
    ruta = []
    while True:
        try:
            linea = input()
            operacion, argumentos = auxiliares.procesar_input(linea)

            if operacion == "camino_mas":
                ruta = comandos.camino_mas(argumentos, grafo, ciudades)
            elif operacion == "camino_escalas":
                ruta = comandos.camino_escalas(argumentos, grafo, ciudades)
            elif operacion == "centralidad":
                comandos.centralidad(argumentos, grafo)
            elif operacion == "nueva_aerolinea":
                comandos.nueva_aerolinea(argumentos, grafo)
            elif operacion == "itinerario":
                comandos.itinerario(argumentos, ciudades, grafo)
            elif operacion == "exportar_kml":
                comandos.exportar_klm(argumentos, ruta, aeropuertos)
            else:
                print(f"Comando desconocido {operacion}")

        except EOFError:
            break


if __name__ == "__main__":
    main()
