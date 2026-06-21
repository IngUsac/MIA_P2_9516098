package comandos

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"MIA_P1_9516098/estructuras"
)

// MKGRP: Crea un nuevo grupo dentro de users.txt.
// Requisitos:
// - Debe existir una sesión activa.
// - El usuario debe ser root.
// - Debe venir el parámetro name.

func MKGRP(
	parametros map[string]string,
) {
	fmt.Println(" MKGRP ")
	fmt.Println()

	name := parametros["name"]

	if !estructuras.SesionActual.Activa {

		fmt.Println(
			"ERROR: no existe una sesion activa",
		)

		return
	}

	if estructuras.SesionActual.User != "root" {

		fmt.Println(
			"ERROR: solo root puede ejecutar MKGRP",
		)

		return
	}

	if name == "" {

		fmt.Println(
			"ERROR: falta parametro name",
		)

		return
	}

	fmt.Println(
		"Grupo:",
		name,
	)

//**---> 
// ------------------------------------------------------------
// Obtener la partición donde existe la sesión activa
// ------------------------------------------------------------

particion, existe := BuscarParticionMontadaPorID(
	estructuras.SesionActual.ID,
)

if !existe {

	fmt.Println(
		"ERROR: particion no encontrada",
	)

	return
}

// ------------------------------------------------------------
// Abrir disco
// ------------------------------------------------------------

archivo, err := os.OpenFile(
	particion.Path,
	os.O_RDWR,
	0644,
)

if err != nil {

	fmt.Println(
		"ERROR: no se pudo abrir el disco",
	)

	return
}

defer archivo.Close()

// ------------------------------------------------------------
// Leer SuperBlock
// ------------------------------------------------------------

sb, err := LeerSuperBlock(
	archivo,
	particion.Start,
)

if err != nil {

	fmt.Println(
		"ERROR leyendo SuperBlock",
	)

	return
}

// ------------------------------------------------------------
// Leer users.txt real
// ------------------------------------------------------------

contenido, err := ObtenerContenidoUsersTXT(
	archivo,
	sb,
)

if err != nil {

	fmt.Println(
		"ERROR leyendo users.txt",
	)

	return
}

fmt.Println()
fmt.Println("  USERS.TXT  ")

fmt.Println(
	contenido,
)

//**--
fmt.Println()
fmt.Println(
    "Bytes usados:",
    len(contenido),
)
//**--

// ------------------------------------------------------------
// Validar si el grupo ya existe
// ------------------------------------------------------------

if ExisteGrupo(
	contenido,
	name,
) {

	fmt.Println(
		"ERROR: grupo ya existe",
	)

	return
}

// ------------------------------------------------------------
// Obtener siguiente ID
// ------------------------------------------------------------

id := ObtenerSiguienteIDGrupo(
	contenido,
)

fmt.Println(
	"Siguiente ID:",
	id,
)

// ------------------------------------------------------------
// Construir nuevo registro de grupo
// ------------------------------------------------------------

nuevoRegistro := fmt.Sprintf(
	"%d,G,%s\n",
	id,
	name,
)

// ------------------------------------------------------------
// Agregar al contenido existente
// ------------------------------------------------------------

contenido += nuevoRegistro

// ------------------------------------------------------------
// Guardar users.txt actualizado
// ------------------------------------------------------------

err = GuardarUsersTXT(
	archivo,
	sb,
	contenido,
)

if err != nil {

	fmt.Println(
		"ERROR guardando users.txt",
	)

	return
}

fmt.Println()
fmt.Println("Grupo creado correctamente")

}

// ExisteGrupo: Verifica si un grupo ya existe dentro del contenido de users.txt.
// Formato esperado:
// 1,G,root
// 2,G,developers
//
// Retorna: true  -> grupo ya existe o false -> grupo disponible

func ExisteGrupo(
	contenido string,
	nombre string,
) bool {

	lineas := strings.Split(
		contenido,
		"\n",
	)

	for _, linea := range lineas {

		linea = strings.TrimSpace(
			linea,
		)

		if linea == "" {
			continue
		}

		campos := strings.Split(
			linea,
			",",
		)

		// Grupo = ID,G,NOMBRE

		if len(campos) != 3 {
			continue
		}

		if campos[1] != "G" {
			continue
		}

		id, err := strconv.Atoi(
			campos[0],
		)

		if err != nil {
			continue
		}

		// Ignorar grupos eliminados
		if id == 0 {
			continue
		}

		if strings.EqualFold(
			campos[2],
			nombre,
		) {
			return true
		}


	}

	return false
}

// ObtenerSiguienteIDGrupo: Recorre users.txt y devuelve el siguiente ID disponible para crear un nuevo grupo.

func ObtenerSiguienteIDGrupo(
	contenido string,
) int {

	maxID := 0

	lineas := strings.Split(
		contenido,
		"\n",
	)

	for _, linea := range lineas {

		linea = strings.TrimSpace(
			linea,
		)

		if linea == "" {
			continue
		}

		campos := strings.Split(
			linea,
			",",
		)

		if len(campos) != 3 {
			continue
		}

		if campos[1] != "G" {
			continue
		}

		id, err := strconv.Atoi(
			campos[0],
		)

		if err != nil {
			continue
		}

		if id > maxID {
			maxID = id
		}
	}

	return maxID + 1
}