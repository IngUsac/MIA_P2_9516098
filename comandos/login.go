package comandos

import (
	"fmt"
	"os"
	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
	"strconv"
	"strings"
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

	fmt.Println()
	fmt.Println("===== LOGIN =====")

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
	archivo, err := os.Open(
		particion.Path,
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
	//**--
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

	//****----
	IniciarSesion(
		usuario,
		id,
	)

	fmt.Println()
	fmt.Println("===== LOGIN EXITOSO =====")

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
	//****----

	if !encontrado {

		fmt.Println(
			"ERROR: usuario no encontrado",
		)

		return
	}

	fmt.Println()
	fmt.Println("===== USUARIO ENCONTRADO =====")

	fmt.Println(
		"UID:",
		usuario.UID,
	)

	fmt.Println(
		"Grupo:",
		usuario.Grupo,
	)

	fmt.Println(
		"User:",
		usuario.User,
	)

	fmt.Println(
		"Password:",
		usuario.Password,
	)
	//**--


	if err != nil {

		fmt.Println(
			"ERROR leyendo users.txt",
		)

		return
	}

	fmt.Println()
	fmt.Println("===== USERS.TXT =====")
	fmt.Println(contenido)


	if err != nil {

		fmt.Println(
			"ERROR: no se pudo leer el SuperBlock",
		)

		return
	}

	fmt.Println()
	fmt.Println("===== SUPERBLOCK =====")

	fmt.Println(
		"Magic:",
		sb.SMagic,
	)

	fmt.Println(
		"Inodes:",
		sb.SInodesCount,
	)

	fmt.Println(
		"Blocks:",
		sb.SBlocksCount,
	)

	fmt.Println(
		"Free Inodes:",
		sb.SFreeInodesCount,
	)

	fmt.Println(
		"Free Blocks:",
		sb.SFreeBlocksCount,
	)

	fmt.Println(
		"First Inode:",
		sb.SFirstIno,
	)

	fmt.Println(
		"First Block:",
		sb.SFirstBlo,
	)


	inodeRaiz, err := LeerInodo(
		archivo,
		sb.SInodeStart,
	)

	if err != nil {

		fmt.Println(
			"ERROR: no se pudo leer el inodo raiz",
		)

		return
	}

	fmt.Println()
	fmt.Println("===== INODO RAIZ =====")

	fmt.Println(
		"UID:",
		inodeRaiz.IUid,
	)

	fmt.Println(
		"GID:",
		inodeRaiz.IGid,
	)

	fmt.Println(
		"Tipo:",
			string(inodeRaiz.IType),
	)

	fmt.Println(
		"Bloque[0]:",
		inodeRaiz.IBlock[0],
	)

	folderRaiz, err := LeerFolderBlock(
		archivo,
		sb.SBlockStart,
	)

	if err != nil {

		fmt.Println(
			"ERROR: no se pudo leer folder raiz",
		)

		return
	}

	fmt.Println()
	fmt.Println("===== FOLDER RAIZ =====")

	for i := 0; i < 4; i++ {

		fmt.Println(
			"Nombre:",
			utilidades.BytesAString(
				folderRaiz.BContent[i].BName[:],
			),
			"Inodo:",
			folderRaiz.BContent[i].BInodo,
		)
	}

	inodeSize := int32(
		utilidades.ObtenerTamano(
			estructuras.Inode{},
		),
	)

	// users.txt está en el inodo 1
	posUsersInode := sb.SInodeStart + inodeSize

	inodeUsers, err := LeerInodo(
		archivo,
		posUsersInode,
	)

	if err != nil {

		fmt.Println(
			"ERROR: no se pudo leer inodo users.txt",
		)

		return
	}

	fmt.Println()
	fmt.Println("===== INODO USERS.TXT =====")

	fmt.Println(
		"UID:",
		inodeUsers.IUid,
	)

	fmt.Println(
		"GID:",
		inodeUsers.IGid,
	)

	fmt.Println(
		"Tipo:",
		string(inodeUsers.IType),
	)

	fmt.Println(
		"Bloque[0]:",
		inodeUsers.IBlock[0],
	)

	
	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FileBlock{},
		),
	)

	posUsersBlock := sb.SBlockStart + blockSize

	usersFile, err := LeerFileBlock(
		archivo,
		posUsersBlock,
	)

	if err != nil {

		fmt.Println(
			"ERROR: no se pudo leer users.txt",
		)

		return
	}

	fmt.Println()
	fmt.Println("===== USERS.TXT =====")

	fmt.Println(
		utilidades.BytesAString(
			usersFile.BContent[:],
		),
	)
	//**--
    
	//**--

}  // fin LOGIN ()

		
		
