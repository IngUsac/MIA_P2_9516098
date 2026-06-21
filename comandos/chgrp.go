package comandos

import (
	"fmt"
	"os"
	"strings"

	"MIA_P1_9516098/estructuras"
)

// CHGRP: Cambia el grupo asignado a un usuario activo. Solo puede ejecutarlo el usuario root.
// Requisitos:
// - Debe existir una sesión activa.
// - El usuario debe existir.
// - El grupo destino debe existir.

func CHGRP(
	parametros map[string]string,
) {
	fmt.Println("  CHGRP ") 
	fmt.Println()
	
	user := parametros["user"]
	grp := parametros["grp"]

	if !estructuras.SesionActual.Activa {

		fmt.Println(
			"ERROR: no existe una sesion activa",
		)

		return
	}

	if estructuras.SesionActual.User != "root" {

		fmt.Println(
			"ERROR: solo root puede ejecutar chgrp",
		)

		return
	}

	if user == "" {

		fmt.Println(
			"ERROR: falta parametro user",
		)

		return
	}

	if grp == "" {

		fmt.Println(
			"ERROR: falta parametro grp",
		)

		return
	}


	if len(user) > 10 {

		fmt.Println(
			"ERROR: user excede 10 caracteres",
		)

		return
	}

	if len(grp) > 10 {

		fmt.Println(
			"ERROR: grp excede 10 caracteres",
		)

		return
	}



	fmt.Println(
		"User:",
		user,
	)

	fmt.Println(
		"Grupo:",
		grp,
	)

	// Buscar la partición de la sesión activa.

	particion, existe := BuscarParticionMontadaPorID(
		estructuras.SesionActual.ID,
	)

	if !existe {

		fmt.Println(
			"ERROR: particion no encontrada",
		)

		return
	}

	// Abrir disco.

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

	// Leer SuperBlock.

	sb, err := LeerSuperBlock(
		archivo,
		particion.Start,
	)

	if err != nil {

		fmt.Println(
			"ERROR: no se pudo leer el SuperBlock",
		)

		return
	}

	// Leer users.txt.

	contenido, err := ObtenerContenidoUsersTXT(
		archivo,
		sb,
	)

	if err != nil {

		fmt.Println(
			"ERROR: no se pudo leer users.txt",
		)

		return
	}

	fmt.Println()
	fmt.Println("  USERS.TXT  ")

	fmt.Println(
		contenido,
	)

	// Verificar grupo destino.

	if !ExisteGrupoActivo(
		contenido,
		grp,
	) {

		fmt.Println(
			"ERROR: grupo no existe",
		)

		return
	}

	// Verificar usuario.

	if !ExisteUsuarioActivo(
		contenido,
		user,
	) {

		fmt.Println(
			"ERROR: usuario no existe",
		)

		return
	}

	fmt.Println(
		"Grupo encontrado",
	)

	fmt.Println(
		"Usuario encontrado",
	)

  // Fin Buscar la partición de la sesión activa.

	contenido = CambiarGrupoUsuario(
		contenido,
		user,
		grp,
	)

	fmt.Println()
	fmt.Println(
		"  USERS.TXT actualizado en CHGRP  ",
	)

	fmt.Println(
		contenido,
	)

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
		"Grupo actualizado correctamente",
	)

 }

// CambiarGrupoUsuario: Modifica el grupo asociado a un usuario activo.
// Parámetros:
// contenido   -> contenido completo de users.txt
// user        -> usuario a modificar
// grupoNuevo  -> nuevo grupo
// Retorna: contenido actualizado.

func CambiarGrupoUsuario(
	contenido string,
	user string,
	grupoNuevo string,
) string {

	lineas := strings.Split(
		contenido,
		"\n",
	)

	for i, linea := range lineas {

		campos := strings.Split(
			linea,
			",",
		)

		if len(campos) != 5 {
			continue
		}

		if campos[1] != "U" {
			continue
		}

		if campos[0] == "0" {
			continue
		}

		if strings.EqualFold(
			campos[3],
			user,
		) {

			campos[2] = grupoNuevo

			lineas[i] = strings.Join(
				campos,
				",",
			)

			break
		}
	}

	return strings.Join(
		lineas,
		"\n",
	)
}
