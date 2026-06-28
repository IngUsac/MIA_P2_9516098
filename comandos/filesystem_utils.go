package comandos

import (
	"fmt"
	"os"
	//"strings"

	"MIA_P1_9516098/estructuras"
)


// Obtiene la partición montada correspondiente a la sesión activa.

func ObtenerParticionSesion() (*estructuras.ParticionMontada, error) {

	particion, existe := BuscarParticionMontadaPorID(
		estructuras.SesionActual.ID,
	)

	if !existe {
		return nil, fmt.Errorf("no existe una sesión activa")
	}

	return &particion, nil
}

/*
Abre el disco correspondiente a la sesión actual.
*/
func AbrirDiscoSesion() (*os.File, *estructuras.ParticionMontada, error) {

	particion, err := ObtenerParticionSesion()

	if err != nil {
		return nil, nil, err
	}

	archivo, err := os.OpenFile(
		particion.Path,
		os.O_RDWR,
		0644,
	)

	if err != nil {
		return nil, nil, err
	}

	return archivo, particion, nil
}

/*
Lee el SuperBlock de la partición montada.
*/
func ObtenerSuperBlockSesion(
	archivo *os.File,
	particion *estructuras.ParticionMontada,
) (estructuras.SuperBlock, error) {

	return LeerSuperBlock(
		archivo,
		particion.Start,
	)

}


func BuscarEntradaEnDirectorio(
	archivo *os.File,
	sb estructuras.SuperBlock,
	inodo estructuras.Inode,
	nombre string,
) (estructuras.EntradaDirectorio, error) {

	return estructuras.EntradaDirectorio{}, nil
}