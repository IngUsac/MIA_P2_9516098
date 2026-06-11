package comandos

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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

		nombre := utilidades.BytesAString(
			particion.PartName[:],
		)

	
		fmt.Printf(
			"Particion %d -> Nombre=%s Tipo=%c Inicio=%d Tamano=%d\n",
			i+1,
			nombre,
			particion.PartType,
			particion.PartStart,
			particion.PartSize,
		)

	}
	indice := BuscarParticionLibre(mbr)
	if indice == -1 {
		fmt.Println("ERROR: no existen entradas libres para particiones")
		return
	}

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

	size, err := strconv.Atoi(sizeStr)

	if err != nil || size <= 0 {
		fmt.Println("ERROR: size invalido")
		return
	}

	unit := "K"

	tipo := "P"

	if valor, ok := parametros["type"]; ok {
		tipo = strings.ToUpper(valor)
	}

	if valor, ok := parametros["unit"]; ok {
		unit = strings.ToUpper(valor)
	}

	sizeBytes := ObtenerTamanoBytes(size, unit)

	if tipo == "E" && ExisteExtendida(mbr) {
		fmt.Println("ERROR: ya existe una particion extendida")
		return
	}


	CrearParticionPrimaria(
		&mbr,
		indice,
		sizeBytes,
		name,
		tipo,
	)

	err = utilidades.EscribirObjeto(
		archivo,
		&mbr,
		0,
	)

	if err != nil {
		fmt.Println("ERROR escribiendo MBR")
		return
	}

	if err != nil {
		fmt.Println("ERROR escribiendo MBR")
		return
	}

	var mbrActualizado estructuras.MBR

	err = utilidades.LeerObjeto(
		archivo,
		&mbrActualizado,
		0,
	)

	if err != nil {
		fmt.Println("ERROR leyendo MBR actualizado")
		return
	}

	fmt.Println()
	fmt.Println("===== PARTICIONES ACTUALIZADAS =====")

	for i, particion := range mbrActualizado.MbrPartitions {

		if particion.PartSize == 0 {
			fmt.Printf("Particion %d: Libre\n", i+1)
			continue
		}

		nombre := utilidades.BytesAString(
			particion.PartName[:],
		)

		fmt.Printf(
			"Particion %d -> Nombre=%s Tipo=%c Inicio=%d Tamano=%d\n",
			i+1,
			nombre,
			particion.PartType,
			particion.PartStart,
			particion.PartSize,
		)
	}
	
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

func CrearParticionPrimaria(
	mbr *estructuras.MBR,
	indice int,
	sizeBytes int32,
	nombre string,
	tipo string,
) {

	inicio := int32(utilidades.ObtenerTamano(*mbr))

	// Si existen particiones anteriores,
	// el inicio debe calcularse después de la última.
	for _, particion := range mbr.MbrPartitions {

		if particion.PartSize > 0 {

			fin := particion.PartStart + particion.PartSize

			if fin > inicio {
				inicio = fin
			}
		}
	}

	mbr.MbrPartitions[indice] = estructuras.Partition{
		PartStatus: '1',
		PartType: tipo[0],
		PartFit:    'F',
		PartStart:  inicio,
		PartSize:   sizeBytes,
		PartName:   utilidades.StringABytes16(nombre),
	}

	fmt.Println("DEBUG")
	fmt.Println("Indice:", indice)
	fmt.Println("Size:", mbr.MbrPartitions[indice].PartSize)
	fmt.Println("Start:", mbr.MbrPartitions[indice].PartStart)
	fmt.Println("Type:", string(mbr.MbrPartitions[indice].PartType))
			
	}


func MostrarParticiones(mbr estructuras.MBR) {

	fmt.Println("\n===== PARTICIONES =====")

	for i, particion := range mbr.MbrPartitions {

		if particion.PartSize == 0 {
			fmt.Printf("Particion %d: Libre\n", i+1)
			continue
		}

		fmt.Printf(
			"Particion %d -> Nombre=%s Tipo=%c Inicio=%d Tamano=%d\n",
			i+1,
			utilidades.BytesAString(particion.PartName[:]),
			particion.PartType,
			particion.PartStart,
			particion.PartSize,
		)
		/*
		fmt.Printf(
			"Particion %d -> Nombre=%s Tipo=%c Inicio=%d Tamano=%d\n",
			i+1,
			utilidades.BytesAString(particion.PartName[:]),
			particion.PartType,
			particion.PartStart,
			particion.PartSize,
		)*/
	}
}

func ExisteExtendida(mbr estructuras.MBR) bool {

	for _, particion := range mbr.MbrPartitions {

		if particion.PartSize > 0 &&
			particion.PartType == 'E' {

			return true
		}
	}

	return false
}