package comandos	
import (
	"fmt"
	"os"
	"strings"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)

// LeerInodo:  Lee un inodo desde una posición específica del disco.
// Parámetros: 
// archivo  -> disco abierto
// posicion -> byte donde inicia el inodo
//
// Retorna: Inodo leído o  error en caso de fallo
//
func LeerInodo(
	archivo *os.File,
	posicion int32,
) (estructuras.Inode, error) {

	var inode estructuras.Inode

	err := utilidades.LeerObjeto(
		archivo,
		&inode,
		int64(posicion),
	)

	if err != nil {
		return estructuras.Inode{}, err
	}

	return inode, nil
}

	
// LeerFolderBlock: Lee un bloque de carpeta desde disco.

func LeerFolderBlock(
	archivo *os.File,
	posicion int32,
) (estructuras.FolderBlock, error) {

	var folder estructuras.FolderBlock

	err := utilidades.LeerObjeto(
		archivo,
		&folder,
		int64(posicion),
	)

	if err != nil {
		return estructuras.FolderBlock{}, err
	}

	return folder, nil
}

// LeerFileBlock: Lee un bloque de archivo desde disco.
//
func LeerFileBlock(
	archivo *os.File,
	posicion int32,
) (estructuras.FileBlock, error) {

	var file estructuras.FileBlock

	err := utilidades.LeerObjeto(
		archivo,
		&file,
		int64(posicion),
	)

	if err != nil {
		return estructuras.FileBlock{}, err
	}

	return file, nil
}

// BuscarEntradaEnFolder: Busca un nombre dentro de un FolderBlock.
//
// Parámetros:
// folder -> bloque de carpeta
// nombre -> nombre buscado
//
// Retorna:
// número de inodo
// true si existe
//
func BuscarEntradaEnFolder(
	folder estructuras.FolderBlock,
	nombre string,
) (int32, bool) {

	for _, entrada := range folder.BContent {

		nombreActual := utilidades.BytesAString(
			entrada.BName[:],
		)

		if nombreActual == nombre {

			return entrada.BInodo, true
		}
	}

	return -1, false
}

// ObtenerInodoPorRuta:
// ------------------------------------------------------------
// Busca un archivo dentro de la raíz.
//
// Ejemplo:
// /users.txt
//
// Retorna:
// inodo encontrado
// posición del inodo
//
func ObtenerInodoPorRuta(
	archivo *os.File,
	sb estructuras.SuperBlock,
	ruta string,
) (estructuras.Inode, int32, error) {

	nombre := strings.TrimPrefix(
		ruta,
		"/",
	)

	folder, err := LeerFolderBlock(
		archivo,
		sb.SBlockStart,
	)

	if err != nil {
		return estructuras.Inode{}, 0, err
	}

	numInodo, existe := BuscarEntradaEnFolder(
		folder,
		nombre,
	)

	if !existe {

		return estructuras.Inode{},
			0,
			fmt.Errorf(
				"archivo no existe",
			)
	}

	inodeSize := int32(
		utilidades.ObtenerTamano(
			estructuras.Inode{},
		),
	)

	posicion := sb.SInodeStart +
		(numInodo * inodeSize)

	inode, err := LeerInodo(
		archivo,
		posicion,
	)

	if err != nil {
		return estructuras.Inode{}, 0, err
	}

	return inode,
		posicion,
		nil



}
/*
inode, _, err := ObtenerInodoPorRuta(
	archivo,
	sb,
	"/users.txt",
)
fmt.Println(
	"Inodo encontrado:",
	inode.IUid,
)*/


