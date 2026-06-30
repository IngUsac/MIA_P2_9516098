package comandos

import (
	"fmt"
	"os"

	"MIA_P1_9516098/estructuras"
)

func EjecutarMove(parametros map[string]string) {

	fmt.Println()
	fmt.Println("MOVE")
	fmt.Println()

	path := parametros["path"]
	destino := parametros["destino"]

	if path == "" {
		fmt.Println("ERROR: falta parametro path")
		return
	}

	if destino == "" {
		fmt.Println("ERROR: falta parametro destino")
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
		fmt.Println("ERROR: particion no montada")
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

	err = MoverEntrada(
		archivo,
		&sb,
		path,
		destino,
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

	fmt.Println("Movimiento realizado correctamente.")
}