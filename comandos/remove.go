package comandos

import (
	"fmt"
	"os"

	"MIA_P1_9516098/estructuras"
)

func EjecutarRemove(parametros map[string]string) {

	fmt.Println()
	fmt.Println("REMOVE")
	fmt.Println()

	path := parametros["path"]

	if path == "" {
		fmt.Println("ERROR: falta parámetro path.")
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
		fmt.Println("ERROR: no se pudo abrir el disco")
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

	inode, numeroInodo, err := ObtenerInodoPorRutaCompleta(
		archivo,
		sb,
		path,
	)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Ruta localizada.")
	fmt.Println("Inodo:", numeroInodo)
	fmt.Println("Tipo:", string(inode.IType))

	// Verificar si es archivo o carpeta

	switch inode.IType {

	case '1':

		fmt.Println("Archivo localizado.")
		err = EliminarArchivo(
			archivo,
			sb,
			numeroInodo,
			inode,
		)

		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("Archivo eliminado correctamente.")


		err = EliminarEntradaPadre(
			archivo,
			sb,
			path,
		)

		if err != nil {
			fmt.Println("ERROR:", err)
		}

		fmt.Println("Entrada eliminada del directorio padre.")


	case '0':

		fmt.Println("Carpeta localizada.")
		err = EliminarCarpetaRecursiva(
			archivo,
			sb,
			numeroInodo,
			inode,
		)

		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("Carpeta eliminada correctamente.")

	default:

		fmt.Println("ERROR: tipo de inodo desconocido.")
	}


	// TODO:
	// - Verificar permisos
	// - Eliminar archivo
	// - Eliminar carpeta recursivamente
	// - Liberar bloques
	// - Liberar inodo
	// - Actualizar bitmaps
}
