package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

}

func TestPilaVaciaApilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	for i := 0; i < 10; i++ {
		pila.Apilar(i)
	}
	pila.Desapilar()
	tope := pila.VerTope()

	require.EqualValues(t, 8, tope, "El tope deberia ser 8")

}

func TestPilaString(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pila.EstaVacia())

	pila.Apilar("Esta")
	pila.Apilar("Es")
	pila.Apilar("Una")
	pila.Apilar("Prueba")
	pila.Apilar("TP_Pila")

	tope := pila.VerTope()

	require.EqualValues(t, "TP_Pila", tope)
	pila.Desapilar()

	nuevo_tope := pila.VerTope()
	require.EqualValues(t, "Prueba", nuevo_tope)

}

func TestPilaApilarVolumen1000(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	for i := 0; i < 1000; i++ {
		pila.Apilar(i)
		pila.VerTope()
	}
	for i := 999; i >= 0; i-- {
		require.EqualValues(t, i, pila.VerTope())
		ultimo := pila.Desapilar()
		require.EqualValues(t, i, ultimo)
	}

	require.True(t, pila.EstaVacia())
}

func TestPilaApilarVolumen(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	for i := 0; i < 1000000; i++ {
		pila.Apilar(i)
		pila.VerTope()
	}
	for i := 999999; i >= 0; i-- {
		require.EqualValues(t, i, pila.VerTope())
		ultimo := pila.Desapilar()
		require.EqualValues(t, i, ultimo)
	}

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

}
