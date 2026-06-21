package comandos

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)

func EjecutarMKDISK(parametros map[string]string) {

	fmt.Println(" MKDISK, parametros", parametros )
	fmt.Println()

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

	if !strings.HasSuffix(  // valida la extension .dsk 
		strings.ToLower(path),
		".dsk",
	) {

		fmt.Println(
			"ERROR: el disco debe tener extension .dsk",
		)

		return
	}

	if !existe {
		fmt.Println("ERROR: Falta parametro obligatorio -path")
		return
	}

	// UNIT por defecto M
	unit := "M"
	if valor, ok := parametros["unit"]; ok {
		unit = strings.ToUpper(valor)
	}

	// Leer el tipo de ajuste (fit) recibido en MKDISK. Guardar el fit del disco para usarlo más adelante.
	// Todavía NO se aplica el algoritmo, solo se almacena en el MBR.

	fit := "FF"

	if valor, ok := parametros["fit"]; ok {
		fit = strings.ToUpper(valor)
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

	// Convertir FF/BF/WF al valor que se almacenará dentro del MBR.
	// Preparar el valor para guardarlo en DskFit.

	fitByte := byte('F')

	switch fit {

	case "BF":
		fitByte = 'B'

	case "WF":
		fitByte = 'W'
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


	fecha := time.Now().Format("2006-01-02 15:04")

	mbr := estructuras.MBR{
		MbrTamano:        	int32(tamanoBytes),
		MbrFechaCreacion: 	utilidades.StringABytes20(fecha),
		MbrDiskSignature: 	rand.Int31(),
		DskFit: 			fitByte,
	}

	err = utilidades.EscribirObjeto(archivo, &mbr, 0)

	if err != nil {
		fmt.Println("ERROR escribiendo MBR:", err)
		return
	}


	if err != nil {
		fmt.Println("ERROR asignando tamaño:", err)
		return
	}

	var mbrLeido estructuras.MBR

	err = utilidades.LeerObjeto(
		archivo,
		&mbrLeido,
		0,
	)

	if err != nil {
		fmt.Println("ERROR leyendo MBR:", err)
		return
	}

	fmt.Println("  MBR LEIDO  ")
	fmt.Println("Tamano:", mbrLeido.MbrTamano)
	fmt.Println("Signature:", mbrLeido.MbrDiskSignature)
	fmt.Println("Fit:", string(mbrLeido.DskFit))
	fmt.Println()
	fmt.Println("Disco creado correctamente")
	fmt.Println("Ruta:", path)
	fmt.Println("Tamaño:", tamanoBytes, "bytes")
}