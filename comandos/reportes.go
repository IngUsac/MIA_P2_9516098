package comandos

import (
	"fmt"
	"os"
	//"strings"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)

// ReporteSB: Genera el reporte del SuperBlock y lo exporta a PDF.

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

	datos := [][]string{

		{
			"SFilesystemType",
			fmt.Sprintf(
				"%d",
				sb.SFilesystemType,
			),
		},

		{
			"SInodesCount",
			fmt.Sprintf(
				"%d",
				sb.SInodesCount,
			),
		},

		{
			"SBlocksCount",
			fmt.Sprintf(
				"%d",
				sb.SBlocksCount,
			),
		},

		{
			"SFreeBlocksCount",
			fmt.Sprintf(
				"%d",
				sb.SFreeBlocksCount,
			),
		},

		{
			"SFreeInodesCount",
			fmt.Sprintf(
				"%d",
				sb.SFreeInodesCount,
			),
		},

		{
			"SMntCount",
			fmt.Sprintf(
				"%d",
				sb.SMntCount,
			),
		},

		{
			"SMagic",
			fmt.Sprintf(
				"%d",
				sb.SMagic,
			),
		},

		{
			"SInodeSize",
			fmt.Sprintf(
				"%d",
				sb.SInodeSize,
			),
		},

		{
			"SBlockSize",
			fmt.Sprintf(
				"%d",
				sb.SBlockSize,
			),
		},

		{
			"SFirstIno",
			fmt.Sprintf(
				"%d",
				sb.SFirstIno,
			),
		},

		{
			"SFirstBlo",
			fmt.Sprintf(
				"%d",
				sb.SFirstBlo,
			),
		},

		{
			"SBmInodeStart",
			fmt.Sprintf(
				"%d",
				sb.SBmInodeStart,
			),
		},

		{
			"SBmBlockStart",
			fmt.Sprintf(
				"%d",
				sb.SBmBlockStart,
			),
		},

		{
			"SInodeStart",
			fmt.Sprintf(
				"%d",
				sb.SInodeStart,
			),
		},

		{
			"SBlockStart",
			fmt.Sprintf(
				"%d",
				sb.SBlockStart,
			),
		},
	}

	err = GenerarPDFTabla(
		path,
		"Reporte de SUPERBLOQUE",
		datos,
	)

	if err != nil {

		fmt.Println(
			"ERROR generando PDF",
		)

		return
	}

	fmt.Println(
		"PDF generado:",
		path,
	)
}



// ReporteBMInode: Muestra todos los inodos ocupados encontrados en el bitmap de inodos.

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

	var contenido string

	for i := int32(0); i < sb.SInodesCount; i++ {

		valor, err := LeerByte(
			archivo,
			sb.SBmInodeStart+i,
		)

		if err != nil {
			return
		}

		contenido += fmt.Sprintf(
			"%d ",
			valor,
		)

		if (i+1)%20 == 0 {
			contenido += "\n"
		}
	}

	

	err = GenerarPDFTexto(
		path,
		"BITMAP DE INODOS",
		contenido,
	)



	if err != nil {

		fmt.Println(
			"ERROR generando PDF",
		)

		return
	}

	fmt.Println(
		"PDF generado:",
		path,
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

// ReporteBMBlock: Recorre los inodos ocupados y muestra los bloques asociados a cada uno.

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

	var contenido string

	for i := int32(0); i < sb.SBlocksCount; i++ {

		valor, err := LeerByte(
			archivo,
			sb.SBmBlockStart+i,
		)

		if err != nil {
			return
		}

		contenido += fmt.Sprintf(
			"%d ",
			valor,
		)

		if (i+1)%20 == 0 {
			contenido += "\n"
		}
	}


	err = GenerarPDFTexto(
		path,
		"BITMAP DE BLOQUES",
		contenido,
	)

	if err != nil {

		fmt.Println(
			"ERROR generando PDF",
		)

		return
	}

	fmt.Println(
		"PDF generado:",
		path,
	)
}


