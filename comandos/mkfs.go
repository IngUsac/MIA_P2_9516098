package comandos

import (
	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
	"fmt"
	"os"
)

func MKFS(
	parametros map[string]string,
) {

	fmt.Println(" MKSF, parametros", parametros)
	fmt.Println()
	id := parametros["id"]

	if id == "" {

		fmt.Println(
			"ERROR: falta parametro id",
		)

		return
	}

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
		"ID:",
		particion.ID,
	)

	fmt.Println(
		"Path:",
		particion.Path,
	)

	fmt.Println(
		"Name:",
		particion.Name,
	)

	fmt.Println(
		"Start:",
		particion.Start,
	)

	fmt.Println(
		"Size:",
		particion.Size,
	)

	fmt.Println(
		"Type:",
		string(particion.Type),
	)
	
	n := CalcularNumeroInodos(
		particion.Size,
	)

	sb := CrearSuperBlock(
		particion.Start,
		n,
	)

	archivo, err := AbrirDisco(
		particion.Path,
	)

	if err != nil {

		fmt.Println(
			"ERROR abriendo disco",
		)

		return
	}

	defer archivo.Close()

	err = EscribirSuperBlock(
		archivo,
		sb,
		particion.Start,
	)

	err = InicializarBitmapInodos(
		archivo,
		sb.SBmInodeStart,
		n,
	)

	err = InicializarBitmapBloques(
		archivo,
		sb.SBmBlockStart,
		n,
	)

	if err != nil {

		fmt.Println(
			"ERROR inicializando bitmap de bloques",
		)

		return
	}

	fmt.Println(
		"Bitmap de bloques inicializado",
	)

	inodoRaiz := CrearInodoRaiz()

	err = EscribirInodo(
		archivo,
		inodoRaiz,
		sb.SInodeStart,
	)

	if err != nil {

		fmt.Println(
			"ERROR escribiendo inodo raiz",
		)

		return
	}

	fmt.Println(
		"Inodo raiz creado",
	)

	
	folderRaiz := CrearFolderRaiz()

	err = EscribirFolderBlock(
		archivo,
		folderRaiz,
		sb.SBlockStart,
	)

	if err != nil {

		fmt.Println(
			"ERROR escribiendo folder raiz",
		)

		return
	}

	fmt.Println(
		"Folder raiz creado",
	)
	

	inodoUsers := CrearInodoUsers()

	inodeSize := int32(
		utilidades.ObtenerTamano(
			estructuras.Inode{},
		),
	)

	err = EscribirInodo(
		archivo,
		inodoUsers,
		sb.SInodeStart+inodeSize,
	)

	if err != nil {

		fmt.Println(
			"ERROR escribiendo inodo users.txt",
		)

		return
	}

	fmt.Println(
		"Inodo users.txt creado",
	)

	usersFile := CrearUsersFile()


	err = GuardarFileBlock(
		archivo,
		sb,
		1,
		usersFile,
	)
		
		
	if err != nil {
		fmt.Println(
			"ERROR escribiendo users.txt",
		)

		return
	}


	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FileBlock{},
		),
	)

	var bloqueVacio estructuras.FileBlock

	for i := int32(2); i <= 4; i++ {

		posBloque := sb.SBlockStart +
			(i * blockSize) //

		err = EscribirFileBlock(
			archivo,
			bloqueVacio,
			posBloque,
		)

		if err != nil {

			fmt.Println(
				"ERROR inicializando bloque users.txt",
			)

			return
		}
	}
