package comandos


import (
	"os"
	"fmt"
	"strings"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)

func CambiarNombreEntrada(
	archivo *os.File,
	sb estructuras.SuperBlock,
	ruta string,
	nuevoNombre string,
) error {

	ruta = strings.TrimRight(ruta, "/")

	partes := strings.Split(ruta, "/")

	if len(partes) < 2 {
		return fmt.Errorf("ruta invalida")
	}

	nombreActual := partes[len(partes)-1]

	rutaPadre := "/"

	if len(partes) > 2 {
		rutaPadre += strings.Join(
			partes[1:len(partes)-1],
			"/",
		)
	}

	inodoPadre, _, err := ObtenerInodoPorRutaCompleta(
		archivo,
		sb,
		rutaPadre,
	)

	if err != nil {
		return err
	}

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FolderBlock{},
		),
	)

	for _, bloque := range inodoPadre.IBlock {

		if bloque == -1 {
			continue
		}

		posicion := sb.SBlockStart +
			(bloque * blockSize)

		folder, err := LeerFolderBlock(
			archivo,
			posicion,
		)

		if err != nil {
			continue
		}

		for i := 0; i < 4; i++ {

			nombre := utilidades.BytesAString(
				folder.BContent[i].BName[:],
			)

			if nombre == nuevoNombre {

				return fmt.Errorf(
					"ya existe una entrada con ese nombre",
				)
			}
		}

		for i := 0; i < 4; i++ {

			nombre := utilidades.BytesAString(
				folder.BContent[i].BName[:],
			)

			if nombre == nombreActual {

				for j := range folder.BContent[i].BName {
					folder.BContent[i].BName[j] = 0
				}

				copy(
					folder.BContent[i].BName[:],
					[]byte(nuevoNombre),
				)

				return EscribirFolderBlock(
					archivo,
					folder,
					posicion,
				)
			}
		}
	}

	return fmt.Errorf("entrada no encontrada")
}