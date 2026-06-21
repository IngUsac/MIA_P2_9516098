package comandos		

import (
	"fmt"
	"os"
	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
	"strconv"
	"strings"
)


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


// GuardarUsersTXT: Guarda el contenido completo de users.txt. Si el contenido crece, reserva nuevos bloques automáticamente.
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

	posUsersInode := sb.SInodeStart +
		inodeSize

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

	cantidadBloques :=
		(len(bytesContenido) + 63) / 64

	if cantidadBloques == 0 {
		cantidadBloques = 1
	}

	if cantidadBloques > 12 {

		return fmt.Errorf(
			"users.txt excede bloques directos",
		)
	}

	for i := 0; i < cantidadBloques; i++ {

		if inodeUsers.IBlock[i] == -1 {

			numBloque, err :=
				BuscarPrimerBloqueLibre(
					archivo,
					sb,
				)

			if err != nil {
				return err
			}

			err = OcuparBloque(
				archivo,
				sb,
				numBloque,
			)

			if err != nil {
				return err
			}

			inodeUsers.IBlock[i] =
				numBloque
		}
	}

	err = EscribirInodo(
		archivo,
		inodeUsers,
		posUsersInode,
	)

	if err != nil {
		return err
	}

	for i := 0; i < cantidadBloques; i++ {

		inicio := i * 64
		fin := inicio + 64

		if fin > len(bytesContenido) {
			fin = len(bytesContenido)
		}

		var file estructuras.FileBlock

		copy(
			file.BContent[:],
			bytesContenido[inicio:fin],
		)

		err = GuardarFileBlock(
			archivo,
			sb,
			inodeUsers.IBlock[i],
			file,
		)

		if err != nil {
			return err
		}
	}

	inodeUsers.ISize =
		int32(len(bytesContenido))

	err = EscribirInodo(
		archivo,
		inodeUsers,
		posUsersInode,
	)

	if err != nil {
		return err
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
