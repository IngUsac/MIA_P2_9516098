package comandos

import (
	"fmt"
	"os"
	"strings"

	"MIA_P1_9516098/estructuras"
)

func MKDIR(parametros map[string]string) {

	fmt.Println(" MKDIR ")
	fmt.Println()

	if !estructuras.SesionActual.Activa {

		fmt.Println(
			"ERROR: no existe una sesion activa",
		)

		return
	}

	path := parametros["path"]

	if path == "" {

		fmt.Println(
			"ERROR: falta parametro path",
		)

		return
	}

	path = strings.Trim(
		path,
		"\"",
	)

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

	sb, err := LeerSuperBlock(
		archivo,
		estructuras.SesionActual.Start,
	)

	if err != nil {

		fmt.Println(
			"ERROR leyendo SuperBlock",
		)

		return
	}

	componentes := SepararRuta(path)

	if len(componentes) == 0 {

		fmt.Println(
			"ERROR: ruta invalida",
		)

		return
	}

	numeroActual := int32(0)

	rutaActual := ""

	seCreo := false

	for _, nombre := range componentes {

		if rutaActual == "" {
			rutaActual = "/" + nombre
		} else {
			rutaActual += "/" + nombre
		}

		_, numeroExistente, errRuta :=
			ObtenerInodoPorRutaCompleta(
				archivo,
				sb,
				rutaActual,
			)

		if errRuta == nil {
				fmt.Println(
				"Ya existe carpeta: ",
				rutaActual,
			)
		

			numeroActual = numeroExistente
			continue
		}

		fmt.Println( "Creando... ", rutaActual,)


		err = MKDIRInterno(
			archivo,
			&sb,
			numeroActual,
			nombre,
		)

		seCreo = true

		if err != nil {

			fmt.Println(
				"ERROR:",
				err,
			)

			return
		}

		_, numeroNuevo, errRuta :=
			ObtenerInodoPorRutaCompleta(
				archivo,
				sb,
				rutaActual,
			)

		if errRuta != nil {

			fmt.Println(
				"ERROR creando:",
				rutaActual,
			)

			return
		}

		numeroActual = numeroNuevo
	}

		

	err = ActualizarSuperBlock(
		archivo,
		sb,
		estructuras.SesionActual.Start,
	)

	if err != nil {

		fmt.Println(
			"ERROR actualizando SuperBlock",
		)

		return
	}

	if seCreo {

		fmt.Println(
			"Directorio creado --> ",
			path,
		)

	} else {

		fmt.Println(
			"ERROR: el directorio ya existe",
		)
	}
}