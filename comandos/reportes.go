package comandos

import (
	"fmt"
	"os"
	//"strings"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)


// ReporteSB: Lee y muestra la información del SuperBlock.

func ReporteSB(
	particion estructuras.ParticionMontada,
	path string,
) {

	archivo, err := os.Open(
		particion.Path,
	)

	if err != nil {

		fmt.Println(
			"ERROR abriendo disco",
		)

		return
	}

	defer archivo.Close()

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

	fmt.Println()
	fmt.Println("===== SUPERBLOCK =====")

	fmt.Println(
		"SFilesystemType:",
		sb.SFilesystemType,
	)

	fmt.Println(
		"SInodesCount:",
		sb.SInodesCount,
	)

	fmt.Println(
		"SBlocksCount:",
		sb.SBlocksCount,
	)

	fmt.Println(
		"SFreeBlocksCount:",
		sb.SFreeBlocksCount,
	)

	fmt.Println(
		"SFreeInodesCount:",
		sb.SFreeInodesCount,
	)

	fmt.Println(
		"SMntCount:",
		sb.SMntCount,
	)

	fmt.Println(
		"SMagic:",
		sb.SMagic,
	)

	fmt.Println(
		"SInodeSize:",
		sb.SInodeSize,
	)

	fmt.Println(
		"SBlockSize:",
		sb.SBlockSize,
	)

	fmt.Println(
		"SFirstIno:",
		sb.SFirstIno,
	)

	fmt.Println(
		"SFirstBlo:",
		sb.SFirstBlo,
	)

	fmt.Println(
		"SBmInodeStart:",
		sb.SBmInodeStart,
	)

	fmt.Println(
		"SBmBlockStart:",
		sb.SBmBlockStart,
	)

	fmt.Println(
		"SInodeStart:",
		sb.SInodeStart,
	)

	fmt.Println(
		"SBlockStart:",
		sb.SBlockStart,
	)

	fmt.Println("======================")
	fmt.Println()
}

// ReporteBMInode: Muestra el bitmap de inodos ocupados/libres.

func ReporteBMInode(
	particion estructuras.ParticionMontada,
	path string,
) {

	archivo, err := os.Open(
		particion.Path,
	)

	if err != nil {

		fmt.Println(
			"ERROR abriendo disco",
		)

		return
	}

	defer archivo.Close()

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

	fmt.Println()
	fmt.Println("===== BITMAP INODOS =====")

	for i := int32(0); i < sb.SInodesCount; i++ {

		valor, err := LeerByte(
			archivo,
			sb.SBmInodeStart+i,
		)

		if err != nil {
			return
		}

		fmt.Printf("%d ", valor)

		if (i+1)%20 == 0 {
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println("=========================")
}


// ReporteBMBlock: Muestra el bitmap de bloques ocupados/libres.

func ReporteBMBlock(
	particion estructuras.ParticionMontada,
	path string,
) {

	archivo, err := os.Open(
		particion.Path,
	)

	if err != nil {

		fmt.Println(
			"ERROR abriendo disco",
		)

		return
	}

	defer archivo.Close()

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

	fmt.Println()
	fmt.Println("===== BITMAP BLOQUES =====")

	for i := int32(0); i < sb.SBlocksCount; i++ {

		valor, err := LeerByte(
			archivo,
			sb.SBmBlockStart+i,
		)

		if err != nil {
			return
		}

		fmt.Printf("%d ", valor)

		if (i+1)%20 == 0 {
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println("=========================")
}


// ReporteINODE: Muestra todos los inodos ocupados encontrados en el bitmap de inodos.

func ReporteINODE(
	particion estructuras.ParticionMontada,
	path string,
) {

	archivo, err := os.Open(
		particion.Path,
	)

	if err != nil {

		fmt.Println(
			"ERROR abriendo disco",
		)

		return
	}

	defer archivo.Close()

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

	fmt.Println()
	fmt.Println("===== REPORTE INODE =====")

	for i := int32(0); i < sb.SInodesCount; i++ {

		valor, err := LeerByte(
			archivo,
			sb.SBmInodeStart+i,
		)

		if err != nil {
			return
		}

		if valor == 0 {
			continue
		}

		inodo, err := LeerInodoPorNumero(
			archivo,
			sb,
			i,
		)

		if err != nil {
			return
		}

		fmt.Println()
		fmt.Println(
			"INODO",
			i,
		)

		fmt.Println(
			"UID:",
			inodo.IUid,
		)

		fmt.Println(
			"GID:",
			inodo.IGid,
		)

		fmt.Println(
			"SIZE:",
			inodo.ISize,
		)

		fmt.Printf(
			"TYPE: %c\n",
			inodo.IType,
		)

		fmt.Println(
			"PERM:",
			inodo.IPerm,
		)

		for j := 0; j < 15; j++ {

			fmt.Printf(
				"IBlock[%d] = %d\n",
				j,
				inodo.IBlock[j],
			)
		}

		fmt.Println(
			"---------------------",
		)
	}

	fmt.Println(
		"======================",
	)
}

// LeerInodoPorNumero: Lee un inodo utilizando su número lógico.

func LeerInodoPorNumero(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numero int32,
) (
	estructuras.Inode,
	error,
) {

	posicion := ObtenerPosicionInodo(
		sb,
		numero,
	)

	return LeerInodo(
		archivo,
		posicion,
	)
}



// ReporteBLOCK: Recorre los inodos ocupados y muestra los bloques asociados a cada uno.

func ReporteBLOCK(
	particion estructuras.ParticionMontada,
	path string,
) {

	archivo, err := os.Open(
		particion.Path,
	)

	if err != nil {

		fmt.Println(
			"ERROR abriendo disco",
		)

		return
	}

	defer archivo.Close()

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

	fmt.Println()
	fmt.Println("===== REPORTE BLOCK =====")

	for i := int32(0); i < sb.SInodesCount; i++ {

		valor, err := LeerByte(
			archivo,
			sb.SBmInodeStart+i,
		)

		if err != nil {
			return
		}

		if valor == 0 {
			continue
		}

		inodo, err := LeerInodoPorNumero(
			archivo,
			sb,
			i,
		)

		if err != nil {
			return
		}

		// Directorio
		if inodo.IType == '0' {

			for j := 0; j < 12; j++ {

				if inodo.IBlock[j] == -1 {
					continue
				}

				folder, err :=
					LeerFolderPorNumero(
						archivo,
						sb,
						inodo.IBlock[j],
					)

				if err != nil {
					continue
				}

				fmt.Println()
				fmt.Println(
					"FOLDER BLOCK",
					inodo.IBlock[j],
				)

				for k := 0; k < 4; k++ {

					nombre :=
						utilidades.BytesAString(
							folder.BContent[k].BName[:],
						)

					fmt.Printf(
						"[%d] %s -> %d\n",
						k,
						nombre,
						folder.BContent[k].BInodo,
					)
				}

				fmt.Println(
					"---------------------",
				)
			}
		}

		// Archivo
		if inodo.IType == '1' {

			for j := 0; j < 12; j++ {

				if inodo.IBlock[j] == -1 {
					continue
				}

				fileBlock, err :=
					LeerFilePorNumero(
						archivo,
						sb,
						inodo.IBlock[j],
					)

				if err != nil {
					continue
				}

				fmt.Println()
				fmt.Println(
					"FILE BLOCK",
					inodo.IBlock[j],
				)

				fmt.Println(
					utilidades.BytesAString(
						fileBlock.BContent[:],
					),
				)

				fmt.Println(
					"---------------------",
				)
			}
		}
	}

	fmt.Println(
		"======================",
	)
}