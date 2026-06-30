package comandos

import (
	"fmt"
	"os"
	"strings"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)

func EliminarArchivo(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroInodo int32,
	inodo estructuras.Inode,
) error {

	fmt.Println("Liberando bloques del archivo...")

	for _, bloque := range inodo.IBlock {

		if bloque == -1 {
			continue
		}

		posicionBloque := sb.SBlockStart + (bloque * sb.SBlockSize)

		err := LiberarBloqueArchivo(
			archivo,
			posicionBloque,
		)

		if err != nil {
			return err
		}

		// Liberar bitmap del bloque
		err = utilidades.EscribirObjeto(
			archivo,
			new(byte),
			int64(sb.SBmBlockStart+bloque),
		)

		if err != nil {
			return err
		}

		sb.SFreeBlocksCount++
	}

	posicionInodo := sb.SInodeStart + (numeroInodo * sb.SInodeSize)

	err := LiberarInodo(
		archivo,
		posicionInodo,
	)

	if err != nil {
		return err
	}

	// Liberar bitmap del inodo
	err = utilidades.EscribirObjeto(
		archivo,
		new(byte),
		int64(sb.SBmInodeStart+numeroInodo),
	)

	if err != nil {
		return err
	}

	sb.SFreeInodesCount++

	err = EscribirSuperBlock(
		archivo,
		sb,
		estructuras.SesionActual.Start,
	)

	if err != nil {
		return err
	}

	fmt.Println("Bloques liberados.")
	fmt.Println("Inodo liberado.")

	return nil
}



func EliminarCarpetaRecursiva(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroInodo int32,
	inodo estructuras.Inode,
) error {

	fmt.Println("Recorriendo carpeta:", numeroInodo)

	for _, bloque := range inodo.IBlock {

		if bloque == -1 {
			continue
		}

		posicionBloque := sb.SBlockStart + (bloque * sb.SBlockSize)

		folder, err := LeerFolderBlock(
			archivo,
			posicionBloque,
		)

		if err != nil {
			continue
		}

		for _, entrada := range folder.BContent {

			if entrada.BInodo == -1 {
				continue
			}

			nombre := strings.TrimRight(
				string(entrada.BName[:]),
				"\x00",
			)

			if nombre == "" ||
				nombre == "." ||
				nombre == ".." {
				continue
			}

			fmt.Println("Encontrado:", nombre)

			posicionInodo := sb.SInodeStart +
				(entrada.BInodo * sb.SInodeSize)

			hijo, err := LeerInodo(
				archivo,
				posicionInodo,
			)

			if err != nil {
				continue
			}

			if hijo.IType == '1' {

				err = EliminarArchivo(
					archivo,
					sb,
					entrada.BInodo,
					hijo,
				)

			} else {

				err = EliminarCarpetaRecursiva(
					archivo,
					sb,
					entrada.BInodo,
					hijo,
				)

			}

			if err != nil {
				return err
			}
		}
	}

	fmt.Println("Liberando carpeta:", numeroInodo)

	return nil
}

func LiberarInodo(
	archivo *os.File,
	posicion int32,
) error {

	var inodo estructuras.Inode

	for i := range inodo.IBlock {
		inodo.IBlock[i] = -1
	}

	return utilidades.EscribirObjeto(
		archivo,
		&inodo,
		int64(posicion),
	)
}

func LiberarBloqueArchivo(
	archivo *os.File,
	posicion int32,
) error {

	var bloque estructuras.FileBlock

	return EscribirFileBlock(
		archivo,
		bloque,
		posicion,
	)
}

func EliminarEntradaPadre(
	archivo *os.File,
	sb estructuras.SuperBlock,
	ruta string,
) error {

	ruta = strings.TrimRight(ruta, "/")

	partes := strings.Split(ruta, "/")

	if len(partes) < 2 {
		return nil
	}

	nombre := partes[len(partes)-1]

	rutaPadre := "/"

	if len(partes) > 2 {
		rutaPadre += strings.Join(
			partes[1:len(partes)-1],
			"/",
		)
	}

	inodoPadre, _, err := ObtenerInodoPorRutaCompleta(
		archivo,
		sb,
		rutaPadre,
	)

	if err != nil {
		return err
	}

	for _, bloque := range inodoPadre.IBlock {

		if bloque == -1 {
			continue
		}

		posicionBloque := sb.SBlockStart + (bloque * sb.SBlockSize)

		fb, err := LeerFolderBlock(
			archivo,
			posicionBloque,
		)

		if err != nil {
			continue
		}

		for i := 0; i < 4; i++ {

			n := strings.TrimRight(
				string(fb.BContent[i].BName[:]),
				"\x00",
			)

			if n == nombre {

				fb.BContent[i].BInodo = -1

				for j := range fb.BContent[i].BName {
					fb.BContent[i].BName[j] = 0
				}

				return EscribirFolderBlock(
					archivo,
					fb,
					posicionBloque,
				)
			}
		}
	}

	return nil
}
