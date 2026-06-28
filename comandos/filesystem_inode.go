package comandos

import (
	//"fmt"
	"os"
//	"strings"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)




// Gguarda un inodo en disco.

func EscribirInodo(
	archivo *os.File,
	inode estructuras.Inode,
	posicion int32,
) error {

	return utilidades.EscribirObjeto(
		archivo,
		&inode,
		int64(posicion),
	)
}


// ObtenerPosicionInodo: Calcula la posición física de un inodo.
// Parámetros: 
// sb      -> SuperBlock
// numero  -> número de inodo
// Retorna:  posición en bytes.

func ObtenerPosicionInodo(
	sb estructuras.SuperBlock,
	numero int32,
) int32 {

	inodeSize := int32(
		utilidades.ObtenerTamano(
			estructuras.Inode{},
		),
	)

	return sb.SInodeStart +
		(numero * inodeSize)
}