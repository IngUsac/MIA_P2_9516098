package comandos

import (
	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
	"fmt"
	"os"
	"strings"
)

//----- funcion para hacer el montaje de particiones -----
// Mount monta una partición utilizando:  -path  y  -name
// Flujo:  validar parámetros --> validar montaje duplicado --> (después leer MBR) --> (después buscar partición) -->  (después generar ID)

func Mount(
	parametros map[string]string,
) {

	path := parametros["path"]
	name := parametros["name"]

	if path == "" {

		fmt.Println(
			"ERROR: falta parametro path",
		)

		return
	}

	if name == "" {

		fmt.Println(
			"ERROR: falta parametro name",
		)

		return
	}

	if ExisteParticionMontada(
		path,
		name,
	) {

		fmt.Println(
			"ERROR: particion ya montada",
		)

		return
	}

	archivo, err := AbrirDisco(
		path,
	)

	if err != nil {

		fmt.Println(
			"ERROR: no se pudo abrir el disco",
		)

		return
	}

	defer archivo.Close()

	mbr, err := LeerMBR(
		archivo,
	)

	if err != nil {

		fmt.Println(
			"ERROR: no se pudo leer el MBR",
		)

		return
	}

	fmt.Println()
	fmt.Println("===== MOUNT =====")

	fmt.Println(
		"Path:",
		path,
	)

	fmt.Println(
		"Name:",
		name,
	)

	fmt.Println(
		"Tamano disco:",
		mbr.MbrTamano,
	)

	// Buscar primaria o extendida
	particion, encontrada := BuscarParticionPrimaria(
		mbr,
		name,
	)

if encontrada {

        if particion.PartType == 'E' {

                fmt.Println(
                        "ERROR: no se puede montar una particion extendida",
                )

                return
        }

        numeroDisco := ObtenerNumeroDisco(path)
        letra := ObtenerLetraParticion(path)
        id := GenerarID(numeroDisco, letra)

		RegistrarParticionMontada(
			id,
			path,
			name,
			particion.PartStart,
			particion.PartSize,
			particion.PartType,
		)

        err = ActualizarParticionPrimariaMontada(
                archivo,
                &mbr,
                name,
                id,
        )

        if err != nil {

                fmt.Println(
                        "ERROR actualizando particion:",
                        err,
                )

                return
        }

        fmt.Println()
        fmt.Println("===== PARTICION MONTADA =====")
        fmt.Println("ID:", id)
        fmt.Println("Nombre:", name)
		fmt.Println()

        return
}


	// Buscar lógica

	var ebr estructuras.EBR

	ebr, encontrada = BuscarParticionLogica(
			archivo,
			mbr,
			name,
	)
	if encontrada {

		numeroDisco := ObtenerNumeroDisco(
			path,
		)

		letra := ObtenerLetraParticion(
			path,
		)

		id := GenerarID(
			numeroDisco,
			letra,
		)


		ebrSize := int32(
				utilidades.ObtenerTamano(
						estructuras.EBR{},
				),
		)

		RegistrarParticionMontada(
				id,
				path,
				name,
				ebr.PartStart + ebrSize,
				ebr.PartSize - ebrSize,
				'L',
		)



		/*
		RegistrarParticionMontada(
			id,
			path,
			name,
			ebr.PartStart,
			ebr.PartSize,
			'L',
		)
		*/

		err = ActualizarParticionLogicaMontada(
			archivo,
			mbr,
			name,
		)

		if err != nil {

			fmt.Println(
				"ERROR actualizando logica:",
				err,
			)

			return
		}

		fmt.Println()
		fmt.Println("===== PARTICION MONTADA =====")

		fmt.Println(
			"ID:",
			id,
		)

		fmt.Println(
			"Nombre:",
			name,
		)

		return
	}

	fmt.Println(
		"ERROR: particion no encontrada",
	)
}


//----- fin Mount -----


// ObtenerNumeroDisco devuelve el número asignado a un disco 
// Si el disco ya fue montado anteriormente:	retorna su número existente.
// Si es la primera vez que se monta: 			asigna el siguiente número disponible.
// disco2.mia -> 2

func ObtenerNumeroDisco(
	path string,
) int {

	// Buscar si el disco ya existe
	for _, disco := range estructuras.DiscosMontados {

		if disco.Path == path {
			return disco.Numero
		}
	}

	// Nuevo número de disco
	nuevoNumero := len(
		estructuras.DiscosMontados,
	) + 1

	// Registrar disco
	estructuras.DiscosMontados = append(
		estructuras.DiscosMontados,
		estructuras.DiscoMontado{
			Path:   path,
			Numero: nuevoNumero,
		},
	)

	return nuevoNumero
}

// ObtenerLetraParticion devuelve la letra que corresponde a una nueva partición montada dentro de un mismo disco
// Disco 1: 981A o 981B o  981C
// Disco 2: 982A
// La letra depende de cuántas particiones del mismo disco ya están montadas en memoria.

func ObtenerLetraParticion(
	path string,
) string {

	contador := 0

	for _, particion := range estructuras.ParticionesMontadas {

		if particion.Path == path {
			contador++
		}
	}

	return string(
		rune('A' + contador),
	)
}


// GenerarID construye el identificador de montaje utilizando el formato solicitado por el proyecto.
// Formato:  98 + NumeroDisco + Letra
//
// Ejemplos: 
// Disco1 -> 981A 
// Disco1 -> 981B
// Disco2 -> 982A

func GenerarID(
	numeroDisco int,
	letra string,
) string {

	return fmt.Sprintf(
		"98%d%s",
		numeroDisco,
		letra,
	)
}