//


	err = EscribirFileBlock(
		archivo,
		usersFile,
		sb.SBlockStart+blockSize, //
	)

	if err != nil {

		fmt.Println(
			"ERROR escribiendo users.txt",
		)

		return
	}

	fmt.Println(
		"users.txt creado",
	)

	
	// Inodo raíz

	err = MarcarBitmapInodo(
		archivo,
		sb.SBmInodeStart,
		0,
	)

	if err != nil {
		fmt.Println("ERROR bitmap inode 0")
		return
	}

	// Inodo users.txt
	err = MarcarBitmapInodo(
		archivo,
		sb.SBmInodeStart,
		1,
	)

	if err != nil {
		fmt.Println("ERROR bitmap inode 1")
		return
	}

	// Bloque raíz
	err = MarcarBitmapBloque(
		archivo,
		sb.SBmBlockStart,
		0,
	)

	if err != nil {
		fmt.Println("ERROR bitmap block 0")
		return
	}

	// Bloque users.txt 
/*
	for i := int32(1); i <= 4; i++ {

		err = MarcarBitmapBloque(
			archivo,
			sb.SBmBlockStart,
			i,
		)

		if err != nil {
			return
		}
	}*/

	err = MarcarBitmapBloque(
        archivo,
        sb.SBmBlockStart,
        1,
)

	if err != nil {
		fmt.Println("ERROR bitmap block 1")
		return
	}

	fmt.Println(
		"Bitmaps actualizados",
	)



	b0, _ := LeerByte(
		archivo,
		sb.SBmInodeStart,
	)

	b1, _ := LeerByte(
		archivo,
		sb.SBmInodeStart+1,
	)

	fmt.Println()
	fmt.Println("  BITMAP INODOS  ")
	fmt.Println("0:", b0)
	fmt.Println("1:", b1)

	bb0, _ := LeerByte(
		archivo,
		sb.SBmBlockStart,
	)

	bb1, _ := LeerByte(
		archivo,
		sb.SBmBlockStart+1,
	)

	fmt.Println()
	fmt.Println("  BITMAP BLOQUES  ")
	fmt.Println("0:", bb0)
	fmt.Println("1:", bb1)


	fmt.Println(
		"Bitmap de inodos inicializado",
	)


	fmt.Println()
	fmt.Println(
		"SuperBlock escrito correctamente",
	)


	fmt.Println()
	fmt.Println("  SUPERBLOCK  ")

	fmt.Println(
		"Inodes:",
		sb.SInodesCount,
	)

	fmt.Println(
		"Blocks:",
		sb.SBlocksCount,
	)

	fmt.Println(
		"BM Inodos:",
		sb.SBmInodeStart,
	)

	fmt.Println(
		"BM Bloques:",
		sb.SBmBlockStart,
	)

	fmt.Println(
		"Inodos:",
		sb.SInodeStart,
	)

	fmt.Println(
		"Bloques:",
		sb.SBlockStart,
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

}

func LeerSuperBlock(
	archivo *os.File,
	inicio int32,
) (estructuras.SuperBlock, error) {

	var sb estructuras.SuperBlock

	err := utilidades.LeerObjeto(
		archivo,
		&sb,
		int64(inicio),
	)

	return sb, err
}





// CalcularNumeroInodos calcula n para EXT2.

func CalcularNumeroInodos(
	sizeParticion int32,
) int32 {

	superSize := int32(
		utilidades.ObtenerTamano(
			estructuras.SuperBlock{},
		),
	)

	inodeSize := int32(
		utilidades.ObtenerTamano(
			estructuras.Inode{},
		),
	)
//
	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FileBlock{},
		),
	)

	n := (sizeParticion - superSize) /
		(4 + inodeSize + (3 * blockSize))//

	return n
}

// CrearSuperBlock construye la estructura base EXT2.

