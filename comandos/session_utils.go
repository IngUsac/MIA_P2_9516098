package comandos

import (
	"MIA_P1_9516098/estructuras"
	"os"
)

func ObtenerParticionYSuperBlock(
	id string,
) (
	*os.File,
	estructuras.ParticionMontada,
	estructuras.SuperBlock,
	error,
) {

	particion, existe := BuscarParticionMontadaPorID(id)

	if !existe {
		return nil,
			estructuras.ParticionMontada{},
			estructuras.SuperBlock{},
			ErrParticionNoMontada
	}

	archivo, err := AbrirDisco(
		particion.Path,
	)

	if err != nil {
		return nil,
			estructuras.ParticionMontada{},
			estructuras.SuperBlock{},
			err
	}

	sb, err := LeerSuperBlock(
		archivo,
		particion.Start,
	)

	if err != nil {

		archivo.Close()

		return nil,
			estructuras.ParticionMontada{},
			estructuras.SuperBlock{},
			err
	}

	return archivo,
		particion,
		sb,
		nil
}