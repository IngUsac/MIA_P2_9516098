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
	fmt.Println(" FDISK ")
	fmt.Println()

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

	fmt.Println("  INFORMACION DEL DISCO  ")
	fmt.Println("Tamano:", mbr.MbrTamano)
	fmt.Println("Signature:", mbr.MbrDiskSignature)
	fmt.Println("Fit:", string(mbr.DskFit))
	fmt.Println("\n")
	fmt.Println("\n  PARTICIONES ACTIVAS  ")

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

//<>

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

	//----- Verificar nombres duplicados dentro de las particiones lógicas almacenadas en la lista EBR.
	// Evitar que dos particiones lógicas tengan
	// el mismo nombre.

	if ExisteNombreLogico(
			archivo,
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


	// Validar que una partición lógica solo pueda crearse si existe una partición extendida.
	// Primera validación para soportar -type=L. 
	// agui solo se verifica que exista una extendida.

	if tipo == "L" {

		_, _, existe := ObtenerParticionExtendida(
			mbr,
		)

		if !existe {
			fmt.Println(
				"ERROR: no existe una particion extendida",
			)
			return
		}

		// --- verificar que la funcion obtiene el ultimo EBR de la lista enlazada de particiones lógicas.
		// Prueba temporal antes de utilizarla para crear Logica3 y posteriores.

		ultimoEBR, err := ObtenerUltimoEBR(
			archivo,
			mbr,
		)

		if err != nil {

			fmt.Println(
				"ERROR obteniendo ultimo EBR",
			)

			return
		}

		// --- fin verificar que la funcion obtiene el ultimo EBR de la lista enlazada de particiones lógicas.




		// Verificar si la primera partición lógica utilizará el EBR inicial.
		// Detectar si todavía no existen lógicas.


		// Comenzar siempre desde el EBR inicial. 
		// Más adelante se utilizará ObtenerUltimoEBR() cuando ya existan lógicas.
		// Por ahora centralizamos la lectura en una sola variable para simplificar el flujo.

		ebrInicial, err := ObtenerEBRInicial(
				archivo,
				mbr,
		)

		if err != nil {
				fmt.Println(
						"ERROR leyendo EBR inicial",
				)
				return
		}
		if EsPrimeraLogica(
				ebrInicial,
		) {

		//----------- ------primera particion logica ------------------
		// Crear la primera partición lógica reutilizando el EBR inicial de la extendida.
		// Transformar el EBR vacío en una lógica válida.

		CrearPrimeraLogica(
			&ebrInicial,
			sizeBytes,
			name,
			fitByte,
		)

		err = utilidades.EscribirObjeto(
			archivo,
			&ebrInicial,
			int64(ebrInicial.PartStart),
		)

		if err != nil {

			fmt.Println(
				"ERROR escribiendo primera logica",
			)

			return
		}

		// Verificar inmediatamente lo escrito.

			ebrVerificacion, err := LeerEBR(
				archivo,
				ebrInicial.PartStart,
			)

			if err != nil {

				fmt.Println(
					"ERROR verificando primera logica",
				)

				return
			}

			fmt.Println()
			fmt.Println("  LOGICA CREADA  ")
			fmt.Println(
				"Nombre:",
				utilidades.BytesAString(
					ebrVerificacion.PartName[:],
				),
			)
			fmt.Println(
				"Fit:",
				string(ebrVerificacion.PartFit),
			)
			fmt.Println(
				"Start:",
				ebrVerificacion.PartStart,
			)
			fmt.Println(
				"Size:",
				ebrVerificacion.PartSize,
			)
			fmt.Println(
				"Next:",
				ebrVerificacion.PartNext,
			)

			// La primera lógica ya fue creada y verificada.
			// No debe continuar con la lógica utilizada
			// para segundas o posteriores particiones lógicas.
			return
		}  //---- fin primera logica

		//----- ya existen lógicas, se debe crear un nuevo EBR para la nueva lógica.
		//---- apartar memoria para el nuevo EBR de la siguiente lógica.
		// Construir en memoria el siguiente EBR. Validar la estructura antes de escribirla físicamente en el disco.

		nuevoInicio := CalcularSiguienteEBR(
			ultimoEBR,
		)

		// Verificar que la nueva lógica cabe dentro de la partición extendida.
		// Evitar que una lógica exceda los límites físicos de espacio de la extendida.

		_, extendida, existe := ObtenerParticionExtendida(
				mbr,
		)

		if existe {

				finExtendida :=
						extendida.PartStart +
						extendida.PartSize

				finNuevaLogica :=
						nuevoInicio +
						sizeBytes

				if finNuevaLogica > finExtendida {

						fmt.Println(
								"ERROR: espacio insuficiente dentro de la particion extendida",
						)

						return
				}
		}


		nuevoEBR := CrearSiguienteLogica(
			nuevoInicio,
			sizeBytes,
			name,
			fitByte,
		)


		// --- enlazar el EBR actual con el nuevo EBR creado.
		// Actualizar PartNext del EBR actual. Todavía no se escribe el nuevo EBR. solo actualiza el enlace.

		EnlazarEBR(
			&ultimoEBR,
			nuevoInicio,
		)

		// --- Persistir el EBR actualizado
		// Guardar el nuevo valor de PartNext dentro del EBR actual.

		err = EscribirEBR(
			archivo,
			ultimoEBR,
		)


		if err != nil {

			fmt.Println(
				"ERROR escribiendo EBR actual",
			)

			return
		}

		// --- fin persistir el EBR actualizado.

		// --- Persistir el nuevo EBR correspondiente a la nueva partición lógica.
		// Guardar el nuevo nodo de la lista enlazada.

		err = EscribirEBR(
			archivo,
			nuevoEBR,
		)
		if err != nil {

				fmt.Println(
						"ERROR escribiendo nuevo EBR",
				)

				return
		}

		fmt.Println()
		fmt.Println("  LOGICA CREADA  ")

		fmt.Println("Nombre:",
				utilidades.BytesAString(
						nuevoEBR.PartName[:],
				),
		)

		fmt.Println("Fit:",
				string(nuevoEBR.PartFit),
		)

		fmt.Println("Start:",
				nuevoEBR.PartStart,
		)

		fmt.Println("Size:",
				nuevoEBR.PartSize,
		)

		fmt.Println("Next:",
				nuevoEBR.PartNext,
		)

		return

		// -- fin verificar que el nuevo EBR se escribió correctamente leyendo directamente desde el disco.
		// fin ya existen logicas
	}


	CrearParticionPrimaria(
		&mbr,
		indice,
		sizeBytes,
		name,
		tipo,
		fitByte,
	)


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
	// verificar persistencia real y no solamente que el EBR exista en memoria.

	
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
	fmt.Println("  PARTICIONES ACTUALIZADAS  ")

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


	}


