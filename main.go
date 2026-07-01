package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"MIA_P1_9516098/analizador"
	"MIA_P1_9516098/api"
)

func main() {

	os.MkdirAll("SALIDAS", 0755)
	os.MkdirAll("SALIDAS/discos", 0755)
	os.MkdirAll("SALIDAS/reportes", 0755)
	
	// MODO SERVIDOR (FASE 2)
	// Ejecutar con:
	// go run . server
	
	if len(os.Args) > 1 && strings.ToLower(os.Args[1]) == "server" {

		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("   PROYECTO 1 - BACKEND REST - Main")
		fmt.Println(" ")

		api.StartServer()
		return
	}

	
	// MODO CONSOLA (FASE 1)
	// Ejecutar con:
	// go run .
	

	fmt.Println(" ")
	fmt.Println("  MANEJO E IMPLEMENTACION DE ARCHIVOS")
	fmt.Println("              PROYECTO 1 ")
	fmt.Println("  SISTEMA DE ARCHIVOS EXT2")
	fmt.Println(" ")

	lector := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println(" ")
		fmt.Print("Comando --->> ")

		if !lector.Scan() {
			break
		}

		comando := strings.TrimSpace(lector.Text())

		if strings.EqualFold(comando, "exit") {

			fmt.Println("Finalizando programa...")
			break
		}

		analizador.Analizar(comando)
	}
}