// LeerInodo:  Lee un inodo desde una posición específica del disco.
// Parámetros: 
// archivo  -> disco abierto
// posicion -> byte donde inicia el inodo
//
// Retorna: Inodo leído o  error en caso de fallo
//
func LeerInodo(
	archivo *os.File,
	posicion int32,
) (estructuras.Inode, error) {

	var inode estructuras.Inode

	err := utilidades.LeerObjeto(
		archivo,
		&inode,
		int64(posicion),
	)

	if err != nil {
		return estructuras.Inode{}, err
	}

	return inode, nil
}

	
// LeerFolderBlock: Lee un bloque de carpeta desde disco.

func LeerFolderBlock(
	archivo *os.File,
	posicion int32,
) (estructuras.FolderBlock, error) {

	var folder estructuras.FolderBlock

	err := utilidades.LeerObjeto(
		archivo,
		&folder,
		int64(posicion),
	)

	if err != nil {
		return estructuras.FolderBlock{}, err
	}

	return folder, nil
}

// LeerFileBlock: Lee un bloque de archivo desde disco.
//
func LeerFileBlock(
	archivo *os.File,
	posicion int32,
) (estructuras.FileBlock, error) {

	var file estructuras.FileBlock

	err := utilidades.LeerObjeto(
		archivo,
		&file,
		int64(posicion),
	)

	if err != nil {
		return estructuras.FileBlock{}, err
	}

	return file, nil
}

// ObtenerContenidoUsersTXT: Recorre las estructuras EXT2 y devuelve el contenido completo del archivo users.txt

func ObtenerContenidoUsersTXT(
	archivo *os.File,
	sb estructuras.SuperBlock,
) (string, error) {

	inodeSize := int32(
		utilidades.ObtenerTamano(
			estructuras.Inode{},
		),
	)

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FileBlock{},
		),
	)

	// users.txt está en el inodo 1
	posUsersInode := sb.SInodeStart + inodeSize

	inodeUsers, err := LeerInodo(
		archivo,
		posUsersInode,
	)

	if err != nil {
		return "", err
	}

	// bloque 1
	posUsersBlock := sb.SBlockStart + blockSize

	usersFile, err := LeerFileBlock(
		archivo,
		posUsersBlock,
	)

	if err != nil {
		return "", err
	}

	_ = inodeUsers

	return utilidades.BytesAString(
		usersFile.BContent[:],
	), nil
}

// BuscarUsuario:  Busca un usuario dentro del contenido de users.txt.

func BuscarUsuario(
	contenido string,
	user string,
) (estructuras.Usuario, bool) {

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

		// Registro de usuario:
		// UID,U,GRUPO,USER,PASS

		if len(campos) != 5 {
			continue
		}

		if campos[1] != "U" {
			continue
		}

		if !strings.EqualFold(
			campos[3],
			user,
		) {
			continue
		}

		uid, _ := strconv.Atoi(
			campos[0],
		)

		return estructuras.Usuario{
			UID:      int32(uid),
			Grupo:    campos[2],
			User:     campos[3],
			Password: campos[4],
		}, true
	}

	return estructuras.Usuario{}, false
}

// IniciarSesion: crea la sesión activa del sistema.

func IniciarSesion(
	usuario estructuras.Usuario,
	id string,
) {

	estructuras.SesionActual = estructuras.Sesion{
		Activa: true,
		User:   usuario.User,
		Pass:   usuario.Password,
		ID:     id,
		UID:    usuario.UID,
		GID:    1,
	}
}