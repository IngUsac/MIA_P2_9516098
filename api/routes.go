package api

import (
	"fmt"
	"net/http"
)

/*
EnableCORS

Middleware sencillo para permitir que el Frontend
React pueda consumir la API REST.
*/
func EnableCORS(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	

		w.Header().Set(
			"Access-Control-Allow-Origin",
			"http://localhost:3000",
		)

		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, OPTIONS",
		)

		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type",
		)

		if r.Method == http.MethodOptions {

			w.WriteHeader(http.StatusOK)
			return

		}

		handler.ServeHTTP(w, r)

	})

}
func StartServer() {

	// Estado del Backend
	http.HandleFunc("/api/status", StatusHandler)

	// Intérprete de comandos
	http.HandleFunc("/api/execute", GenericCommandHandler)

	// Visualizador - Discos
	http.HandleFunc("/api/disks", GetDisksHandler)

	// Visualizador - Particiones
	http.HandleFunc("/api/partitions", GetPartitionsHandler)

	// Particiones montadas
	http.HandleFunc("/api/mounts", GetMountsHandler)

	// Visualizador - Árbol
	http.HandleFunc("/api/tree", GetTreeHandler)

	// Visualizador - Reportes
	http.HandleFunc("/api/reports", GetReportsHandler)

	// Sesión
	http.HandleFunc("/api/session", GetSessionHandler)

	http.HandleFunc("/api/files", GetFilesHandler)

	


	
	fmt.Println("URL    : http://localhost:8080")
	fmt.Println()

	//**--

	http.Handle(
		"/reportes/",
		http.StripPrefix(
			"/reportes/",
			http.FileServer(
				http.Dir("./SALIDAS/reportes"),
			),
		),
	)
	//**--


	err := http.ListenAndServe(
		":8080",
		EnableCORS(http.DefaultServeMux),
	)

	if err != nil {
		fmt.Println("ERROR iniciando servidor:", err)
	}
}