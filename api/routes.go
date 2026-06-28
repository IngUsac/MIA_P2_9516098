package api

import (
	"fmt"
	"net/http"
)

func StartServer() {

	// Estado del Backend
	http.HandleFunc("/api/status", StatusHandler)

	// Intérprete de comandos
	http.HandleFunc("/api/execute", GenericCommandHandler)

	// Visualizador - Discos
	http.HandleFunc("/api/disks", GetDisksHandler)

	http.HandleFunc("/api/partitions",GetPartitionsHandler,)

	fmt.Println()
	fmt.Println()
	fmt.Println("   PROYECTO 1 - BACKEND REST")
	fmt.Println()
	fmt.Println("        API REST MIA")
	fmt.Println()
	fmt.Println("Puerto : 8080")
	fmt.Println("URL    : http://localhost:8080")
	fmt.Println()
	fmt.Println("GET  /api/status")
	fmt.Println("POST /api/execute")
	fmt.Println("GET  /api/disks")
	fmt.Println()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error iniciando servidor:", err)
	}
}