func CrearSuperBlock(
	inicioParticion int32,
	n int32,
) estructuras.SuperBlock {

	var sb estructuras.SuperBlock

	sb.SFilesystemType = 2

	sb.SInodesCount = n 
	sb.SBlocksCount = 3 * n

	sb.SFreeInodesCount = n -2
	sb.SFreeBlocksCount = (3 * n) - 2

	sb.SMagic = 0xEF53

	sb.SInodeSize = int32(
		utilidades.ObtenerTamano(
			estructuras.Inode{},
		),
	)

	sb.SBlockSize = int32(
		utilidades.ObtenerTamano(
			estructuras.FileBlock{},
		),
	)

	sb.SFirstIno = 2
	sb.SFirstBlo = 2

	superSize := int32(
		utilidades.ObtenerTamano(
			estructuras.SuperBlock{},
		),
	)

	sb.SBmInodeStart =
		inicioParticion + superSize

	sb.SBmBlockStart =
		sb.SBmInodeStart + n

	sb.SInodeStart =
		sb.SBmBlockStart + (3 * n)

	sb.SBlockStart =
		sb.SInodeStart +
			(n * sb.SInodeSize)

	return sb
}

// Crea el inodo 0 correspondiente al directorio raíz "/"

func CrearInodoRaiz() estructuras.Inode {

	var inode estructuras.Inode

	inode.IUid = 1
	inode.IGid = 1

	inode.ISize = 0

	for i := 0; i < 15; i++ {
		inode.IBlock[i] = -1
	}

	inode.IBlock[0] = 0

	for i := 1; i < 15; i++ {
			inode.IBlock[i] = -1
	}

	inode.IType = '0' // carpeta

	inode.IPerm = 664

	return inode
}


// CrearInodoUsers: crea el inodo inicial del archivo users.txt.

func CrearInodoUsers() estructuras.Inode {

	var inode estructuras.Inode

	inode.IUid = 1
	inode.IGid = 1

	for i := 0; i < 15; i++ {
		inode.IBlock[i] = -1
	}

	inode.IBlock[0] = 1

	contenido :=
		"1,G,root\n1,U,root,root,123\n"

	inode.ISize = int32(
		len(contenido),
	)

	inode.IType = '1'
	inode.IPerm = 664

	return inode
}


// CrearFolderRaiz crea el bloque de carpeta "/".
func CrearFolderRaiz() estructuras.FolderBlock {

	var folder estructuras.FolderBlock

	copy(
		folder.BContent[0].BName[:],
		".",
	)

	folder.BContent[0].BInodo = 0

	copy(
		folder.BContent[1].BName[:],
		"..",
	)

	folder.BContent[1].BInodo = 0

	copy(
		folder.BContent[2].BName[:],
		"users.txt",
	)

	folder.BContent[2].BInodo = 1

	folder.BContent[3].BInodo = -1

	return folder
}


// Crear el contenido inicial de users.txt  

func CrearUsersFile() estructuras.FileBlock {

	var file estructuras.FileBlock

	contenido := "1,G,root\n1,U,root,root,123\n"

	copy(
		file.BContent[:],
		contenido,
	)

	return file
}







// Guarda el SuperBlock al inicio de la partición

func EscribirSuperBlock(
	archivo *os.File,
	sb estructuras.SuperBlock,
	inicio int32,
) error {

	return utilidades.EscribirObjeto(
		archivo,
		&sb,
		int64(inicio),
	)
}


// Gguarda un inodo en disco.

func EscribirInodo(
	archivo *os.File,
	inode estructuras.Inode,
	posicion int32,
) error {

	return utilidades.EscribirObjeto(
		archivo,
		&inode,
		int64(posicion),
	)
}

// Guarda un FolderBlock en disco.

func EscribirFolderBlock(
	archivo *os.File,
	folder estructuras.FolderBlock,
	posicion int32,
) error {

	return utilidades.EscribirObjeto(
		archivo,
		&folder,
		int64(posicion),
	)
} 


func EscribirFileBlock(
	archivo *os.File,
	file estructuras.FileBlock,
	posicion int32,
) error {

	return utilidades.EscribirObjeto(
		archivo,
		&file,
		int64(posicion),
	)
}


// llena con 0 el bitmap de inodos.

