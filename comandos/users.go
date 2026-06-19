package comandos		

import (
	
	"os"
	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
	"strconv"
	"strings"
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


// ObtenerContenidoUsersTXT: Recorre las estructuras EXT2 y devuelve el contenido completo del archivo users.txt

func ObtenerContenidoUsersTXT(
	archivo *os.File,
	sb estructuras.SuperBlock,
) (string, error) {

	inodeSize := int32(
		utilidades.ObtenerTamano(
			estructuras.Inode{},
		),
	)

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FileBlock{},
		),
	)

	posUsersInode := sb.SInodeStart + inodeSize

	inodeUsers, err := LeerInodo(
		archivo,
		posUsersInode,
	)

	if err != nil {
		return "", err
	}

	var contenido string

	for i := 0; i < 15; i++ {

		if inodeUsers.IBlock[i] == -1 {
			break
		}

		posBloque := sb.SBlockStart +
			(inodeUsers.IBlock[i] * blockSize)

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



// BuscarUsuario:  Busca un usuario dentro del contenido de users.txt.

func BuscarUsuario(
	contenido string,
	user string,
) (estructuras.Usuario, bool) {

	lineas := strings.Split(
		contenido,
		"\n",
	)

	for _, linea := range lineas {

		linea = strings.TrimSpace(
			linea,
		)

		if linea == "" {
			continue
		}

		campos := strings.Split(
			linea,
			",",
		)

		// Registro de usuario:
		// UID,U,GRUPO,USER,PASS

		if len(campos) != 5 {
			continue
		}

		if campos[1] != "U" {
			continue
		}

		if !strings.EqualFold(
			campos[3],
			user,
		) {
			continue
		}

		uid, _ := strconv.Atoi(
			campos[0],
		)

		return estructuras.Usuario{
			UID:      int32(uid),
			Grupo:    campos[2],
			User:     campos[3],
			Password: campos[4],
		}, true
	}

	return estructuras.Usuario{}, false
}




// GuardarUsersTXT: Guarda el contenido completo de users.txt utilizando los bloques definidos en el inodo users.txt.
// Actualmente se utilizan hasta 4 bloques y Cada bloque almacena 64 bytes.
// IBlock[0]
// IBlock[1]
// IBlock[2]
// IBlock[3]
// Parámetros:
// archivo   -> disco abierto
// sb        -> SuperBlock de la partición
// contenido -> nuevo contenido de users.txt

func GuardarUsersTXT(
	archivo *os.File,
	sb estructuras.SuperBlock,
	contenido string,
) error {

	inodeSize := int32(
		utilidades.ObtenerTamano(
			estructuras.Inode{},
		),
	)

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FileBlock{},
		),
	)

	// users.txt = inodo 1

	posUsersInode := sb.SInodeStart + inodeSize

	inodeUsers, err := LeerInodo(
		archivo,
		posUsersInode,
	)

	if err != nil {
		return err
	}

	bytesContenido := []byte(
		contenido,
	)

	const tamBloque = 64

	for i := 0; i < 4; i++ {

		if inodeUsers.IBlock[i] == -1 {
			break
		}

		inicio := i * tamBloque

		var file estructuras.FileBlock

		// Limpiar bloque

		for j := range file.BContent {
			file.BContent[j] = 0
		}

		if inicio < len(bytesContenido) {

			fin := inicio + tamBloque

			if fin > len(bytesContenido) {
				fin = len(bytesContenido)
			}

			copy(
				file.BContent[:],
				bytesContenido[inicio:fin],
			)
		}

		posBloque := sb.SBlockStart +
			(inodeUsers.IBlock[i] * blockSize)

		err = EscribirFileBlock(
			archivo,
			file,
			posBloque,
		)

		if err != nil {
			return err
		}
	}

	return nil
}


// ExisteUsuarioActivo: Verifica si un usuario existe y no ha sido eliminado.
// Parámetros: 
// contenido -> contenido completo de users.txt
// user      -> usuario a buscar
// Retorna: True  -> usuario activo encontrado  o  false -> no existe o está eliminado

func ExisteUsuarioActivo(
	contenido string,
	user string,
) bool {

	lineas := strings.Split(
		contenido,
		"\n",
	)

	for _, linea := range lineas {

		linea = strings.TrimSpace(
			linea,
		)

		if linea == "" {
			continue
		}

		campos := strings.Split(
			linea,
			",",
		)

		if len(campos) != 5 {
			continue
		}

		if campos[1] != "U" {
			continue
		}

		if campos[0] == "0" {
			continue
		}

		if strings.EqualFold(
			campos[3],
			user,
		) {
			return true
		}
	}

	return false
}
