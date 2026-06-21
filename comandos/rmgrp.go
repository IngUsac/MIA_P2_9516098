package comandos

import (
    "fmt"
    "os"
    "strconv"
    "strings"

    "MIA_P1_9516098/estructuras"
)

// RMGRP
// ------------------------------------------------------------
// Elimina lógicamente un grupo de users.txt.
//
// La eliminación lógica consiste en colocar
// el ID del grupo en 0.
//
// Ejemplo:
//
// Antes:
// 2,G,developers
//
// Después:
// 0,G,developers
//
func RMGRP(
	parametros map[string]string,
) {

	fmt.Println(" RMGRP ")
	fmt.Println()
	
	name := parametros["name"]

	// Debe existir sesión activa

	if !estructuras.SesionActual.Activa {

		fmt.Println(
			"ERROR: no existe una sesion activa",
		)

		return
	}

	// Solo root puede ejecutar RMGRP

	if estructuras.SesionActual.User != "root" {

		fmt.Println(
			"ERROR: solo root puede ejecutar RMGRP",
		)

		return
	}

	// Validar parámetro

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

	// ------------------------------------------------------------
	// Obtener partición de la sesión actual
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

	// ------------------------------------------------------------
	// Validar existencia del grupo
	// ------------------------------------------------------------

	if !ExisteGrupoActivo(
		contenido,
		name,
	) {

		fmt.Println(
			"ERROR: grupo no existe",
		)

		return
	}

	fmt.Println(
		"Grupo encontrado",
	)

	// ------------------------------------------------------------
	// Eliminar grupo
	// ------------------------------------------------------------

	contenido = EliminarGrupo(
		contenido,
		name,
	)

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
	fmt.Println(
		"Grupo eliminado correctamente",
	)



}


// ExisteGrupoActivo:  Verifica que un grupo exista y que no esté eliminado.
// Grupo activo:     2,G,developers
// Grupo eliminado:  0,G,developers
// Retorna:  true  -> existe y está activo  o   false -> no existe o está eliminado

func ExisteGrupoActivo(
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

		if len(campos) != 3 {
			continue
		}

		if campos[1] != "G" {
			continue
		}

		if !strings.EqualFold(
			campos[2],
			nombre,
		) {
			continue
		}

		id, err := strconv.Atoi(
			campos[0],
		)

		if err != nil {
			return false
		}

		return id != 0
	}

	return false
}

// EliminarGrupo: Marca un grupo como eliminado colocando su ID en 0.
// Antes:    	2,G,developers
// Después:		0,G,developers
// Retorna el contenido actualizado.

func EliminarGrupo(
	contenido string,
	nombre string,
) string {

	lineas := strings.Split(
		contenido,
		"\n",
	)

	for i, linea := range lineas {

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

		if !strings.EqualFold(
			campos[2],
			nombre,
		) {
			continue
		}

		// Eliminación lógica
		campos[0] = "0"

		lineas[i] = strings.Join(
			campos,
			",",
		)

		break
	}

	return strings.Join(
		lineas,
		"\n",
	)
}