package comandos	
import (
	"fmt"
	"os"
	"strings"

	"MIA_P1_9516098/estructuras"
	"MIA_P1_9516098/utilidades"
)

	

// LeerInodo:  Lee un inodo desde una posición específica del disco.
// Parámetros: 
// archivo  -> disco abierto
// posicion -> byte donde inicia el inodo

// Retorna: Inodo leído o  error en caso de fallo

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

// BuscarEntradaEnFolder: Busca un nombre dentro de un FolderBlock.
//
// Parámetros:
// folder -> bloque de carpeta
// nombre -> nombre buscado
//
// Retorna:
// número de inodo
// true si existe
//
func BuscarEntradaEnFolder(
	folder estructuras.FolderBlock,
	nombre string,
) (int32, bool) {

	for _, entrada := range folder.BContent {

		nombreActual := utilidades.BytesAString(
			entrada.BName[:],
		)

		if nombreActual == nombre {

			return entrada.BInodo, true
		}
	}

	return -1, false
}

// ObtenerInodoPorRuta: Busca un archivo dentro de la raíz.
// Ejemplo:  /users.txt
// Retorna: 
// inodo encontrado
// posición del inodo

func ObtenerInodoPorRuta(
	archivo *os.File,
	sb estructuras.SuperBlock,
	ruta string,
) (estructuras.Inode, int32, error) {

	nombre := strings.TrimPrefix(
		ruta,
		"/",
	)

	folder, err := LeerFolderBlock(
		archivo,
		sb.SBlockStart,
	)

	if err != nil {
		return estructuras.Inode{}, 0, err
	}

	numInodo, existe := BuscarEntradaEnFolder(
		folder,
		nombre,
	)

	if !existe {

		return estructuras.Inode{},
			0,
			fmt.Errorf(
				"archivo no existe",
			)
	}

	inodeSize := int32(
		utilidades.ObtenerTamano(
			estructuras.Inode{},
		),
	)

	posicion := sb.SInodeStart +
		(numInodo * inodeSize)

	inode, err := LeerInodo(
		archivo,
		posicion,
	)

	if err != nil {
		return estructuras.Inode{}, 0, err
	}



	return inode,
		posicion,
		nil



}

// LeerContenidoArchivo: Lee el contenido completo de un archivo utilizando
// bloques directos e indirecto simple.
// Parámetros:
// archivo -> disco abierto
// sb      -> SuperBlock
// inode   -> inodo del archivo
// Retorna: contenido completo respetando inode.ISize.

func LeerContenidoArchivo(
	archivo *os.File,
	sb estructuras.SuperBlock,
	inode estructuras.Inode,
) (string, error) {

	var contenido []byte

	// Directos
	for i := 0; i < 12; i++ {

		if inode.IBlock[i] == -1 {
			break
		}

		fileBlock, err := LeerFilePorNumero(
			archivo,
			sb,
			inode.IBlock[i],
		)

		if err != nil {
			return "", err
		}

		contenido = append(
			contenido,
			fileBlock.BContent[:]...,
		)
	}

	// Indirecto simple
	if inode.IBlock[12] != -1 {

		pointer, err := LeerPointerPorNumero(
			archivo,
			sb,
			inode.IBlock[12],
		)

		if err != nil {
			return "", err
		}

		for i := 0; i < 16; i++ {

			if pointer.BPointers[i] == -1 {
				break
			}

			fileBlock, err :=
				LeerFilePorNumero(
					archivo,
					sb,
					pointer.BPointers[i],
				)

			if err != nil {
				return "", err
			}

			contenido = append(
				contenido,
				fileBlock.BContent[:]...,
			)
		}
	}

	if int(inode.ISize) < len(contenido) {

		contenido =
			contenido[:inode.ISize]
	}

	return string(contenido), nil
}

// SepararRuta:  Convierte una ruta absoluta en sus componentes.
// Ejemplo:  "/home/docs/test"
// Retorna:  ["home", "docs", "test"]

func SepararRuta(
	ruta string,
) []string {

	ruta = strings.TrimSpace(
		ruta,
	)

	ruta = strings.TrimPrefix(
		ruta,
		"/",
	)

	if ruta == "" {
		return []string{}
	}

	componentes := strings.Split(
		ruta,
		"/",
	)

	return componentes
}


