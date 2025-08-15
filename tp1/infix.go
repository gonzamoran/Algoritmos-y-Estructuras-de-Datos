package main

import (
	"tdas/cola"
	"tdas/pila"
)

const (
	menos              = "-"
	mas                = "+"
	multiplicacion     = "*"
	division           = "/"
	potencia           = "^"
	parentesisApertura = "("
	parentesisClausura = ")"
	vacio              = " "
)

func Precedencia(operador string) int {
	switch operador {
	case mas, menos:
		return 1
	case multiplicacion, division:
		return 2
	case potencia:
		return 3
	default:
		return 0
	}
}
func EsAsociativaDer(operador string) bool {
	return operador == potencia
}

func EsOperador(operador string) bool {
	return operador == mas || operador == menos || operador == multiplicacion || operador == division || operador == potencia
}

func ConversionNotacion(entradaEstandar []string) cola.Cola[string] {
	colaFinal := cola.CrearColaEnlazada[string]()
	pilaOperadores := pila.CrearPilaDinamica[string]()

	for _, token := range entradaEstandar {

		if !EsOperador(token) && token != parentesisApertura && token != parentesisClausura {
			colaFinal.Encolar(token)
		} else if EsOperador(token) {
			for !pilaOperadores.EstaVacia() && (Precedencia(pilaOperadores.VerTope()) > Precedencia(token) || Precedencia(pilaOperadores.VerTope()) == Precedencia(token) && !EsAsociativaDer(token)) {
				colaFinal.Encolar(pilaOperadores.Desapilar())
			}
			pilaOperadores.Apilar(token)
		} else if token == parentesisApertura {
			pilaOperadores.Apilar(token)
		} else if token == parentesisClausura {
			for !pilaOperadores.EstaVacia() && pilaOperadores.VerTope() != parentesisApertura {
				colaFinal.Encolar(pilaOperadores.Desapilar())
			}
			pilaOperadores.Desapilar()
		}
	}

	for !pilaOperadores.EstaVacia() {
		colaFinal.Encolar(pilaOperadores.Desapilar())
	}
	return colaFinal
}

// funcion para que sea valida la expresion que sea leida de stdin.
func Tokenizar(expresion string) []string {
	var tokens []string
	var tokenActual string

	for _, r := range expresion {
		if !EsOperador(string(r)) && string(r) != parentesisApertura && string(r) != parentesisClausura && string(r) != vacio {
			tokenActual += string(r)
		} else if EsOperador(string(r)) || string(r) == parentesisApertura || string(r) == parentesisClausura {
			if tokenActual != "" {
				tokens = append(tokens, tokenActual)
				tokenActual = ""
			}
			tokens = append(tokens, string(r))
		}
	}
	if tokenActual != "" {
		tokens = append(tokens, tokenActual)
	}
	return tokens
}
