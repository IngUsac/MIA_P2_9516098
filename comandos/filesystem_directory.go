package comandos


import (
	"os"
	"fmt"
	"strings"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)

func CambiarNombreEntrada(
	archivo *os.File,
	sb estructuras.SuperBlock,
	ruta string,
	nuevoNombre string,
) error {

	ruta = strings.TrimRight(ruta, "/")

	partes := strings.Split(ruta, "/")

	if len(partes) < 2 {
		return fmt.Errorf("ruta invalida")
	}

	nombreActual := partes[len(partes)-1]

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

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FolderBlock{},
		),
	)

	for _, bloque := range inodoPadre.IBlock {

		if bloque == -1 {
			continue
		}

		posicion := sb.SBlockStart +
			(bloque * blockSize)

		folder, err := LeerFolderBlock(
			archivo,
			posicion,
		)

		if err != nil {
			continue
		}

		for i := 0; i < 4; i++ {

			nombre := utilidades.BytesAString(
				folder.BContent[i].BName[:],
			)

			if nombre == nuevoNombre {

				return fmt.Errorf(
					"ya existe una entrada con ese nombre",
				)
			}
		}

		for i := 0; i < 4; i++ {

			nombre := utilidades.BytesAString(
				folder.BContent[i].BName[:],
			)

			if nombre == nombreActual {

				for j := range folder.BContent[i].BName {
					folder.BContent[i].BName[j] = 0
				}

				copy(
					folder.BContent[i].BName[:],
					[]byte(nuevoNombre),
				)

				return EscribirFolderBlock(
					archivo,
					folder,
					posicion,
				)
			}
		}
	}

	return fmt.Errorf("entrada no encontrada")
}

// CrearDirectorioCompleto:
// Crea un directorio y lo agrega inmediatamente al directorio padre.

func CrearDirectorioCompleto(
	archivo *os.File,
	sb *estructuras.SuperBlock,
	numeroPadre int32,
	nombre string,
) (int32, error) {

	numeroInodo, err := CrearDirectorio(
		archivo,
		sb,
		numeroPadre,
		nombre,
	)

	if err != nil {
		return -1, err
	}

	err = AgregarArchivoEnPadre(
		archivo,
		*sb,
		numeroPadre,
		nombre,
		numeroInodo,
	)

	if err != nil {
		return -1, err
	}

	return numeroInodo, nil
}

// CopiarArchivo:  Copia un archivo existente hacia un directorio destino.  Retorna el número de inodo creado.

func CopiarArchivo(
	archivo *os.File,
	sb *estructuras.SuperBlock,
	rutaOrigen string,
	numeroPadreDestino int32,
	nuevoNombre string,
) (int32, error) {

	// Obtener el archivo origen

	inodoOrigen, _, err := ObtenerInodoPorRutaCompleta(
		archivo,
		*sb,
		rutaOrigen,
	)

	if err != nil {
		return -1, err
	}

	if inodoOrigen.IType != '1' {

		return -1,
			fmt.Errorf("la ruta origen no es un archivo")
	}

	// Leer contenido

	contenido, err := LeerContenidoArchivo(
		archivo,
		*sb,
		inodoOrigen,
	)

	if err != nil {
		return -1, err
	}

	// Crear copia

	numNuevo, err := CrearArchivo(
		archivo,
		sb,
		numeroPadreDestino,
		nuevoNombre,
		contenido,
	)

	if err != nil {
		return -1, err
	}

	return numNuevo, nil
}

// CopiarDirectorio:  Copia un directorio completo de forma recursiva.

func CopiarDirectorio(
	archivo *os.File,
	sb *estructuras.SuperBlock,
	rutaOrigen string,
	numeroPadreDestino int32,
	nuevoNombre string,
) (int32, error) {

	// Obtener inodo origen

	inodoOrigen, _, err := ObtenerInodoPorRutaCompleta(
		archivo,
		*sb,
		rutaOrigen,
	)

	if err != nil {
		return -1, err
	}

	if inodoOrigen.IType != '0' {

		return -1,
			fmt.Errorf("la ruta origen no es un directorio")
	}

	// Crear directorio destino

	numNuevo, err := CrearDirectorioCompleto(
		archivo,
		sb,
		numeroPadreDestino,
		nuevoNombre,
	)

	if err != nil {
		return -1, err
	}

	//**--

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FolderBlock{},
		),
	)

	for _, bloque := range inodoOrigen.IBlock {

		if bloque == -1 {
			continue
		}

		posicion := sb.SBlockStart +
			(bloque * blockSize)

		folder, err := LeerFolderBlock(
			archivo,
			posicion,
		)

		if err != nil {
			continue
		}

		for _, entrada := range folder.BContent {

			if entrada.BInodo == -1 {
				continue
			}

			nombre := utilidades.BytesAString(
				entrada.BName[:],
			)

			if nombre == "." ||
				nombre == ".." {
				continue
			}

			posHijo := ObtenerPosicionInodo(
				*sb,
				entrada.BInodo,
			)

			inodoHijo, err := LeerInodo(
				archivo,
				posHijo,
			)

			if err != nil {
				continue
			}

			rutaHijo := rutaOrigen + "/" + nombre

			if inodoHijo.IType == '1' {

				_, err = CopiarArchivo(
					archivo,
					sb,
					rutaHijo,
					numNuevo,
					nombre,
				)

			} else {

				_, err = CopiarDirectorio(
					archivo,
					sb,
					rutaHijo,
					numNuevo,
					nombre,
				)
			}

			if err != nil {
				return -1, err
			}
		}
	}
	//**--

	
	// La copia recursiva se implementará en el siguiente paso

	return numNuevo, nil
}