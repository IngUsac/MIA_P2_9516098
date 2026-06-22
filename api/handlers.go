package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"MIA_P1_9516098/analizador"
)

/*
	Estructura esperada en el JSON.
*/
type MkDiskRequest struct {
	Size int    `json:"size"`
	Unit string `json:"unit"`
	Fit  string `json:"fit"`
	Path string `json:"path"`
}

/*
	POST /mkdisk

	Ejemplo:

	{
		"size":75,
		"unit":"m",
		"fit":"wf",
		"path":"./SALIDAS/pruebas/d1.dsk"
	}
*/
func MkDiskHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Solo aceptar POST
	if r.Method != http.MethodPost {

		http.Error(
			w,
			"Metodo no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	var req MkDiskRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"JSON invalido",
			http.StatusBadRequest,
		)

		return
	}

	// Construir comando MIA
	comando := fmt.Sprintf(
		"mkdisk -size=%d -unit=%s -fit=%s -path=%s",
		req.Size,
		req.Unit,
		req.Fit,
		req.Path,
	)

	// Mostrar en consola para depuración
	fmt.Println(" ")
	fmt.Println("COMANDO RECIBIDO DESDE API: ",comando)
	fmt.Println(" ")

	// Ejecutar utilizando tu analizador existente
	analizador.Analizar(comando)

	// Respuesta JSON
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]string{
			"mensaje": "Solicitud enviada correctamente",
		},
	)
}