package comandos

import (
	"fmt"
	"os"
	"strings"

	"MIA_P1_9516098/estructuras"
)

// RMUSR: Elimina lógicamente un usuario registrado en users.txt. Solo puede ejecutarlo el usuario root. La eliminación lógica consiste en cambiar el UID a 0.
 
func RMUSR(
	parametros map[string]string,
) {
	fmt.Println(" RMUSR ")
	fmt.Println()
	user := parametros["user"]

	if !estructuras.SesionActual.Activa {

		fmt.Println(
			"ERROR: no existe una sesion activa",
		)

		return
	}

	if estructuras.SesionActual.User != "root" {

		fmt.Println(
			"ERROR: solo root puede ejecutar rmusr",
		)

		return
	}

	if user == "" {

		fmt.Println(
			"ERROR: falta parametro user",
		)

		return
	}

	if strings.EqualFold(
		user,
		"root",
	) {

		fmt.Println(
			"ERROR: no se puede eliminar el usuario root",
		)

		return
	}

	if len(user) > 10 {

		fmt.Println(
			"ERROR: user excede 10 caracteres",
		)

		return
	}


	fmt.Println(
		"User:",
		user,
	)

	// Buscar la partición donde está la sesión activa.

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
		"Usuario encontrado",
	)

// Fin Buscar partición de sesión activa.


// EliminarUsuario: Marca un usuario como eliminado dentro de users.txt. La eliminación lógica consiste en cambiar el UID a 0.
// Parámetros:
// contenido -> contenido completo de users.txt
// user      -> usuario a eliminar
// Retorna: contenido actualizado.

	contenido = EliminarUsuario(
		contenido,
		user,
	)

	fmt.Println()
	fmt.Println("  USERS.TXT ACTUALIZADO  ")

	fmt.Println(
		contenido,
	)

	// Guardar cambios.

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
		"Usuario eliminado correctamente",
	)


}



// EliminarUsuario: Realiza la eliminación lógica de un usuario dentro de users.txt cambiando su UID a 0.
// Parámetros:
// contenido -> contenido completo de users.txt
// user      -> usuario a eliminar

func EliminarUsuario(
	contenido string,
	user string,
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

		if len(campos) != 5 {
			continue
		}

		if campos[1] != "U" {
			continue
		}

		if strings.EqualFold(
			campos[3],
			user,
		) {

			campos[0] = "0"

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