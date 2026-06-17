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

	fmt.Println()
	fmt.Println("===== MKFS =====")

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


	if err != nil {

		fmt.Println(
			"ERROR inicializando bitmap de inodos",
		)

		return
	}

	fmt.Println(
		"Bitmap de inodos inicializado",
	)

	if err != nil {

		fmt.Println(
			"ERROR escribiendo SuperBlock",
		)

		return
	}

	fmt.Println()
	fmt.Println(
		"SuperBlock escrito correctamente",
	)


	fmt.Println()
	fmt.Println("===== SUPERBLOCK =====")

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
}



//**--

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


//**--


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

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FileBlock{},
		),
	)

	n := (sizeParticion - superSize) /
		(4 + inodeSize + (3 * blockSize))

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

	sb.SFreeInodesCount = n
	sb.SFreeBlocksCount = 3 * n

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

	sb.SFirstIno = 0
	sb.SFirstBlo = 0

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

	inode.IType = '0' // carpeta

	inode.IPerm = 664

	return inode
}











// Gguarda el SuperBlock al inicio de la partición

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