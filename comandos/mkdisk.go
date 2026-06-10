package comandos

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func EjecutarMKDISK(parametros map[string]string) {

	fmt.Println("===== MKDISK =====")

	// Validar SIZE
	sizeStr, existe := parametros["size"]
	if !existe {
		fmt.Println("ERROR: Falta parametro obligatorio -size")
		return
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size <= 0 {
		fmt.Println("ERROR: -size debe ser un numero mayor a cero")
		return
	}

	// Validar PATH
	path, existe := parametros["path"]
	if !existe {
		fmt.Println("ERROR: Falta parametro obligatorio -path")
		return
	}

	// UNIT por defecto M
	unit := "M"
	if valor, ok := parametros["unit"]; ok {
		unit = strings.ToUpper(valor)
	}

	var tamanoBytes int64

	switch unit {
	case "K":
		tamanoBytes = int64(size) * 1024
	case "M":
		tamanoBytes = int64(size) * 1024 * 1024
	default:
		fmt.Println("ERROR: Unit debe ser K o M")
		return
	}

	// Crear carpetas si no existen
	carpeta := filepath.Dir(path)

	err = os.MkdirAll(carpeta, 0755)
	if err != nil {
		fmt.Println("ERROR creando carpetas:", err)
		return
	}

	// Crear archivo
	archivo, err := os.Create(path)
	if err != nil {
		fmt.Println("ERROR creando disco:", err)
		return
	}
	defer archivo.Close()

	// Reservar espacio
	err = archivo.Truncate(tamanoBytes)
	if err != nil {
		fmt.Println("ERROR asignando tamaño:", err)
		return
	}

	fmt.Println("Disco creado correctamente")
	fmt.Println("Ruta:", path)
	fmt.Println("Tamaño:", tamanoBytes, "bytes")
}