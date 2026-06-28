package estructuras

type EntradaDirectorio struct {
	NumeroInodo int32
	Bloque      int32
	Posicion    int
	Nombre      string
	Existe      bool
}