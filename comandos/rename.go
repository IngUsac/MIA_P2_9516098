package comandos

import (
	"fmt"
	"path/filepath"
	"strings"
)

func EjecutarRename(parametros map[string]string) {

	fmt.Println()
	fmt.Println("RENAME")
	fmt.Println("")

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

	fmt.Println("Ruta :", path)
	fmt.Println("Nuevo nombre :", name)

	directorio := filepath.Dir(path)
	nombreActual := filepath.Base(path)

	fmt.Println("Directorio   :", directorio)
	fmt.Println("Nombre actual:", nombreActual)

	partes := strings.Split(
		strings.Trim(directorio, "/"),
		"/",
	)

	fmt.Println("Ruta dividida:", partes)

	fmt.Println()
	fmt.Println("Base de RENAME preparada.")	

	
}