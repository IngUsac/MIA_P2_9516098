package comandos

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"MIA_P1_9516098/estructuras"
	)

func MKFILE(parametros map[string]string) {

	fmt.Println("  MKFILE, parametros", parametros ) 
	fmt.Println()
	
	// Validar sesión activa

	if !estructuras.SesionActual.Activa {

		fmt.Println(
			"ERROR: no existe una sesion activa",
		)

		return
	}

	// Path obligatorio

	path, existe := parametros["path"]

	if !existe || strings.TrimSpace(path) == "" {

		fmt.Println(
			"ERROR: falta parametro path",
		)

		return
	}

	// Parámetro -r

	_, crearPadres := parametros["r"]

	// Parámetro -size

	size := 0

	if valor, ok := parametros["size"]; ok {

		numero, err := strconv.Atoi(
			valor,
		)

		if err != nil || numero < 0 {

			fmt.Println(
				"ERROR: size invalido",
			)

			return
		}

		size = numero
	}

	// Parámetro -cont

	contenido := ""

	if valor, ok := parametros["cont"]; ok {

		contenidoBytes, err :=
			os.ReadFile(
				valor,
			)

		if err != nil {

			fmt.Println(
				"ERROR: archivo cont no existe",
			)

			return
		}

		contenido =
			string(
				contenidoBytes,
			)
	}

	// Buscar partición montada asociada a la sesión

	particion, existe :=
		BuscarParticionMontadaPorID(
			estructuras.SesionActual.ID,
		)

	if !existe {

		fmt.Println(
			"ERROR: particion no montada",
		)

		return
	}

	archivo, err := os.OpenFile(
		particion.Path,
		os.O_RDWR,
		0644,
	)

	if err != nil {

		fmt.Println(
			"ERROR abriendo disco:",
			err,
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
			"ERROR leyendo superblock:",
			err,
		)

		return
	}

	// Verificar si ya existe

	_, _, errExiste :=
		ObtenerInodoPorRutaCompleta(
			archivo,
			sb,
			path,
		)

	if errExiste == nil {

		fmt.Println(
			"ERROR: el archivo ya existe",
		)

		return
	}

		 
	// Obtener carpeta padre y nombre archivo
	 

	partes := SepararRuta(
		path,
	)

	if len(partes) == 0 {

		fmt.Println(
			"ERROR: path invalido",
		)

		return
	}
		


	nombreArchivo :=partes[len(partes)-1]

	if len(nombreArchivo) > 12 {
		fmt.Println("ERROR: nombre de archivo excede 12 caracteres",nombreArchivo)

		return
	}

	rutaPadre := "/"

	if len(partes) > 1 {

		rutaPadre += strings.Join(
			partes[:len(partes)-1],
			"/",
		)
	}

	 
	// Buscar carpeta padre
	 

	_, numeroPadre, err :=
		ObtenerInodoPorRutaCompleta(
			archivo,
			sb,
			rutaPadre,
		)

	if err != nil {

		if !crearPadres {

			fmt.Println( "ERROR: ruta padre no existe", )

			return
		}

		componentes := SepararRuta(
			rutaPadre,
		)

		rutaActual := "/"

		numeroActual := int32(0)

		for _, carpeta := range componentes {

			if rutaActual == "/" {

				rutaActual += carpeta

			} else {

				rutaActual += "/" + carpeta
			}

			_, numeroExistente, errRuta :=
				ObtenerInodoPorRutaCompleta(
					archivo,
					sb,
					rutaActual,
				)

			if errRuta == nil {

				numeroActual =
					numeroExistente

				continue
			}

			err = MKDIRInterno(
				archivo,
				&sb,
				numeroActual,
				carpeta,
			)

			if err != nil {

				fmt.Println(
					"ERROR:",
					err,
				)

				return
			}
			
			// Recargar SuperBlock actualizado

			sb, err = LeerSuperBlock(
				archivo,
				particion.Start,
			)

			if err != nil {

				fmt.Println(
					"ERROR leyendo superblock:",
					err,
				)

				return
			}

			_, numeroActual, err =
				ObtenerInodoPorRutaCompleta(
					archivo,
					sb,
					rutaActual,
				)

			if err != nil {

				fmt.Println(
					"ERROR:",
					err,
				)

				return
			}

			
		}
		sb, err = LeerSuperBlock(
			archivo,
			particion.Start,
		)

		if err != nil {

			fmt.Println(
				"ERROR leyendo superblock:",
				err,
			)

			return
		}
		_, numeroPadre, err =
			ObtenerInodoPorRutaCompleta(
				archivo,
				sb,
				rutaPadre,
			)

		if err != nil {

			fmt.Println(
				"ERROR: no se pudo crear ruta padre",
			)

			return
		}
	}

	 
	// Generar contenido por size
	 

	if contenido == "" &&
		size > 0 {

		for i := 0; i < size; i++ {

			contenido += fmt.Sprintf(
				"%d",
				i%10,
			)
		}
	}

	 
	// Crear archivo
	 

	_, err = CrearArchivo(
		archivo,
		&sb,
		numeroPadre,
		nombreArchivo,
		contenido,
	)

	if err != nil {

		fmt.Println(
			"ERROR:",
			err,
		)

		return
	}

	fmt.Println(
		"Archivo creado:",
		path,
	)
}