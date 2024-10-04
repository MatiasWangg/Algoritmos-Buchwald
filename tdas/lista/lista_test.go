package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLista_alCrearEstaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "Deberia estar vacia al ser creada")
}

func TestLista_verPrimeroPanicsAlEstarVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "VerPrimero deberia causar panico en una lista vacia")
}

func TestLista_borrarPrimeroPanicsAlEstarVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "BorrarPrimero deberia causar panico en una lista vacia")
}

func TestInsertarBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(0)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia despues de insertar")
	require.Equal(t, 0, lista.VerPrimero(), "El primer valor debería ser el insertado")

	lista.InsertarUltimo(1)
	require.Equal(t, 0, lista.VerPrimero(), "El primer valor deberia seguir siendo el mismo")

	borrado := lista.BorrarPrimero()
	require.Equal(t, 0, borrado, "Se deberia haber borrado el primero")
	require.Equal(t, 1, lista.VerPrimero(), "El primer valor ahora deberia ser el segundo insertado")

	borrado = lista.BorrarPrimero()
	require.Equal(t, 1, borrado, "Se deberia haber borrado el ultimo")
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrar todo")
}

func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	tam := 1000

	for i := 0; i < tam; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, 0, lista.VerPrimero())
	}
	for j := 0; j < tam; j++ {
		require.Equal(t, j, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrar todo")
}

func TestListaGenerica(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaInt.InsertarUltimo(10)
	require.Equal(t, 10, listaInt.BorrarPrimero(), "Deberia borrar el entero 10")

	listaString := TDALista.CrearListaEnlazada[string]()
	listaString.InsertarUltimo("Hola")
	require.Equal(t, "Hola", listaString.BorrarPrimero(), "Deberia borrar la cadena 'Hola'")

	listaBool := TDALista.CrearListaEnlazada[bool]()
	listaBool.InsertarUltimo(true)
	require.Equal(t, true, listaBool.BorrarPrimero(), "Deberia borrar el valor booleano true")
}

// Pruebas del iterador Externo

func TestInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente(), "No deberia haber nada para ver")
	iter.Insertar(2.718281828459045235360)
	require.EqualValues(t, lista.VerPrimero(), iter.VerActual(), "El valor primer elemento de la lista debe ser el que inserto el iterador")
	require.True(t, iter.HaySiguiente(), "Deberia haber algo mas para ver")
}

func TestInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int32]()
	lista.InsertarPrimero(24)
	lista.InsertarUltimo(34)
	iter := lista.Iterador()
	require.EqualValues(t, lista.VerPrimero(), iter.VerActual(), "el elemento actual del iterador debe ser el primero de la lista")
	iter.Siguiente()
	require.EqualValues(t, lista.VerUltimo(), iter.VerActual(), "El elemento actual deberia ser el ultimo")
	iter.Siguiente()
	require.False(t, iter.HaySiguiente(), " No deberia haber algo mas para ver")
	iter.Insertar(47)
	require.EqualValues(t, lista.VerUltimo(), iter.VerActual(), "El elemento actual del iterador debe ser el ultimo de la lista")
}

func TestInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(10)
	iter.Siguiente()
	iter.Insertar(20)
	iter.Siguiente()
	iter.Insertar(40)
	require.EqualValues(t, lista.VerUltimo(), iter.VerActual(), "el elemento actual debe ser 40")
	iter.Insertar(30)
	require.EqualValues(t, 30, iter.VerActual(), "se debe haber insertado el 30")
	iter.Siguiente()
	require.EqualValues(t, 40, iter.VerActual(), "el elemento posterior al 30 debe ser el que era el actual antes de insertar")
}

func TestBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("A")
	lista.InsertarUltimo("B")
	iter := lista.Iterador()
	require.EqualValues(t, lista.VerPrimero(), iter.VerActual(), "El elemento del iterador deberia ser el primero de la lista")
	borrado := iter.Borrar()
	require.EqualValues(t, "A", borrado, "El valor borrado debe ser el que estaba de primero")
	require.EqualValues(t, lista.VerPrimero(), iter.VerActual(), "El primero deberia ser ahora el elemento actual")
	require.EqualValues(t, lista.VerUltimo(), iter.VerActual(), "El ultimo deberia ser ahora el elemento actual")
}

func TestBorrarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float32]()
	lista.InsertarUltimo(3.9)
	lista.InsertarPrimero(2.39)
	lista.InsertarPrimero(34.23)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	borrado := iter.Borrar()
	require.EqualValues(t, float32(3.9), borrado, "El elemento borrado deberia ser el ultimo")
	require.EqualValues(t, float32(2.39), lista.VerUltimo(), "El ultimo de la lista deberia ser el anterior al borrado")

}

func TestNoEstaBorrado(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Borrar()
	require.NotEqualValues(t, 20, lista.VerPrimero())
	require.NotEqualValues(t, 20, lista.VerUltimo())
	require.NotEqualValues(t, 20, iter.VerActual())
	require.EqualValues(t, 2, lista.Largo())

}
