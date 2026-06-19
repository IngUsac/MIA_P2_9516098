package comandos

import (
	"fmt"
	"os"
	"sort"
	"strings"
	

	"MIA_P1_9516098/estructuras"
)

// CAT: Muestra el contenido de uno o varios archivos.
// Ejemplos:  cat -file1="/users.txt"     / cat -file1="/users.txt" -file2="/otro.txt"
// El contenido se muestra en el mismo orden en que se reciben los parámetros file1, file2, file3, etc.

func CAT(
	parametros map[string]string,
) {

	fmt.Println()
	fmt.Println(
		"===== CAT =====",
	)

	// --------------------------------------------------------
	// Validar sesión
	// --------------------------------------------------------

	if !estructuras.SesionActual.Activa {

		fmt.Println(
			"ERROR: no existe una sesion activa",
		)

		return
	}

	// --------------------------------------------------------
	// Obtener rutas
	// --------------------------------------------------------

	type ArchivoCAT struct {
		Indice int
		Ruta   string
	}

	var archivos []ArchivoCAT

	for clave, valor := range parametros {

		claveLower := strings.ToLower(
			strings.TrimSpace(clave),
		)

		if strings.HasPrefix(
			claveLower,
			"file",
		) {

			var indice int

			fmt.Sscanf(
				claveLower,
				"file%d",
				&indice,
			)

			ruta := strings.Trim(
				strings.TrimSpace(valor),
				"\"",
			)

			archivos = append(
				archivos,
				ArchivoCAT{
					Indice: indice,
					Ruta:   ruta,
				},
			)
		}
	}

	if len(archivos) == 0 {

		fmt.Println(
			"ERROR: falta parametro file",
		)

		return
	}

	// --------------------------------------------------------
	// Ordenar file1, file2, file3...
	// --------------------------------------------------------

	sort.Slice(
		archivos,
		func(i, j int) bool {

			return archivos[i].Indice <
				archivos[j].Indice
		},
	)

	// --------------------------------------------------------
	// Abrir disco
	// --------------------------------------------------------

	archivo, err := os.OpenFile(
		estructuras.SesionActual.Path,
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

	// --------------------------------------------------------
	// Leer SuperBlock
	// --------------------------------------------------------

	sb, err := LeerSuperBlock(
		archivo,
		estructuras.SesionActual.Start,
		//estructuras.SesionActual.Path,
	)

	if err != nil {

		fmt.Println(
			"ERROR leyendo SuperBlock",
		)

		return
	}

	// --------------------------------------------------------
	// Leer cada archivo solicitado
	// --------------------------------------------------------

	for _, item := range archivos {

		fmt.Println()
		fmt.Println(
			"Archivo:",
			item.Ruta,
		)

		inode, _, err := ObtenerInodoPorRuta(
			archivo,
			sb,
			item.Ruta,
		)

		if err != nil {

			fmt.Println(
				"ERROR:",
				item.Ruta,
				"no existe",
			)

			continue
		}

		// Validar que sea archivo

		if inode.IType != '1' {

			fmt.Println(
				"ERROR:",
				item.Ruta,
				"no es un archivo",
			)

			continue
		}

		contenido, err := LeerContenidoArchivo(
			archivo,
			sb,
			inode,
		)

		if err != nil {

			fmt.Println(
				"ERROR leyendo archivo",
			)

			continue
		}

		fmt.Println(
			contenido,
		)
	}
}