func MostrarParticiones(mbr estructuras.MBR) {

	fmt.Println("\n  PARTICIONES  ")

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
//------------------- Verificar si existe partición extendida ------------------
// Obtener la partición extendida existente dentro del disco.
// Las particiones lógicas necesitan conocer dónde inicia y cuánto espacio tiene la extendida.
// Debe retornar:
// indice  -> posición dentro del MBR
// particion -> estructura completa
// existe -> indica si fue encontrada

func ObtenerParticionExtendida(
	mbr estructuras.MBR,
) (int, estructuras.Partition, bool) {

	for i, particion := range mbr.MbrPartitions {

		if particion.PartSize > 0 &&
			particion.PartType == 'E' {

			return i, particion, true
		}
	}

	return -1, estructuras.Partition{}, false
}



// Leer un EBR desde una posición específica del disco. 
// Centralizar la lectura de EBR para reutilizarla cuando existan varias particiones lógicas.

func LeerEBR(
	archivo *os.File,
	inicio int32,
) (estructuras.EBR, error) {

	var ebr estructuras.EBR

	err := utilidades.LeerObjeto(
		archivo,
		&ebr,
		int64(inicio),
	)

	return ebr, err
}


// Determinar si un EBR no contiene una partición lógica.
// El primer EBR creado dentro de una extendida inicia vacío y será reutilizado por la primera  partición lógica.
// Debe retornar
// true  -> EBR vacío
// false -> EBR ocupado

func EBRVacio(
	ebr estructuras.EBR,
) bool {

	return ebr.PartSize == 0
}


// Obtener el primer EBR de la partición extendida. 
// Toda creación de particiones lógicas inicia  leyendo el primer EBR almacenado dentro de la partición extendida.
// debe retornar: EBR encontrado o error de lectura. 
// Centraliza la obtención del primer EBR de la partición extendida, todas las operaciones sobre particiones lógicas comenzarán desde aquí.

func ObtenerEBRInicial(
	archivo *os.File,
	mbr estructuras.MBR,
) (estructuras.EBR, error) {

	_, extendida, existe := ObtenerParticionExtendida(mbr)

	if !existe {
		return estructuras.EBR{}, fmt.Errorf(
			"no existe particion extendida",
		)
	}

	return LeerEBR(
		archivo,
		extendida.PartStart,
	)
}

// --- Recorrer toda la lista enlazada de EBR y devolver el último EBR existente.
// Permitir la creación de una cantidad arbitraria de particiones lógicas.
// debe retornar el 
// último EBR encontrado dentro de la extendida.

func ObtenerUltimoEBR(
	archivo *os.File,
	mbr estructuras.MBR,
) (estructuras.EBR, error) {

	ebr, err := ObtenerEBRInicial(
		archivo,
		mbr,
	)

	if err != nil {
		return estructuras.EBR{}, err
	}

	for ebr.PartNext != -1 {

		ebr, err = LeerEBR(
			archivo,
			ebr.PartNext,
		)

		if err != nil {
			return estructuras.EBR{}, err
		}
	}

	return ebr, nil
}

// --- fin recorrer toda la lista





// Determinar si la partición lógica que se desea crear sería la primera dentro de la extendida. 
// Si el EBR inicial está vacío, la primera lógica reutilizará ese EBR. 
// debe retornar:
// true  -> primera lógica
// false -> ya existen lógicas

func EsPrimeraLogica(
	ebr estructuras.EBR,
) bool {

	return EBRVacio(ebr)
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

// ------ Verificar si ya existe el nombre de una partición lógica dentro de la extendida.
// Evitar nombres duplicados dentro del mismo disco, incluyendo las particiones lógicas.
// La comparación se realiza contra todas las entradas del MBR y también contra los EBR de las particiones lógicas.
// debe retornar:
// true  -> nombre ya existe
// false -> nombre disponible

func ExisteNombreLogico(
    archivo *os.File,
    mbr estructuras.MBR,
    nombre string,
) bool {

    ebr, err := ObtenerEBRInicial(
        archivo,
        mbr,
    )

    if err != nil {
        return false
    }

    for {

        if ebr.PartSize > 0 {

            nombreActual := utilidades.BytesAString(
                ebr.PartName[:],
            )

            if strings.EqualFold(
                nombreActual,
                nombre,
            ) {
                return true
            }
        }

        if ebr.PartNext == -1 {
            break
        }

        ebr, err = LeerEBR(
            archivo,
            ebr.PartNext,
        )

        if err != nil {
            break
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


// Reutilizar el EBR inicial vacío para crear la primera partición lógica dentro de una partición extendida.
// La primera lógica NO crea un nuevo EBR. Se reutiliza el EBR inicial. 
// PartStart se conserva.
// PartNext permanece en -1.

func CrearPrimeraLogica(
	ebr *estructuras.EBR,
	sizeBytes int32,
	nombre string,
	fit byte,
) {

	ebr.PartFit = fit
	ebr.PartSize = sizeBytes
	ebr.PartName = utilidades.StringABytes16(
		nombre,
	)
}

// Calcular la posición donde debe ubicarse el siguiente EBR dentro de la partición extendida. 
// Determinar el inicio del siguiente nodo de la lista enlazada de particiones lógicas.
// Se utiliza la formula:
// siguienteEBR = EBRActual + sizeof(EBR) + tamañoParticionLogica

func CalcularSiguienteEBR(
	ebr estructuras.EBR,
) int32 {

	return ebr.PartStart +
		int32(
			utilidades.ObtenerTamano(
				estructuras.EBR{},
			),
		) +
		ebr.PartSize
}


// Construir un nuevo EBR para una partición lógica posterior.
// Crear el siguiente nodo de la lista enlazada de particiones lógicas.
// El nuevo EBR inicia apuntando al final de la lista (PartNext = -1).

func CrearSiguienteLogica(
	inicio int32,
	sizeBytes int32,
	nombre string,
	fit byte,
) estructuras.EBR {

	return estructuras.EBR{
		PartMount: '0',
		PartFit:   fit,
		PartStart: inicio,
		PartSize:  sizeBytes,
		PartNext:  -1,
		PartName: utilidades.StringABytes16(
			nombre,
		),
	}
}


// Actualizar el enlace hacia el siguiente EBR por medio de apuntar el campo PartNext del EBR actual hacia la posición del siguiente EBR creado.
// Conectar dos nodos de la lista enlazada de particiones lógicas.
// Ejemplo: EBR1.Next -> EBR2

func EnlazarEBR(
	ebr *estructuras.EBR, // apuntador al EBR actual que se desea actualizar
	siguiente int32,
) {
	ebr.PartNext = siguiente
}


// Guardar un EBR en una posición específica del disco. 
// Centralizar la escritura de EBR para reutilizarla en todas las operaciones sobre particiones lógicas.
// debe retornar error de escritura si ocurre algún problema.

func EscribirEBR(
	archivo *os.File,
	ebr estructuras.EBR,
) error {

	return utilidades.EscribirObjeto(
		archivo,
		&ebr,
		int64(ebr.PartStart),
	)
}


