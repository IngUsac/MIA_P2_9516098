package api

import (
	"fmt"
	"net/http"
)

/*
	Endpoint principal de la API.
*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API MIA funcionando correctamente")
}

/*
	StartServer

	Inicia el servidor HTTP del Backend.
*/
func StartServer() {

	// Página principal
	http.HandleFunc("/", HomeHandler)

	// Estado del Backend
	http.HandleFunc("/api/status", StatusHandler)

	
	// Endpoint único del intérprete
	
	http.HandleFunc("/api/execute", GenericCommandHandler)

	fmt.Println("")
	fmt.Println(" ")
	fmt.Println("        API REST MIA")
	fmt.Println(" ")
	fmt.Println("Puerto : 8080")
	fmt.Println("URL    : http://localhost:8080")
	fmt.Println("")
	fmt.Println("GET  /api/status")
	fmt.Println("POST /api/execute")
	fmt.Println("")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}