// ReporteBLOCK: Genera el reporte de bloques utilizados y lo exporta a PDF.

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

	var contenido string

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

		// Directorios

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

				contenido +=
					"=================================\n"

				contenido += fmt.Sprintf(
					"FOLDER BLOCK %d\n",
					inodo.IBlock[j],
				)

				contenido +=
					"=================================\n"

				for k := 0; k < 4; k++ {

					nombre :=
						utilidades.BytesAString(
							folder.BContent[k].BName[:],
						)

					contenido += fmt.Sprintf(
						"[%d] %s -> %d\n",
						k,
						nombre,
						folder.BContent[k].BInodo,
					)
				}

				contenido += "\n"
			}
		}

		// Archivos

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

				contenido +=
					"=================================\n"

				contenido += fmt.Sprintf(
					"FILE BLOCK %d\n",
					inodo.IBlock[j],
				)

				contenido +=
					"=================================\n"

				contenido +=
					utilidades.BytesAString(
						fileBlock.BContent[:],
					)

				contenido += "\n\n"
			}
		}
	}

	err = GenerarPDFTexto(
		path,
		"REPORTE DE BLOQUES",
		contenido,
	)

	if err != nil {

		fmt.Println(
			"ERROR generando PDF",
		)

		return
	}

	fmt.Println(
		"PDF generado:",
		path,
	)
}


// ReporteMBR: Genera el reporte del MBR y lo exporta a PDF.

func ReporteMBR(
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

	mbr, err := LeerMBR(
		archivo,
	)

	if err != nil {

		fmt.Println(
			"ERROR leyendo MBR",
		)

		return
	}

	var datos [][]string

	datos = append(
		datos,
		[]string{
			"MbrTamano",
			fmt.Sprintf(
				"%d",
				mbr.MbrTamano,
			),
		},
	)

	datos = append(
		datos,
		[]string{
			"MbrDiskSignature",
			fmt.Sprintf(
				"%d",
				mbr.MbrDiskSignature,
			),
		},
	)

	datos = append(
		datos,
		[]string{
			"DskFit",
			string(
				mbr.DskFit,
			),
		},
	)

	for i := 0; i < 4; i++ {

		part := mbr.MbrPartitions[i]

		if part.PartSize <= 0 {
			continue
		}

		prefijo := fmt.Sprintf(
			"Particion %d ",
			i+1,
		)

		datos = append(
			datos,
			[]string{
				prefijo + "Status",
				fmt.Sprintf(
					"%d",
					part.PartStatus,
				),
			},
		)

		datos = append(
			datos,
			[]string{
				prefijo + "Type",
				string(
					part.PartType,
				),
			},
		)

		datos = append(
			datos,
			[]string{
				prefijo + "Fit",
				string(
					part.PartFit,
				),
			},
		)

		datos = append(
			datos,
			[]string{
				prefijo + "Start",
				fmt.Sprintf(
					"%d",
					part.PartStart,
				),
			},
		)

		datos = append(
			datos,
			[]string{
				prefijo + "Size",
				fmt.Sprintf(
					"%d",
					part.PartSize,
				),
			},
		)

		datos = append(
			datos,
			[]string{
				prefijo + "Name",
				utilidades.BytesAString(
					part.PartName[:],
				),
			},
		)

		if part.PartType == 'E' {

			posActual := part.PartStart
			contadorEBR := 1

			for {

				ebr, err := LeerEBR(
					archivo,
					posActual,
				)

				if err != nil {
					break
				}

				if ebr.PartSize <= 0 {
					break
				}

				prefijoEBR := fmt.Sprintf(
					"EBR %d ",
					contadorEBR,
				)

				datos = append(
					datos,
					[]string{
						prefijoEBR + "Fit",
						string(
							ebr.PartFit,
						),
					},
				)

				datos = append(
					datos,
					[]string{
						prefijoEBR + "Start",
						fmt.Sprintf(
							"%d",
							ebr.PartStart,
						),
					},
				)

				datos = append(
					datos,
					[]string{
						prefijoEBR + "Size",
						fmt.Sprintf(
							"%d",
							ebr.PartSize,
						),
					},
				)

				datos = append(
					datos,
					[]string{
						prefijoEBR + "Next",
						fmt.Sprintf(
							"%d",
							ebr.PartNext,
						),
					},
				)

				datos = append(
					datos,
					[]string{
						prefijoEBR + "Name",
						utilidades.BytesAString(
							ebr.PartName[:],
						),
					},
				)

				if ebr.PartNext == -1 {
					break
				}

				posActual = ebr.PartNext
				contadorEBR++
			}
		}
	}

	err = GenerarPDFTabla(
		path,
		"REPORTE MBR",
		datos,
	)

	if err != nil {

		fmt.Println(
			"ERROR generando PDF",
		)

		return
	}

	fmt.Println(
		"PDF generado:",
		path,
	)
}

