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

	// users.txt está en el inodo 1
	posUsersInode := sb.SInodeStart + inodeSize

	inodeUsers, err := LeerInodo(
		archivo,
		posUsersInode,
	)

	if err != nil {
		return "", err
	}

	// bloque 1
	posUsersBlock := sb.SBlockStart + blockSize

	usersFile, err := LeerFileBlock(
		archivo,
		posUsersBlock,
	)

	if err != nil {
		return "", err
	}

	_ = inodeUsers

	return utilidades.BytesAString(
		usersFile.BContent[:],
	), nil
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




// GuardarUsersTXT:  Sobrescribe el contenido del archivo users.txt dentro del sistema EXT2.
// Recibe:
//   archivo   -> disco abierto
//   sb        -> SuperBlock de la partición
//   contenido -> nuevo contenido completo de users.txt
// Retorna:   error -> si ocurre un problema al escribir.

func GuardarUsersTXT(
	archivo *os.File,
	sb estructuras.SuperBlock,
	contenido string,
) error {

	var file estructuras.FileBlock

	// Copiar el contenido recibido
	copy(
		file.BContent[:],
		contenido,
	)



// users.txt está almacenado en el bloque 1

blockSize := int32(
	utilidades.ObtenerTamano(
		estructuras.FileBlock{},
	),
)

posUsersBlock := sb.SBlockStart + blockSize

// Sobrescribir el bloque en disco
return EscribirFileBlock(
	archivo,
	file,
	posUsersBlock,
)
}
