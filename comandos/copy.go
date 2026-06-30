package comandos

import (
	"fmt"
	"os"
	"path/filepath"

	"MIA_P1_9516098/estructuras"
)

func EjecutarCopy(parametros map[string]string) {

	fmt.Println()
	fmt.Println("COPY")
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

	inodoOrigen, _, err := ObtenerInodoPorRutaCompleta(
		archivo,
		sb,
		path,
	)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	inodoDestino, numeroDestino, err := ObtenerInodoPorRutaCompleta(
		archivo,
		sb,
		destino,
	)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	if inodoDestino.IType != '0' {
		fmt.Println("ERROR: destino no es un directorio")
		return
	}

	nombre := filepath.Base(path)

	if inodoOrigen.IType == '1' {

		_, err = CopiarArchivo(
			archivo,
			&sb,
			path,
			numeroDestino,
			nombre,
		)

	} else {

		_, err = CopiarDirectorio(
			archivo,
			&sb,
			path,
			numeroDestino,
			nombre,
		)
	}

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

	fmt.Println("Copia realizada correctamente.")
}