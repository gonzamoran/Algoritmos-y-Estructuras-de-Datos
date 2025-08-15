package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())

	// mas pruebas para este caso...
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})
}

func TestColaVaciaEncolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia())

	for i := 1; i < 10; i++ {
		cola.Encolar(i)
	}

	primero := cola.VerPrimero()
	require.EqualValues(t, 1, primero, "El primero deberia ser 1")

}

func TestColaVaciaEncolarYDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia())

	for i := 1; i < 10; i++ {
		cola.Encolar(i)
	}

	cola.Desencolar()
	primero := cola.VerPrimero()
	require.EqualValues(t, 2, primero, "El primero deberia ser 2")

}

func TestPilaString(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	require.True(t, cola.EstaVacia())

	cola.Encolar("Esta")
	cola.Encolar("Es")
	cola.Encolar("Una")
	cola.Encolar("Prueba")
	cola.Encolar("De TP_COLA")

	require.EqualValues(t, "Esta", cola.Desencolar())
	require.EqualValues(t, "Es", cola.Desencolar())

}
func TestColaVolumen(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())

	for i := 0; i <= 10000; i++ {
		cola.Encolar(i)
	}
	for i := 0; i <= 10000; i++ {
		require.EqualValues(t, i, cola.VerPrimero())
		cola.Desencolar()
	}

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}

func TestColaEncolarVolumen(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())

	for i := 0; i <= 1000000; i++ {
		cola.Encolar(i)
	}
	for i := 0; i <= 1000000; i++ {
		require.EqualValues(t, i, cola.VerPrimero())
		cola.Desencolar()
	}

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}
