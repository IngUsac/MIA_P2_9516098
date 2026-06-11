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
}