// BuscarComponenteEnInodo: Busca un nombre dentro de los FolderBlocks apuntados por un inodo de carpeta.
// Parámetros:
// archivo    -> disco abierto
// sb         -> SuperBlock
// inode      -> inodo carpeta
// componente -> nombre a buscar
// Retorna: número de inodo encontrado y  true si existe

func BuscarComponenteEnInodo(
	archivo *os.File,
	sb estructuras.SuperBlock,
	inode estructuras.Inode,
	componente string,
) (int32, bool) {

	blockSize := int32(
		utilidades.ObtenerTamano(
			estructuras.FolderBlock{},
		),
	)

	for i := 0; i < 15; i++ {

		if inode.IBlock[i] == -1 {
			break
		}

		posBloque := sb.SBlockStart +
			(inode.IBlock[i] * blockSize)

		folder, err := LeerFolderBlock(
			archivo,
			posBloque,
		)

		if err != nil {
			continue
		}

		for _, entrada := range folder.BContent {

			nombre := utilidades.BytesAString(
				entrada.BName[:],
			)

			if nombre == componente {

				return entrada.BInodo,
					true
			}
		}
	}

	return -1, false

}

// ObtenerInodoPorRutaCompleta: Recorre una ruta absoluta componente por componente.
// Ejemplo:  /home/docs/proyecto
// Retorna:  inodo encontrado y  número de inodo,  error si no existe.

func ObtenerInodoPorRutaCompleta(
	archivo *os.File,
	sb estructuras.SuperBlock,
	ruta string,
) (estructuras.Inode, int32, error) {

	componentes := SepararRuta(
		ruta,
	)

	// Ruta raíz
	if len(componentes) == 0 {

		posRaiz := ObtenerPosicionInodo(
			sb,
			0,
		)

		inodeRaiz, err := LeerInodo(
			archivo,
			posRaiz,
		)

		if err != nil {
			return estructuras.Inode{}, 0, err
		}

		return inodeRaiz, 0, nil
	}

	// Empezar desde la raíz
	numeroActual := int32(0)

	posActual := ObtenerPosicionInodo(
		sb,
		numeroActual,
	)

	inodeActual, err := LeerInodo(
		archivo,
		posActual,
	)

	if err != nil {
		return estructuras.Inode{}, 0, err
	}

	for _, componente := range componentes {

		numeroSiguiente, existe :=
			BuscarComponenteEnInodo(
				archivo,
				sb,
				inodeActual,
				componente,
			)

		if !existe {

			return estructuras.Inode{},
				0,
				fmt.Errorf(
					"ruta no existe: %s",
					ruta,
				)
		}

		posSiguiente := ObtenerPosicionInodo(
			sb,
			numeroSiguiente,
		)

		inodeActual, err = LeerInodo(
			archivo,
			posSiguiente,
		)

		if err != nil {
			return estructuras.Inode{}, 0, err
		}

		numeroActual = numeroSiguiente
	}

	return inodeActual,
		numeroActual,
		nil
}

// BuscarEspacioLibreFolder:  Busca una entrada libre dentro de un FolderBlock.
// Retorna:  índice de la entrada libre y true si existe espacio.

func BuscarEspacioLibreFolder(
	folder estructuras.FolderBlock,
) (int, bool) {

	for i := 0; i < 4; i++ {

		if folder.BContent[i].BInodo == -1 {

			return i, true
		}
	}

	return -1, false
}

// AgregarEntradaFolder:Inserta una nueva entrada dentro de un FolderBlock.
// Parámetros:
// folder       -> bloque a modificar
// nombre       -> nombre del archivo/directorio
// numeroInodo  -> inodo asociado

func AgregarEntradaFolder(
	folder *estructuras.FolderBlock,
	nombre string,
	numeroInodo int32,
) bool {

	indice, existe := BuscarEspacioLibreFolder(
		*folder,
	)

	if !existe {
		return false
	}

	copy(
		folder.BContent[indice].BName[:],
		nombre,
	)

	folder.BContent[indice].BInodo =
		numeroInodo

	return true
}

// CrearFolderBlockDirectorio:  Crea el FolderBlock inicial de un directorio.
// Contenido:
// .  -> apunta a sí mismo
// .. -> apunta al padre
func CrearFolderBlockDirectorio(
	inodoActual int32,
	inodoPadre int32,
) estructuras.FolderBlock {

	var folder estructuras.FolderBlock

	for i := 0; i < 4; i++ {
		folder.BContent[i].BInodo = -1
	}

	copy(
		folder.BContent[0].BName[:],
		".",
	)

	folder.BContent[0].BInodo =
		inodoActual

	copy(
		folder.BContent[1].BName[:],
		"..",
	)

	folder.BContent[1].BInodo =
		inodoPadre

	return folder
}




