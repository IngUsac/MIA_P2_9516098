package api

import (
	"encoding/json"
	"net/http"

	"MIA_P1_9516098/comandos"
	"MIA_P1_9516098/estructuras"
)

func GetFilesHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	if !estructuras.SesionActual.Activa {

		http.Error(
			w,
			"No existe una sesión activa.",
			http.StatusUnauthorized,
		)

		return
	}

	ruta := r.URL.Query().Get("path")

	if ruta == "" {
		ruta = "/"
	}

	archivo, err := comandos.AbrirDisco(
		estructuras.SesionActual.Path,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	defer archivo.Close()

	sb, err := comandos.LeerSuperBlock(
		archivo,
		estructuras.SesionActual.Start,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	lista, err := comandos.ListarDirectorio(
		archivo,
		sb,
		ruta,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		lista,
	)

}