// ReporteDISK: Genera el reporte de distribución del disco y lo exporta a PDF.

func ReporteDISK(
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

	mbr, err := LeerMBR(
		archivo,
	)

	if err != nil {

		fmt.Println(
			"ERROR leyendo MBR",
		)

		return
	}

	var datos [][]string

	var usado int32 = 169 // tamaño aproximado del MBR

	datos = append(
		datos,
		[]string{
			"MBR",
			"Sistema",
			fmt.Sprintf("%d", 169),
			fmt.Sprintf(
				"%.2f%%",
				float64(169)*100.0/
					float64(mbr.MbrTamano),
			),
		},
	)

	for i := 0; i < 4; i++ {

		part := mbr.MbrPartitions[i]

		if part.PartSize <= 0 {
			continue
		}

		nombre :=
			utilidades.BytesAString(
				part.PartName[:],
			)

		tipo := "Primaria"

		if part.PartType == 'E' {
			tipo = "Extendida"
		}

		datos = append(
			datos,
			[]string{
				nombre,
				tipo,
				fmt.Sprintf(
					"%d",
					part.PartSize,
				),
				fmt.Sprintf(
					"%.2f%%",
					float64(part.PartSize)*100.0/
						float64(mbr.MbrTamano),
				),
			},
		)

		usado += part.PartSize

		// Recorrer particiones lógicas

		if part.PartType == 'E' {

			posActual :=
				part.PartStart

			for {

				ebr, err :=
					LeerEBR(
						archivo,
						posActual,
					)

				if err != nil {
					break
				}

				if ebr.PartSize <= 0 {
					break
				}

				nombreLogica :=
					utilidades.BytesAString(
						ebr.PartName[:],
					)

				datos = append(
					datos,
					[]string{
						nombreLogica,
						"Logica",
						fmt.Sprintf(
							"%d",
							ebr.PartSize,
						),
						fmt.Sprintf(
							"%.2f%%",
							float64(ebr.PartSize)*100.0/
								float64(mbr.MbrTamano),
						),
					},
				)

				if ebr.PartNext == -1 {
					break
				}

				posActual =
					ebr.PartNext
			}
		}
	}

	libre :=
		mbr.MbrTamano - usado

	if libre < 0 {
		libre = 0
	}

	datos = append(
		datos,
		[]string{
			"Libre",
			"Libre",
			fmt.Sprintf(
				"%d",
				libre,
			),
			fmt.Sprintf(
				"%.2f%%",
				float64(libre)*100.0/
					float64(mbr.MbrTamano),
			),
		},
	)

	err = GenerarPDFTabla(
		path,
		"REPORTE DISK",
		datos,
	)

	if err != nil {

		fmt.Println(
			"ERROR generando PDF",
		)

		return
	}

	fmt.Println(
		"PDF generado:",
		path,
	)
}

// ReporteINODE: Genera el reporte de todos los inodos ocupados y lo exporta a PDF.

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

	var contenido string

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

		contenido += fmt.Sprintf(
			"=================================\n",
		)

		contenido += fmt.Sprintf(
			"INODO %d\n",
			i,
		)

		contenido += fmt.Sprintf(
			"=================================\n",
		)

		contenido += fmt.Sprintf(
			"UID: %d\n",
			inodo.IUid,
		)

		contenido += fmt.Sprintf(
			"GID: %d\n",
			inodo.IGid,
		)

		contenido += fmt.Sprintf(
			"SIZE: %d\n",
			inodo.ISize,
		)

		contenido += fmt.Sprintf(
			"TYPE: %c\n",
			inodo.IType,
		)

		contenido += fmt.Sprintf(
			"PERM: %d\n",
			inodo.IPerm,
		)

		contenido += "\n"

		for j := 0; j < 15; j++ {

			contenido += fmt.Sprintf(
				"IBlock[%d] = %d\n",
				j,
				inodo.IBlock[j],
			)
		}

		contenido += "\n"
	}

	err = GenerarPDFTexto(
		path,
		"REPORTE DE INODOS",
		contenido,
	)

	if err != nil {

		fmt.Println(
			"ERROR generando PDF",
		)

		return
	}

	fmt.Println(
		"PDF generado:",
		path,
	)
}
/*
func ReporteTREE(
	particion estructuras.ParticionMontada,
	path string,
)*/