// CrearInodoDirectorio: Crea un inodo para un directorio.
// El bloque indicado será el primer FolderBlock asociado al directorio.

func CrearInodoDirectorio(
	numeroBloque int32,
) estructuras.Inode {

	var inode estructuras.Inode

	inode.IUid = 1
	inode.IGid = 1

	for i := 0; i < 15; i++ {

		inode.IBlock[i] = -1
	}

	inode.IBlock[0] =
		numeroBloque

	inode.IType = '0'

	inode.IPerm = 664

	return inode
}

// BuscarPrimerInodoLibre: Busca el primer inodo libre en el bitmap.
// Retorna: número de inodo libre o  error si no existe.

func BuscarPrimerInodoLibre(
	archivo *os.File,
	sb estructuras.SuperBlock,
) (int32, error) {

	for i := int32(0); i < sb.SInodesCount; i++ {

		valor, err := LeerByte(
			archivo,
			sb.SBmInodeStart+i,
		)

		if err != nil {
			return -1, err
		}

		if valor == 0 {
			return i, nil
		}
	}

	return -1,
		fmt.Errorf(
			"no hay inodos libres",
		)
}

// BuscarPrimerBloqueLibre: Busca el primer bloque libre en el bitmap.
// Retorna: número de bloque libre o error si no existe.

func BuscarPrimerBloqueLibre(
	archivo *os.File,
	sb estructuras.SuperBlock,
) (int32, error) {

	for i := int32(0); i < sb.SBlocksCount; i++ {

		valor, err := LeerByte(
			archivo,
			sb.SBmBlockStart+i,
		)

		if err != nil {
			return -1, err
		}

		if valor == 0 {
			return i, nil
		}
	}

	return -1,
		fmt.Errorf(
			"no hay bloques libres",
		)
}

// OcuparInodo: Marca un inodo como ocupado en el bitmap.

func OcuparInodo(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numero int32,
) error {

	var ocupado byte = 1

	return utilidades.EscribirObjeto(
		archivo,
		&ocupado,
		int64(
			sb.SBmInodeStart + numero,
		),
	)
}

// OcuparBloque: Marca un bloque como ocupado en el bitmap.

func OcuparBloque(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numero int32,
) error {

	var ocupado byte = 1

	return utilidades.EscribirObjeto(
		archivo,
		&ocupado,
		int64(
			sb.SBmBlockStart + numero,
		),
	)
}

// ObtenerPosicionBloque: Calcula la posición física de un bloque.

func ObtenerPosicionBloque(
	sb estructuras.SuperBlock,
	numero int32,
) int32 {

	return sb.SBlockStart +
		(numero * sb.SBlockSize)
}

// GuardarFolderBlock: Guarda un FolderBlock en disco.

func GuardarFolderBlock(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroBloque int32,
	folder estructuras.FolderBlock,
) error {

	posicion := ObtenerPosicionBloque(
		sb,
		numeroBloque,
	)

	return EscribirFolderBlock(
		archivo,
		folder,
		posicion,
	)
}

// GuardarInodo: Guarda un inodo en una posición lógica.

func GuardarInodo(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroInodo int32,
	inode estructuras.Inode,
) error {

	posicion := ObtenerPosicionInodo(
		sb,
		numeroInodo,
	)

	return EscribirInodo(
		archivo,
		inode,
		posicion,
	)
}

// ActualizarSuperBlock: Sobrescribe el SuperBlock actualizado al inicio de la partición.

func ActualizarSuperBlock(
	archivo *os.File,
	sb estructuras.SuperBlock,
	inicioParticion int32,
) error {

	return EscribirSuperBlock(
		archivo,
		sb,
		inicioParticion,
	)
}

// ReservarRecursosDirectorio: Reserva un inodo y un bloque para un nuevo directorio.

