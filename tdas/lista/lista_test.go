package lista_test

import(
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
