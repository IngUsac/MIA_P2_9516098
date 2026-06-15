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

	fmt.Println("Tamano EBR:", utilidades.ObtenerTamano(estructuras.EBR{}))

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
			"Particion %d -> Nombre=%s Tipo=%c Fit=%c Inicio=%d Tamano=%d\n",
			i+1,
			nombre,
			particion.PartType,
			particion.PartFit,
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


	// Leer el fit solicitado para la partición. Almacenar FF/BF/WF dentro de la estructura de la partición.
	// En esta etapa solamente se guarda el valor. Todavía NO se aplica el algoritmo de asignación.

	fit := "WF"

	if valor, ok := parametros["fit"]; ok {
		fit = strings.ToUpper(valor)
	}

	// DEBUG - Verificar que FDISK recibe correctamente el parámetro fit.
	fmt.Println("Fit recibido:", fit)

		
	// Convertir FF/BF/WF al valor almacenado dentro de la estructura Partition.
	// Preparar el valor para guardarlo en PartFit. Todavía no se aplica ningún algoritmo solo se almacena el identificador.

	fitByte := byte('W')

	switch fit {

	case "FF":
		fitByte = 'F'

	case "BF":
		fitByte = 'B'
	}


	sizeBytes := ObtenerTamanoBytes(size, unit)

	if tipo == "E" && ExisteExtendida(mbr) {
		fmt.Println("ERROR: ya existe una particion extendida")
		return
	}


	
	// Evitar que existan particiones con nombres repetidos.
	// Validación de nombre unico, no pueden existir dos particiones con el mismo nombre dentro del disco.

	if ExisteNombreParticion(
		mbr,
		name,
	) {
		fmt.Println(
			"ERROR: ya existe una particion con ese nombre",
		)
		return
	}

	//------------------- Verificar espacio para una nueva particion ------------------
	// Verificar que la nueva partición no exceda el tamaño físico del disco.
	// Validación de espacio disponible. Antes de crear cualquier partición se debe comprobar que cabe completamente dentro del disco.

	if !HayEspacioDisponible(
		mbr,
		sizeBytes,
	) {
		fmt.Println(
			"ERROR: espacio insuficiente en el disco",
		)
		return
	}


	for i, p := range mbr.MbrPartitions {
		fmt.Printf(
			"DEBUG P%d Tipo=%c Inicio=%d Tamano=%d\n",
			i+1,
			p.PartType,
			p.PartStart,
			p.PartSize,
		)
	}


	CrearParticionPrimaria(
		&mbr,
		indice,
		sizeBytes,
		name,
		tipo,
		fitByte,
	)

	fmt.Println("DEBUG: entrando a validacion EBR")  //  para ver si entra a la parte de creación del EBR
	
	// Si se creó una partición extendida, crear el primer EBR vacío
	fmt.Println("DEBUG tipo =", tipo)

	if tipo == "E" {

										// PASO FDISK-EXT-01:  Crear automáticamente el EBR inicial dentro de la extendida.
	ebr := CrearEBRVacio(
		mbr.MbrPartitions[indice].PartStart,
	)

	err = utilidades.EscribirObjeto(
		archivo,
		&ebr,
		int64(ebr.PartStart),
	)

	if err != nil {
		fmt.Println("ERROR escribiendo EBR")
		return
	}

	fmt.Println("EBR inicial creado en:", ebr.PartStart)


		// FDISK-EXT-02: Verificar que el EBR recién escrito puede leerse desde disco.
		// verificar persistencia real y no solamente que el EBR exista en memoria.

	var ebrLeido estructuras.EBR

	err = utilidades.LeerObjeto(
		archivo,
		&ebrLeido,
		int64(ebr.PartStart),
	)

	if err != nil {
		fmt.Println("ERROR leyendo EBR")
		return
	}

	fmt.Println("===== EBR LEIDO =====")
	fmt.Println("Start:", ebrLeido.PartStart)
	fmt.Println("Size:", ebrLeido.PartSize)
	fmt.Println("Next:", ebrLeido.PartNext)
	}

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
			"Particion %d -> Nombre=%s Tipo=%c Fit=%c Inicio=%d Tamano=%d\n",
			i+1,
			nombre,
			particion.PartType,
			particion.PartFit,
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

	// Crear una partición almacenando también el tipo de ajuste (fit). Guardar F/B/W dentro de PartFit.

	func CrearParticionPrimaria(
		mbr *estructuras.MBR,
		indice int,
		sizeBytes int32,
		nombre string,
		tipo string,
		fit byte,
	) {

	inicio := int32(utilidades.ObtenerTamano(*mbr))

	// Si existen particiones anteriores, el inicio debe calcularse después de la última.
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
		PartFit:    fit,
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


		// Convertir el nombre almacenado en bytes a texto legible para mostrarlo.

		nombre := utilidades.BytesAString(
			particion.PartName[:],
		)

		// Mostrar el Fit almacenado en cada partición. 
		// Verificar que el valor F/B/W persiste correctamente en el MBR.

		fmt.Printf(
			"Particion %d -> Nombre=%s Tipo=%c Fit=%c Inicio=%d Tamano=%d\n",
			i+1,
			nombre,
			particion.PartType,
			particion.PartFit,
			particion.PartStart,
			particion.PartSize,
		)
		
		
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

//------------------- Verificar nombre repetido ------------------
// Verificar si ya existe una partición con el mismo nombre.
// Evitar nombres duplicados dentro del mismo disco.
// La comparación se realiza contra todas las entradas del MBR que estén ocupadas.

func ExisteNombreParticion(
	mbr estructuras.MBR,
	nombre string,
) bool {

	for _, particion := range mbr.MbrPartitions {

		if particion.PartSize == 0 {
			continue
		}

		nombreActual := utilidades.BytesAString(
			particion.PartName[:],
		)

		if strings.EqualFold(	// devuelve true si los nombres son iguales, ignorando mayúsculas/minúsculas
			nombreActual,
			nombre,
		) {
			return true
		}
	}

	return false
}

//------------------- Verificar espacio para una nueva particion ------------------
// Verificar si la nueva partición cabe dentro del disco.
// Evitar que una partición sobrepase el tamaño total  del disco. Se calcula usando el último byte ocupado por cualquier partición existente.
// Esta función calcula cuál sería el final de la nueva partición y verifica que no exceda el tamaño total del disco almacenado en el MBR.

func HayEspacioDisponible(
	mbr estructuras.MBR,
	sizeBytes int32,
) bool {

	inicio := int32(
		utilidades.ObtenerTamano(mbr),
	)

	for _, particion := range mbr.MbrPartitions {

		if particion.PartSize == 0 {
			continue
		}

		fin := particion.PartStart +
			particion.PartSize

		if fin > inicio {
			inicio = fin
		}
	}

	return inicio+sizeBytes <= mbr.MbrTamano
}



func CrearEBRVacio(inicio int32) estructuras.EBR {

	return estructuras.EBR{
		PartMount: '0',
		PartFit:   'W',
		PartStart: inicio,
		PartSize:  0,
		PartNext:  -1,
	}
}