func ReservarRecursosDirectorio(
	archivo *os.File,
	sb *estructuras.SuperBlock,
) (int32, int32, error) {

	numInodo, err := BuscarPrimerInodoLibre(
		archivo,
		*sb,
	)

	if err != nil {
		return -1, -1, err
	}

	numBloque, err := BuscarPrimerBloqueLibre(
		archivo,
		*sb,
	)

	if err != nil {
		return -1, -1, err
	}

	err = OcuparInodo(
		archivo,
		*sb,
		numInodo,
	)

	if err != nil {
		return -1, -1, err
	}

	err = OcuparBloque(
		archivo,
		*sb,
		numBloque,
	)

	if err != nil {
		return -1, -1, err
	}

	sb.SFreeInodesCount--
	sb.SFreeBlocksCount--

	sb.SFirstIno = numInodo + 1
	sb.SFirstBlo = numBloque + 1

	return numInodo,
		numBloque,
		nil
}



// LeerFolderPorNumero: Lee un FolderBlock utilizando su número lógico.

func LeerFilePorNumero(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroBloque int32,
) (estructuras.FileBlock, error) {

	posicion := ObtenerPosicionBloque(
		sb,
		numeroBloque,
	)

	return LeerFileBlock(
		archivo,
		posicion,
	)
}



// GuardarFolderPorNumero: Guarda un FolderBlock utilizando su número lógico.

func GuardarFolderPorNumero(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroBloque int32,
	folder estructuras.FolderBlock,
) error {

	posicion := ObtenerPosicionBloque(
		sb,
		numeroBloque,
	)

	return EscribirFolderBlock(
		archivo,
		folder,
		posicion,
	)
}

// LeerFolderPorNumero: Lee un FolderBlock utilizando su número lógico.
// Parámetros:
// archivo      -> disco abierto
// sb           -> SuperBlock
// numeroBloque -> número lógico del bloque
// Retorna: FolderBlock leído.

func LeerFolderPorNumero(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroBloque int32,
) (
	estructuras.FolderBlock,
	error,
) {

	posicion := ObtenerPosicionBloque(
		sb,
		numeroBloque,
	)

	return LeerFolderBlock(
		archivo,
		posicion,
	)
}


// AgregarDirectorioEnPadre:  Inserta una nueva entrada dentro de la carpeta padre.
// Parámetros:
// archivo      -> disco abierto
// sb           -> superblock
// numeroPadre  -> inodo del directorio padre
// nombre       -> nombre del nuevo directorio
// numeroNuevo  -> inodo del nuevo directorio

func AgregarDirectorioEnPadre(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroPadre int32,
	nombre string,
	numeroNuevo int32,
) error {

	posPadre := ObtenerPosicionInodo(
		sb,
		numeroPadre,
	)

	inodePadre, err := LeerInodo(
		archivo,
		posPadre,
	)

	if err != nil {
		return err
	}

	// Buscar FolderBlock asociado al padre
	for i := 0; i < 15; i++ {
		if inodePadre.IBlock[i] == -1 {

		numBloque, err := BuscarPrimerBloqueLibre(
			archivo,
			sb,
		)

		if err != nil {
			return err
		}
				
		nuevoFolder := CrearFolderBlockDirectorio(
			numeroPadre,
			numeroPadre,
		)
				

		err = GuardarFolderPorNumero(
			archivo,
			sb,
			numBloque,
			nuevoFolder,
		)

		if err != nil {
			return err
		}

		err = OcuparBloque(
			archivo,
			sb,
			numBloque,
		)

		if err != nil {
			return err
		}

		inodePadre.IBlock[i] = numBloque

		err = GuardarInodo(
			archivo,
			sb,
			numeroPadre,
			inodePadre,
		)

		if err != nil {
			return err
		}

		folder := nuevoFolder

		ok := AgregarEntradaFolder(
			&folder,
			nombre,
			numeroNuevo,
		)

		if !ok {
			return fmt.Errorf("error insertando entrada")
		}

		return GuardarFolderPorNumero(
			archivo,
			sb,
			numBloque,
			folder,
		)
	}

		folder, err := LeerFolderPorNumero(
			archivo,
			sb,
			inodePadre.IBlock[i],
		)

		if err != nil {
			continue
		}

		ok := AgregarEntradaFolder(
			&folder,
			nombre,
			numeroNuevo,
		)

		if ok {

			return GuardarFolderPorNumero(
				archivo,
				sb,
				inodePadre.IBlock[i],
				folder,
			)
		}
	}

	return fmt.Errorf(
		"no hay espacio en carpeta padre",
	)
}


