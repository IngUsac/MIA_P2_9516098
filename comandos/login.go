package comandos

import (
	"fmt"
	"os"
	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)

// LOGIN:  Verifica que la partición indicada por el ID exista dentro de las particiones montadas.

func LOGIN(
	parametros map[string]string,
) {

	user := parametros["user"]
	pass := parametros["pass"]
	id := parametros["id"]

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

	//**--
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

}

		
		
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
