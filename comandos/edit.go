package comandos

import (
	"fmt"
	"os"

	"MIA_P1_9516098/estructuras"
)

func EjecutarEdit(parametros map[string]string) {

	fmt.Println()
	fmt.Println("EDIT")
	fmt.Println()

	path := parametros["path"]
	cont := parametros["cont"]

	if path == "" {

		fmt.Println("ERROR: falta parametro path")
		return
	}

	if cont == "" {

		fmt.Println("ERROR: falta parametro cont")
		return
	}

	if !estructuras.SesionActual.Activa {

		fmt.Println("ERROR: no existe una sesion activa")
		return
	}

	particion, existe := BuscarParticionMontadaPorID(
		estructuras.SesionActual.ID,
	)

	if !existe {

		fmt.Println("ERROR: particion no encontrada")
		return
	}

	archivo, err := os.OpenFile(
		particion.Path,
		os.O_RDWR,
		0644,
	)

	if err != nil {

		fmt.Println("ERROR abriendo disco")
		return
	}

	defer archivo.Close()

	sb, err := LeerSuperBlock(
		archivo,
		particion.Start,
	)

	if err != nil {

		fmt.Println("ERROR leyendo SuperBlock")
		return
	}

	inodo, numeroInodo, err := ObtenerInodoPorRutaCompleta(
		archivo,
		sb,
		path,
	)

	if err != nil {

		fmt.Println("ERROR:", err)
		return
	}

	if inodo.IType != '1' {

		fmt.Println("ERROR: la ruta no corresponde a un archivo")
		return
	}

	err = EscribirContenidoArchivo(
		archivo,
		&sb,
		&inodo,
		cont,
	)

	if err != nil {

		fmt.Println("ERROR:", err)
		return
	}

	err = GuardarInodo(
		archivo,
		sb,
		numeroInodo,
		inodo,
	)

	if err != nil {

		fmt.Println("ERROR:", err)
		return
	}

	err = EscribirSuperBlock(
		archivo,
		sb,
		particion.Start,
	)

	if err != nil {

		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Archivo editado correctamente.")
}