// MKDIRInterno:  Crea un directorio dentro de otro directorio.
// Ejemplo:
// padre = 0 (/)
// nombre = home

func MKDIRInterno(
	archivo *os.File,
	sb *estructuras.SuperBlock,
	numeroPadre int32,
	nombre string,
) error {



	// Crear directorio físicamente y enlazarlo al padre

	_, err := CrearDirectorioCompleto(
		archivo,
		sb,
		numeroPadre,
		nombre,
	)

	if err != nil {
		return err
	}


	// Crear directorio físicamente
	/*
	numNuevo, err := CrearDirectorio(
		archivo,
		sb,
		numeroPadre,
		nombre,
	)

	if err != nil {
		return err
	}

	// Agregar entrada en el padre

	err = AgregarDirectorioEnPadre(
		archivo,
		*sb,
		numeroPadre,
		nombre,
		numNuevo,
	)

	if err != nil {
		return err
	}
*/
	return nil
}

// CrearDirectorio:  Crea físicamente un directorio en disco.
// Retorna el número de inodo asignado.

func CrearDirectorio(
	archivo *os.File,
	sb *estructuras.SuperBlock,
	numeroPadre int32,
	nombre string,
) (int32, error) {

	numInodo,
		numBloque,
		err := ReservarRecursosDirectorio(
		archivo,
		sb,
	)

	if err != nil {
		return -1, err
	}

	folder := CrearFolderBlockDirectorio(
		numInodo,
		numeroPadre,
	)

	err = GuardarFolderBlock(
		archivo,
		*sb,
		numBloque,
		folder,
	)

	if err != nil {
		return -1, err
	}

	inode := CrearInodoDirectorio(
		numBloque,
	)

	err = GuardarInodo(
		archivo,
		*sb,
		numInodo,
		inode,
	)

	if err != nil {
		return -1, err
	}

	return numInodo, nil
}


func CrearInodoArchivo(
	bloques []int32,
	size int32,
) estructuras.Inode {

	var inode estructuras.Inode

	for i := 0; i < 15; i++ {
		inode.IBlock[i] = -1
	}

	inode.IUid = estructuras.SesionActual.UID
	inode.IGid = estructuras.SesionActual.GID

	inode.ISize = size

	inode.IType = '1'

	inode.IPerm = 664

	for i := 0; i < len(bloques) && i < 15; i++ {
		inode.IBlock[i] = bloques[i]
	}

	if len(bloques) > 12 {

		inode.IBlock[12] =
			bloques[12]
	}


	return inode
}

// GuardarFileBlock: guarda un bloque de archivo utilizando número lógico.

func GuardarFileBlock(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroBloque int32,
	fileBlock estructuras.FileBlock,
) error {

	posicion := ObtenerPosicionBloque(
		sb,
		numeroBloque,
	)

	return EscribirFileBlock(
		archivo,
		fileBlock,
		posicion,
	)
}

// ReservarRecursosArchivo: reserva un inodo y un bloque.

func ReservarRecursosArchivo(
	archivo *os.File,
	sb *estructuras.SuperBlock,
) (
	int32,
	int32,
	error,
) {

	return ReservarRecursosDirectorio(
		archivo,
		sb,
	)
}


// CrearArchivo: crea físicamente un archivo dentro del filesystem. El contenido es dividido en bloques de 64 bytes utilizando apuntadores directos e indirecto simple.
// Parámetros:
// archivo      -> disco abierto
// sb           -> SuperBlock
// numeroPadre  -> inodo del directorio padre
// nombre       -> nombre del archivo
// contenido    -> contenido completo a almacenar
// Retorna:
// número de inodo asignado
// error si ocurre algún problema durante la creación.

func CrearArchivo(
	archivo *os.File,
	sb *estructuras.SuperBlock,
	numeroPadre int32,
	nombre string,
	contenido string,
) (
	int32,
	error,
) {

	numInodo, err := BuscarPrimerInodoLibre(
		archivo,
		*sb,
	)

	if err != nil {
		return -1, err
	}

	err = OcuparInodo(
		archivo,
		*sb,
		numInodo,
	)

	if err != nil {
		return -1, err
	}

	sb.SFreeInodesCount--
	sb.SFirstIno = numInodo + 1

	inode := CrearInodoArchivo(
		nil,
		0,
	)

	err = EscribirContenidoArchivo(
		archivo,
		sb,
		&inode,
		contenido,
	)

	if err != nil {
		return -1, err
	}




	err = GuardarInodo(
		archivo,
		*sb,
		numInodo,
		inode,
	)

	if err != nil {
		return -1, err
	}

	err = AgregarArchivoEnPadre(
		archivo,
		*sb,
		numeroPadre,
		nombre,
		numInodo,
	)

	if err != nil {
		return -1, err
	}

	return numInodo, nil
}

