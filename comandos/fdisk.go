package comandos

import (
	"fmt"
	"os"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)

func EjecutarFDISK(parametros map[string]string) {

	fmt.Println("===== FDISK =====")

	path, existe := parametros["path"]

	if !existe {
		fmt.Println("ERROR: falta parametro -path")
		return
	}

	archivo, err := os.OpenFile(path, os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("ERROR abriendo disco:", err)
		return
	}

	defer archivo.Close()

	var mbr estructuras.MBR

	err = utilidades.LeerObjeto(
		archivo,
		&mbr,
		0,
	)

	if err != nil {
		fmt.Println("ERROR leyendo MBR:", err)
		return
	}

	fmt.Println("===== INFORMACION DEL DISCO =====")
	fmt.Println("Tamano:", mbr.MbrTamano)
	fmt.Println("Signature:", mbr.MbrDiskSignature)
	fmt.Println("Fit:", string(mbr.DskFit))
	fmt.Println("\n")
	fmt.Println("\n===== PARTICIONES =====")

	for i, particion := range mbr.MbrPartitions {

		if particion.PartSize == 0 {
			fmt.Printf("Particion %d: Libre\n", i+1)
			continue
		}

		fmt.Printf(
			"Particion %d: Inicio=%d Tamano=%d\n",
			i+1,
			particion.PartStart,
			particion.PartSize,
		)
	}
	indice := BuscarParticionLibre(mbr)

	fmt.Println()
	fmt.Println("Indice libre encontrado:", indice)

	sizeStr, existe := parametros["size"]

	if !existe {
		return
	}

	name, existe := parametros["name"]

	if !existe {
		return
	}

	fmt.Println("Size recibido:", sizeStr)
	fmt.Println("Nombre recibido:", name)

}

func BuscarParticionLibre(mbr estructuras.MBR) int {

	for i, particion := range mbr.MbrPartitions {

		if particion.PartSize == 0 {
			return i
		}
	}

	return -1
}

func ObtenerTamanoBytes(size int, unit string) int32 {

	switch unit {
	case "K":
		return int32(size * 1024)

	case "M":
		return int32(size * 1024 * 1024)

	default:
		return int32(size * 1024)
	}
}