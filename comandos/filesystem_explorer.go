package comandos

import (
	"os"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)

type ExplorerEntry struct {
	Name        string `json:"name"`
	IsDirectory bool   `json:"isDirectory"`
}

func ListarDirectorio(
	archivo *os.File,
	sb estructuras.SuperBlock,
	ruta string,
) ([]ExplorerEntry, error) {

	inodo, _, err := ObtenerInodoPorRutaCompleta(
		archivo,
		sb,
		ruta,
	)

	if err != nil {
		return nil, err
	}

	var lista []ExplorerEntry

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FolderBlock{},
		),
	)



	for i := 0; i < 12; i++ {

		if inodo.IBlock[i] == -1 {
			continue
		}

		posBloque := sb.SBlockStart +
			(inodo.IBlock[i] * blockSize)

		folder, err := LeerFolderBlock(
			archivo,
			posBloque,
		)

		if err != nil {
			continue
		}

		for _, entrada := range folder.BContent {

			if entrada.BInodo == -1 {
				continue
			}

			nombre := utilidades.BytesAString(
				entrada.BName[:],
			)

			if nombre == "" ||
				nombre == "." ||
				nombre == ".." {
				continue
			}

			posInodo := ObtenerPosicionInodo(
				sb,
				entrada.BInodo,
			)

			inodoHijo, err := LeerInodo(
				archivo,
				posInodo,
			)

			if err != nil {
				continue
			}

			lista = append(
				lista,
				ExplorerEntry{
					Name:        nombre,
					IsDirectory: inodoHijo.IType == '0',
				},
			)
		}
	}

	return lista, nil

}