func AgregarArchivoEnPadre(
	archivo *os.File,
	sb estructuras.SuperBlock,
	numeroPadre int32,
	nombre string,
	numeroInodo int32,
) error {

	return AgregarDirectorioEnPadre(
		archivo,
		sb,
		numeroPadre,
		nombre,
		numeroInodo,
	)
}

func ReservarBloqueArchivo(
	archivo *os.File,
	sb *estructuras.SuperBlock,
) (int32, error) {

	numBloque, err := BuscarPrimerBloqueLibre(
		archivo,
		*sb,
	)

	if err != nil {
		return -1, err
	}

	err = OcuparBloque(
		archivo,
		*sb,
		numBloque,
	)

	if err != nil {
		return -1, err
	}

	sb.SFreeBlocksCount--
	sb.SFirstBlo = numBloque + 1

	return numBloque, nil
}

func EscribirContenidoArchivo(
	archivo *os.File,
	sb *estructuras.SuperBlock,
	inode *estructuras.Inode,
	contenido string,
) error {

	// Liberar bloques actuales del archivo
	for _, bloque := range inode.IBlock {

		if bloque == -1 {
			continue
		}

		posicion := ObtenerPosicionBloque(
			*sb,
			bloque,
		)

		err := LiberarBloqueArchivo(
			archivo,
			posicion,
		)

		if err != nil {
			return err
		}

		var libre byte = 0

		err = utilidades.EscribirObjeto(
			archivo,
			&libre,
			int64(sb.SBmBlockStart+bloque),
		)

		if err != nil {
			return err
		}

		sb.SFreeBlocksCount++
	}

	// Limpiar apuntadores actuales
	for i := 0; i < len(inode.IBlock); i++ {
		inode.IBlock[i] = -1
	}

	contenidoBytes := []byte(contenido)

	cantidadBloques := (len(contenidoBytes) + 63) / 64

	if cantidadBloques == 0 {
		cantidadBloques = 1
	}

	if cantidadBloques > 28 {
		return fmt.Errorf("archivo excede indirecto simple")
	}

	var bloquesDirectos []int32
	var bloquesIndirectos []int32

	for i := 0; i < cantidadBloques; i++ {

		numBloque, err := ReservarBloqueArchivo(
			archivo,
			sb,
		)

		if err != nil {
			return err
		}

		var fileBlock estructuras.FileBlock

		inicio := i * 64
		fin := inicio + 64

		if fin > len(contenidoBytes) {
			fin = len(contenidoBytes)
		}

		if inicio < len(contenidoBytes) {

			copy(
				fileBlock.BContent[:],
				contenidoBytes[inicio:fin],
			)
		}

		err = GuardarFileBlock(
			archivo,
			*sb,
			numBloque,
			fileBlock,
		)

		if err != nil {
			return err
		}

		if i < 12 {

			bloquesDirectos = append(
				bloquesDirectos,
				numBloque,
			)

		} else {

			bloquesIndirectos = append(
				bloquesIndirectos,
				numBloque,
			)
		}
	}

	if len(bloquesIndirectos) > 0 {

	numeroPointer, err := ReservarBloqueArchivo(
		archivo,
		sb,
	)

	if err != nil {
		return err
	}

	var pointer estructuras.PointerBlock

	for i := 0; i < 16; i++ {
		pointer.BPointers[i] = -1
	}

	for i := range bloquesIndirectos {
		pointer.BPointers[i] = bloquesIndirectos[i]
	}

	err = GuardarPointerPorNumero(
		archivo,
		*sb,
		numeroPointer,
		pointer,
	)

	if err != nil {
		return err
	}

	inode.IBlock[12] = numeroPointer
}

	for i := range bloquesDirectos {
		inode.IBlock[i] = bloquesDirectos[i]
	}

	inode.ISize = int32(len(contenidoBytes))


	err := EscribirSuperBlock(
		archivo,
		*sb,
		estructuras.SesionActual.Start,
	)

	if err != nil {
		return err
	}

	return nil
}