package comandos

import (
	"fmt"
	"os"

	"MIA_P1_9516098/estructuras"
)

func EjecutarRename(parametros map[string]string) {

	fmt.Println()
	fmt.Println("RENAME")
	fmt.Println()

	path := parametros["path"]
	name := parametros["name"]

	if path == "" {
		fmt.Println("ERROR: falta parametro path")
		return
	}

	if name == "" {
		fmt.Println("ERROR: falta parametro name")
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

	_, _, err = ObtenerInodoPorRutaCompleta(
		archivo,
		sb,
		path,
	)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	err = CambiarNombreEntrada(
		archivo,
		sb,
		path,
		name,
	)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Nombre actualizado correctamente.")
}


/*

package comandos

import (
	"fmt"
	"path/filepath"
)

func EjecutarRename(parametros map[string]string) {

	fmt.Println()
	fmt.Println("RENAME")
	fmt.Println()

	path := parametros["path"]
	name := parametros["name"]

	if path == "" {
		fmt.Println("ERROR: falta parámetro path.")
		return
	}

	if name == "" {
		fmt.Println("ERROR: falta parámetro name.")
		return
	}

	directorio := filepath.Dir(path)
	nombreActual := filepath.Base(path)

	fmt.Println("Ruta          :", path)
	fmt.Println("Directorio    :", directorio)
	fmt.Println("Nombre actual :", nombreActual)
	fmt.Println("Nuevo nombre  :", name)

	fmt.Println()
	fmt.Println("Base de RENAME preparada.")
	fmt.Println("Pendiente: cambiar la entrada dentro del FolderBlock.")
}

*/