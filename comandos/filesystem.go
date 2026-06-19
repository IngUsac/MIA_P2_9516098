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

// ObtenerInodoPorRuta: Busca un archivo dentro de la raíz.
// Ejemplo:  /users.txt
// Retorna: 
// inodo encontrado
// posición del inodo

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


// LeerContenidoArchivo: Lee el contenido completo de un archivo utilizando los bloques apuntados por su inodo.
// Parámetros:
// archivo -> disco abierto
// sb      -> SuperBlock
// inode   -> inodo del archivo
// Retorna: contenido completo del archivo.

func LeerContenidoArchivo(
	archivo *os.File,
	sb estructuras.SuperBlock,
	inode estructuras.Inode,
) (string, error) {

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FileBlock{},
		),
	)

	var contenido string

	for i := 0; i < 15; i++ {

		if inode.IBlock[i] == -1 {
			break
		}

		posBloque := sb.SBlockStart +
			(inode.IBlock[i] * blockSize)

		fileBlock, err := LeerFileBlock(
			archivo,
			posBloque,
		)

		if err != nil {
			return "", err
		}

		contenido += utilidades.BytesAString(
			fileBlock.BContent[:],
		)
	}

	return contenido, nil
}

// SepararRuta:  Convierte una ruta absoluta en sus componentes.
// Ejemplo:  "/home/docs/test"
// Retorna:  ["home", "docs", "test"]

func SepararRuta(
	ruta string,
) []string {

	ruta = strings.TrimSpace(
		ruta,
	)

	ruta = strings.TrimPrefix(
		ruta,
		"/",
	)

	if ruta == "" {
		return []string{}
	}

	componentes := strings.Split(
		ruta,
		"/",
	)

	return componentes
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

// BuscarComponenteEnInodo: Busca un nombre dentro de los FolderBlocks apuntados por un inodo de carpeta.
// Parámetros:
// archivo    -> disco abierto
// sb         -> SuperBlock
// inode      -> inodo carpeta
// componente -> nombre a buscar
// Retorna: número de inodo encontrado y  true si existe

func BuscarComponenteEnInodo(
	archivo *os.File,
	sb estructuras.SuperBlock,
	inode estructuras.Inode,
	componente string,
) (int32, bool) {

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FolderBlock{},
		),
	)

	for i := 0; i < 15; i++ {

		if inode.IBlock[i] == -1 {
			break
		}

		posBloque := sb.SBlockStart +
			(inode.IBlock[i] * blockSize)

		folder, err := LeerFolderBlock(
			archivo,
			posBloque,
		)

		if err != nil {
			continue
		}

		for _, entrada := range folder.BContent {

			nombre := utilidades.BytesAString(
				entrada.BName[:],
			)

			if nombre == componente {

				return entrada.BInodo,
					true
			}
		}
	}

	return -1, false

}

// ObtenerInodoPorRutaCompleta: Recorre una ruta absoluta componente por componente.
// Ejemplo:  /home/docs/proyecto
// Retorna:  inodo encontrado y  número de inodo,  error si no existe.

func ObtenerInodoPorRutaCompleta(
	archivo *os.File,
	sb estructuras.SuperBlock,
	ruta string,
) (estructuras.Inode, int32, error) {

	componentes := SepararRuta(
		ruta,
	)

	// Ruta raíz
	if len(componentes) == 0 {

		posRaiz := ObtenerPosicionInodo(
			sb,
			0,
		)

		inodeRaiz, err := LeerInodo(
			archivo,
			posRaiz,
		)

		if err != nil {
			return estructuras.Inode{}, 0, err
		}

		return inodeRaiz, 0, nil
	}

	// Empezar desde la raíz
	numeroActual := int32(0)

	posActual := ObtenerPosicionInodo(
		sb,
		numeroActual,
	)

	inodeActual, err := LeerInodo(
		archivo,
		posActual,
	)

	if err != nil {
		return estructuras.Inode{}, 0, err
	}

	for _, componente := range componentes {

		numeroSiguiente, existe :=
			BuscarComponenteEnInodo(
				archivo,
				sb,
				inodeActual,
				componente,
			)

		if !existe {

			return estructuras.Inode{},
				0,
				fmt.Errorf(
					"ruta no existe: %s",
					ruta,
				)
		}

		posSiguiente := ObtenerPosicionInodo(
			sb,
			numeroSiguiente,
		)

		inodeActual, err = LeerInodo(
			archivo,
			posSiguiente,
		)

		if err != nil {
			return estructuras.Inode{}, 0, err
		}

		numeroActual = numeroSiguiente
	}

	return inodeActual,
		numeroActual,
		nil
}