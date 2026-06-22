package api

import (
	"fmt"
	"net/http"
)

/*
	Muestra un mensaje al entrar a:
	http://localhost:8080
*/
func HomeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintf(
		w,
		"API MIA funcionando correctamente",
	)
}

/*
	Inicia el servidor y registra las rutas.
*/
func StartServer() {

	// Ruta principal
	http.HandleFunc("/", HomeHandler)

	// Endpoint para crear discos
	http.HandleFunc(
		"/mkdisk",
		MkDiskHandler,
	)

	fmt.Println("")
	fmt.Println("API escuchando en puerto 8080")
	fmt.Println("http://localhost:8080/mkdisk")
	fmt.Println("")

	err := http.ListenAndServe(
		":8080",
		nil,
	)

	if err != nil {
		fmt.Println(
			"Error al iniciar servidor:",
			err,
		)
	}
}