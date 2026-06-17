package estructuras

// Representa un disco que ya fue montado durante la ejecución.
// Se utiliza para asignar el número de disco requerido por el ID.
// Ejemplo:
// Disco1 -> 981A, 981B
// Disco2 -> 982A
type DiscoMontado struct {
	Path   string
	Numero int
}

// Tabla global de discos montados en memoria.
var DiscosMontados []DiscoMontado

// Información de una partición montada.
type ParticionMontada struct {
	ID    string
	Path  string
	Name  string

	Start int32
	Size  int32

	Type  byte
}

// Tabla global de particiones montadas.
var ParticionesMontadas []ParticionMontada


