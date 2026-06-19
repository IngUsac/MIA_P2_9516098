package comandos

import (
	"fmt"
	"os"
	"MIA_P1_9516098/estructuras"
//	"MIA_P1_9516098/utilidades"
//	"strconv"
//	"strings"
)

// LOGIN:  Verifica que la partición indicada por el ID exista dentro de las particiones montadas.

func LOGIN(
	parametros map[string]string,
) {

	user := parametros["user"]
	pass := parametros["pass"]
	id := parametros["id"]

	if estructuras.SesionActual.Activa {

		fmt.Println(
			"ERROR: ya existe una sesion activa",
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

	if id == "" {

		fmt.Println(
			"ERROR: falta parametro id",
		)

		return
	}

	// Buscar partición montada
	particion, existe := BuscarParticionMontadaPorID(
		id,
	)

	if !existe {

		fmt.Println(
			"ERROR: particion no montada",
		)

		return
	}


	fmt.Println(
		"User:",
		user,
	)

	fmt.Println(
		"Pass:",
		pass,
	)

	fmt.Println(
		"ID:",
		id,
	)

	fmt.Println(
		"Path:",
		particion.Path,
	)

	fmt.Println(
		"Start:",
		particion.Start,
	)

	// Abrir disco
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

	// Leer SuperBlock
	sb, err := LeerSuperBlock(
		archivo,
		particion.Start,
	)

	contenido, err := ObtenerContenidoUsersTXT(
		archivo,
		sb,
	)

	usuario, encontrado := BuscarUsuario(
		contenido,
		user,
	)

	if !encontrado {

		fmt.Println(
			"ERROR: usuario no encontrado",
		)

		return
	}

	if usuario.Password != pass {

		fmt.Println(
			"ERROR: contraseña incorrecta",
		)

		return
	}

	
	IniciarSesion(
		usuario,
		id,
		particion.Path,
		particion.Start,
	)

	fmt.Println()
	fmt.Println("===== LOGIN EXITOSO =====")

	//**-------
err = MKDIRInterno(
	archivo,
	&sb,
	0,
	"home",
)

if err != nil {

	fmt.Println(
		"ERROR MKDIR:",
		err,
	)

} else {

	fmt.Println(
		"Directorio home creado",
	)
}

inodeHome,
numHome,
err := ObtenerInodoPorRutaCompleta(
	archivo,
	sb,
	"/home",
)

if err != nil {

	fmt.Println(
		"ERROR buscando home:",
		err,
	)

} else {

	fmt.Println()
	fmt.Println("HOME ENCONTRADO")
	fmt.Println("Inodo:", numHome)
	fmt.Println("Tipo:", string(inodeHome.IType))
}


fmt.Println()
fmt.Println("===== PRUEBA MKDIR =====------------------------------")

err = MKDIRInterno(
	archivo,
	&sb,
	0,
	"home",
)

if err != nil {

	fmt.Println(
		"ERROR MKDIR:------------------------------",
		err,
	)

} else {

	fmt.Println(
		"Directorio home creado------------------------------",
	)
}


fmt.Println()
fmt.Println("===== VALIDANDO /home =====--------------------------")

inodeHome,
	numHome,
	err = ObtenerInodoPorRutaCompleta(
	archivo,
	sb,
	"/home",
)

if err != nil {

	fmt.Println(
		"ERROR buscando home:-----------",
		err,
	)

} else {

	fmt.Println(
		"HOME ENCONTRADO-----------",
	)

	fmt.Println(
		"Inodo:-----------",
		numHome,
	)

	fmt.Println(
		"Tipo:-----------",
		string(inodeHome.IType),
	)
}
	//**-------
	
	fmt.Println(
		"Usuario:",
		usuario.User,
	)

	fmt.Println(
		"UID:",
		usuario.UID,
	)

	fmt.Println(
		"GID:",
		1,
	)

	fmt.Println(
		"ID:",
		id,
	)
	

	if !encontrado {

		fmt.Println(
			"ERROR: usuario no encontrado",
		)

		return
	}


	if err != nil {

		fmt.Println(
			"ERROR leyendo users.txt",
		)

		return
	}

		
	if err != nil {

		fmt.Println(
			"ERROR: no se pudo leer el SuperBlock",
		)

		return
	}

}  // fin LOGIN ()

		


// IniciarSesion: crea la sesión activa del sistema.

func IniciarSesion(
	usuario estructuras.Usuario,
	id string,
	path string,
	start int32,
) {

	estructuras.SesionActual = estructuras.Sesion{
		Activa: true,
		User:   usuario.User,
		Pass:   usuario.Password,
		ID:     id,
		UID:    usuario.UID,
		GID:    1,
		Path:  path,
		Start: start,
	}
}