// RegistrarParticionMontada agrega una partición a la tabla global de montajes en memoria
// Todavía no lee MBR ni EBR.  Se utiliza para validar la generación de IDs y el almacenamiento de montajes.

func RegistrarParticionMontada(
	id string,
	path string,
	name string,
	start int32,
	size int32,
	tipo byte,
) {

	estructuras.ParticionesMontadas = append(
		estructuras.ParticionesMontadas,
		estructuras.ParticionMontada{
			ID:    id,
			Path:  path,
			Name:  name,
			Start: start,
			Size:  size,
			Type:  tipo,
		},
	)
}

// ExisteParticionMontada verifica si una partición ya fue montada previamente.
// Debe comparar:  Path del disco y  Nombre de la partición
// Debe retornar:    true  -> ya está montada  o  false -> no está montada

func ExisteParticionMontada(
	path string,
	name string,
) bool {

	for _, particion := range estructuras.ParticionesMontadas {

		if particion.Path == path &&
			particion.Name == name {

			return true
		}
	}

	return false
}


// AbrirDisco abre el archivo .mia en modo lectura/escritura.  Centraliza la apertura para reutilizarla en MOUNT y otros comandos.


func AbrirDisco(
	path string,
) (*os.File, error) {

	return os.OpenFile(
		path,
		os.O_RDWR,
		0644,
	)
}


// LeerMBR obtiene el MBR almacenado al inicio del disco.

func LeerMBR(
	archivo *os.File,
) (estructuras.MBR, error) {

	var mbr estructuras.MBR

	err := utilidades.LeerObjeto(
		archivo,
		&mbr,
		0,
	)

	return mbr, err
}

// BuscarParticionPrimaria busca una partición primaria o extendida por nombre dentro del MBR.
// Debe retornar: particion encontrada  true  -> encontrada o false -> no encontrada

func BuscarParticionPrimaria(
	mbr estructuras.MBR,
	nombre string,
) (estructuras.Partition, bool) {

	for _, particion := range mbr.MbrPartitions {

		if particion.PartSize == 0 {
			continue
		}

		nombreActual := utilidades.BytesAString(
			particion.PartName[:],
		)

		if strings.EqualFold(
			nombreActual,
			nombre,
		) {

			return particion, true
		}
	}

	return estructuras.Partition{}, false
}

// ActualizarParticionMontada: Marca una partición como montada y guarda su ID. Esta información queda almacenada dentro del MBR.

func ActualizarParticionPrimariaMontada(
	archivo *os.File,
	mbr *estructuras.MBR,
	nombre string,
	id string,
) error {

	for i := range mbr.MbrPartitions {

		particion := &mbr.MbrPartitions[i]

		if particion.PartSize == 0 {
			continue
		}

		nombreActual := utilidades.BytesAString(
			particion.PartName[:],
		)

		if strings.EqualFold(
			nombreActual,
			nombre,
		) {

			// Montada
			particion.PartStatus = 'M'

			// Correlativo dentro del MBR
			particion.PartCorrelative = int32(i + 1)

			// Limpiar ID anterior
			particion.PartID = [4]byte{}

			// Copiar ID
			copy(
				particion.PartID[:],
				id,
			)

			// Persistir MBR
			err := utilidades.EscribirObjeto(
				archivo,
				mbr,
				0,
			)

			if err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf(
		"particion no encontrada",
	)
}

// ActualizarParticionLogicaMontada: Marca una partición lógica como montada actualizando el campo PartMount del EBR.

func ActualizarParticionLogicaMontada(
	archivo *os.File,
	mbr estructuras.MBR,
	nombre string,
) error {

	_, extendida, existe := ObtenerParticionExtendida(
		mbr,
	)

	if !existe {
		return fmt.Errorf(
			"no existe particion extendida",
		)
	}

	ebr, err := LeerEBR(
		archivo,
		extendida.PartStart,
	)

	if err != nil {
		return err
	}

	for {

		nombreActual := utilidades.BytesAString(
			ebr.PartName[:],
		)

		if strings.EqualFold(
			nombreActual,
			nombre,
		) {

			ebr.PartMount = 'M'

			return utilidades.EscribirObjeto(
				archivo,
				&ebr,
				int64(ebr.PartStart),
			)
		}

		if ebr.PartNext == -1 {
			break
		}

		ebr, err = LeerEBR(
			archivo,
			ebr.PartNext,
		)

		if err != nil {
			return err
		}
	}

	return fmt.Errorf(
		"logica no encontrada",
	)
}

// BuscarParticionLogica recorre la lista enlazada de EBR dentro de la partición extendida buscando una lógica por nombre.
// Debe retornar: EBR encontrado --> true  -> encontrado o false -> no encontrado

func BuscarParticionLogica(
	archivo *os.File,
	mbr estructuras.MBR,
	nombre string,
) (estructuras.EBR, bool) {

	_, extendida, existe := ObtenerParticionExtendida(
		mbr,
	)

	if !existe {
		return estructuras.EBR{}, false
	}

	ebr, err := LeerEBR(
		archivo,
		extendida.PartStart,
	)


	if err != nil {
		return estructuras.EBR{}, false
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

				return ebr, true
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

	return estructuras.EBR{}, false
}



// Busca una partición previamente montada utilizando su ID

func BuscarParticionMontadaPorID(
	id string,
) (estructuras.ParticionMontada, bool) {

	for _, particion := range estructuras.ParticionesMontadas {

		if strings.EqualFold(
			particion.ID,
			id,
		) {

			return particion, true
		}
	}

	return estructuras.ParticionMontada{}, false
}