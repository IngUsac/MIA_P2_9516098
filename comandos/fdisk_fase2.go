package comandos

import (
	"fmt"
	"strconv"
	"strings"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)

//Ejecuta:  fdisk -add=...

func EjecutarFDISKAdd(parametros map[string]string) {

	fmt.Println()
	fmt.Println("FDISK ADD")
	fmt.Println("")

	path := parametros["path"]
	name := parametros["name"]

	archivo, err := AbrirDisco(path)

	if err != nil {

		fmt.Println("ERROR: no se pudo abrir el disco")
		return
	}

	defer archivo.Close()

	mbr, err := LeerMBR(archivo)

	if err != nil {

		fmt.Println("ERROR: no se pudo leer el MBR")
		return
	}

	indice := -1

	for i := 0; i < 4; i++ {

		nombre := strings.TrimRight(
			string(mbr.MbrPartitions[i].PartName[:]),
			"\x00",
		)

		if strings.EqualFold(nombre, name) {

			indice = i
			break
		}
	}

	if indice == -1 {

		fmt.Println("ERROR: partición no encontrada")
		return
	}

	fmt.Println()
	fmt.Println("Partición encontrada")
	fmt.Println("Nombre :", name)
	fmt.Println("Índice :", indice)
	fmt.Println("Inicio :", mbr.MbrPartitions[indice].PartStart)
	fmt.Println("Tamaño :", mbr.MbrPartitions[indice].PartSize)


	addBytes, err := convertirAddBytes(
		parametros["add"],
		parametros["unit"],
	)

	if err != nil {

		fmt.Println("ERROR: valor ADD inválido")
		return
	}

	fmt.Println("Incremento :", addBytes, "bytes")

	// No hay cambios que aplicar.
	if addBytes == 0 {

		fmt.Println("No hay cambios que aplicar.")
		return
	}
		

	nuevoTamano := int64(mbr.MbrPartitions[indice].PartSize) + addBytes
	

	if nuevoTamano <= 0 {

		fmt.Println("ERROR: el tamaño resultante de la partición no puede ser menor o igual a cero.")
		return
	}

	fmt.Println("Nuevo tamaño :", nuevoTamano)
	
	siguiente := siguienteParticionOcupada(mbr, indice)
	
	var espacioDisponible int64

	if siguiente == -1 {

		// Última partición
		espacioDisponible =
			int64(mbr.MbrTamano) -
			(int64(mbr.MbrPartitions[indice].PartStart) +
				int64(mbr.MbrPartitions[indice].PartSize))

	} else {

		espacioDisponible =
			int64(mbr.MbrPartitions[siguiente].PartStart) -
			(int64(mbr.MbrPartitions[indice].PartStart) +
				int64(mbr.MbrPartitions[indice].PartSize))
	}

	fmt.Println("Espacio libre contiguo :", espacioDisponible, "bytes")

	// Validar espacio únicamente cuando el incremento es positivo.
	if addBytes > 0 && addBytes > espacioDisponible {

		fmt.Println("ERROR: no existe espacio libre suficiente para ampliar la partición.")
		return
	}

	// Actualizar el tamaño de la partición.
	mbr.MbrPartitions[indice].PartSize = int32(nuevoTamano)

	// Guardar el MBR.
	err = utilidades.EscribirObjeto(
		archivo,
		&mbr,
		0,
	)

	if err != nil {

		fmt.Println("ERROR: no fue posible actualizar el MBR.")
		return
	}

	fmt.Println()
	fmt.Println("Partición actualizada correctamente.")
	fmt.Println("Nuevo tamaño:", nuevoTamano, "bytes")



}

//Ejecuta: fdisk -delete=...

func EjecutarFDISKDelete(parametros map[string]string) {

	fmt.Println()
	fmt.Println("FDISK DELETE")
	fmt.Println()

	fmt.Println("Pendiente de implementar.")
}

func convertirAddBytes(valor string, unidad string) (int64, error) {

	add, err := strconv.ParseInt(valor, 10, 64)

	if err != nil {
		return 0, err
	}

	switch strings.ToUpper(unidad) {

	case "B":
		return add, nil

	case "K":
		return add * 1024, nil

	case "M":
		return add * 1024 * 1024, nil

	default:
		return add * 1024, nil
	}
}

//Obtiene el índice de la siguiente partición ocupada. Retorna -1 si no existe.

func siguienteParticionOcupada(mbr estructuras.MBR, indice int) int {

	actualFin := int64(mbr.MbrPartitions[indice].PartStart) +
		int64(mbr.MbrPartitions[indice].PartSize)

	siguiente := -1
	var menorInicio int64 = 1<<62

	for i := 0; i < 4; i++ {

		if i == indice {
			continue
		}

		// Ignorar entradas libres
		if mbr.MbrPartitions[i].PartStart <= 0 ||
			mbr.MbrPartitions[i].PartSize <= 0 {
			continue
		}

		inicio := int64(mbr.MbrPartitions[i].PartStart)

		if inicio > actualFin && inicio < menorInicio {

			menorInicio = inicio
			siguiente = i
		}
	}

	return siguiente
}

