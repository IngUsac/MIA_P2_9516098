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