func InicializarBitmapInodos(
	archivo *os.File,
	inicio int32,
	n int32,
) error {

	for i := int32(0); i < n; i++ {

		var cero byte = 0

		err := utilidades.EscribirObjeto(
			archivo,
			&cero,
			int64(inicio+i),
		)

		if err != nil {
			return err
		}
	}

	return nil
}


// llena con 0 el bitmap de bloques.

func InicializarBitmapBloques(
	archivo *os.File,
	inicio int32,
	n int32,
) error {

	for i := int32(0); i < (3 * n); i++ {

		var cero byte = 0

		err := utilidades.EscribirObjeto(
			archivo,
			&cero,
			int64(inicio+i),
		)

		if err != nil {
			return err
		}
	}

	return nil
}

// Marca una posición del bitmap de inodos como ocupada.

func MarcarBitmapInodo(
	archivo *os.File,
	inicio int32,
	indice int32,
) error {

	var ocupado byte = 1

	return utilidades.EscribirObjeto(
		archivo,
		&ocupado,
		int64(inicio+indice),
	)
}

// Marca una posición del bitmap de bloques como ocupada.
func MarcarBitmapBloque(
	archivo *os.File,
	inicio int32,
	indice int32,
) error {

	var ocupado byte = 1

	return utilidades.EscribirObjeto(
		archivo,
		&ocupado,
		int64(inicio+indice),
	)
}

func LeerByte(
	archivo *os.File,
	posicion int32,
) (byte, error) {

	var valor byte

	err := utilidades.LeerObjeto(
		archivo,
		&valor,
		int64(posicion),
	)

	return valor, err
}


// GuardarFilePorNumero: guarda un FileBlock usando número lógico.

func GuardarFilePorNumero(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroBloque int32,
	file estructuras.FileBlock,
) error {

	posicion := ObtenerPosicionBloque(
		sb,
		numeroBloque,
	)

	return EscribirFileBlock(
		archivo,
		file,
		posicion,
	)
}

 // LeerPointerBlock: Lee un bloque de apuntadores desde disco.
// Parámetros:
// archivo  -> disco abierto
// posicion -> posición física del bloque
// Retorna: PointerBlock leído y error si ocurre algún problema.

func LeerPointerBlock(
	archivo *os.File,
	posicion int32,
) (
	estructuras.PointerBlock,
	error,
) {

	var pointer estructuras.PointerBlock

	err := utilidades.LeerObjeto(
		archivo,
		&pointer,
		int64(posicion),
	)

	return pointer, err
}

// EscribirPointerBlock: Guarda un bloque de apuntadores en disco.
// Parámetros:
// archivo  -> disco abierto
// pointer  -> bloque a guardar
// posicion -> posición física destino.

func EscribirPointerBlock(
	archivo *os.File,
	pointer estructuras.PointerBlock,
	posicion int32,
) error {

	return utilidades.EscribirObjeto(
		archivo,
		&pointer,
		int64(posicion),
	)
}

// LeerPointerPorNumero: Lee un PointerBlock utilizando su número lógico.
// Parámetros:
// archivo      -> disco abierto
// sb           -> SuperBlock
// numeroBloque -> bloque lógico
// Retorna: PointerBlock leído.

func LeerPointerPorNumero(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroBloque int32,
) (
	estructuras.PointerBlock,
	error,
) {

	posicion := ObtenerPosicionBloque(
		sb,
		numeroBloque,
	)

	return LeerPointerBlock(
		archivo,
		posicion,
	)
}

// GuardarPointerPorNumero: Guarda un PointerBlock utilizando su número lógico.
// Parámetros:
// archivo      -> disco abierto
// sb           -> SuperBlock
// numeroBloque -> bloque lógico destino
// pointer      -> bloque a guardar.

func GuardarPointerPorNumero(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroBloque int32,
	pointer estructuras.PointerBlock,
) error {

	posicion := ObtenerPosicionBloque(
		sb,
		numeroBloque,
	)

	return EscribirPointerBlock(
		archivo,
		pointer,
		posicion,
	)
}

