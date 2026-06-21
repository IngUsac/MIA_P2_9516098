package comandos

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"MIA_P1_9516098/estructuras"
)

func MKUSR(
	parametros map[string]string,
) {
	fmt.Println(" MKUSR ")
	fmt.Println()

	user := parametros["user"]
	pass := parametros["pass"]
	grp := parametros["grp"]

	if !estructuras.SesionActual.Activa {

		fmt.Println(
			"ERROR: no existe una sesion activa",
		)

		return
	}

	if estructuras.SesionActual.User != "root" {

		fmt.Println(
			"ERROR: solo root puede ejecutar mkusr",
		)

		return
	}

	if user == "" {

		fmt.Println(
			"ERROR: falta parametro user",
		)

		return
	}

	if pass == "" {

		fmt.Println(
			"ERROR: falta parametro pass",
		)

		return
	}

	if grp == "" {

		fmt.Println(
			"ERROR: falta parametro grp",
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

	particion, existe := BuscarParticionMontadaPorID(
		estructuras.SesionActual.ID,
	)

	if !existe {

		fmt.Println("ERROR: particion no encontrada",)

		return
	}

	archivo, err := os.OpenFile(
		particion.Path,
		os.O_RDWR,
		0644,
	)

	if err != nil {

		fmt.Println("ERROR abriendo disco",)

		return
	}

	defer archivo.Close()

	sb, err := LeerSuperBlock(
		archivo,
		particion.Start,
	)

	if err != nil {

		fmt.Println(
			"ERROR leyendo superblock",
		)

		return
	}

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
	fmt.Println("===== USERS.TXT =====")

	fmt.Println(
		contenido,
	)

	if !ExisteGrupoActivo(
		contenido,
		grp,
	) {

		fmt.Println(
			"ERROR: grupo no existe",
		)

		return
	}

	if ExisteUsuario(
		contenido,
		user,
	) {

		fmt.Println(
			"ERROR: usuario ya existe",
		)

		return
	}

	fmt.Println(
		"Grupo válido",
	)

	fmt.Println(
		"Usuario disponible",
	)

	

	// Obtener siguiente UID disponible

	nuevoUID := ObtenerSiguienteUID(
		contenido,
	)

	fmt.Println(
		"Siguiente UID:",
		nuevoUID,
	)

	// Crear nuevo registro

	nuevoRegistro := fmt.Sprintf(
		"%d,U,%s,%s,%s\n",
		nuevoUID,
		grp,
		user,
		pass,
	)

	// Agregar al contenido

	contenido += nuevoRegistro

	// Guardar en disco

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
		"Usuario creado correctamente",
	) 
	// Fin obtener sig usuario UID



}

// ExisteUsuario:  Busca un usuario activo dentro de users.txt
// Retorna: true  -> existe  o  false -> no existe

func ExisteUsuario(
	contenido string,
	user string,
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
			return true
		}
	}

	return false
}

// ObtenerSiguienteUID: Busca el UID más alto registrado en users.txt y devuelve el siguiente disponible.

func ObtenerSiguienteUID(
	contenido string,
) int {

	lineas := strings.Split(
		contenido,
		"\n",
	)

	ultimoUID := 0

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

		if len(campos) != 5 {
			continue
		}

		if campos[1] != "U" {
			continue
		}

		uid, err := strconv.Atoi(
			campos[0],
		)

		if err != nil {
			continue
		}

		if uid > ultimoUID {
			ultimoUID = uid
		}
	}

	return ultimoUID + 1
}

// ObtenerSiguienteUID: Busca el UID más alto de los usuarios activos y devuelve el siguiente disponible.
