
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"MIA_P1_9516098/analizador"
//	"MIA_P1_9516098/api"
)

func main() {

	// Inicia API
	// go api.StartServer()

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

		comando := strings.TrimSpace(
			lector.Text(),
		)

		if strings.ToLower(comando) == "exit" {

			fmt.Println("Finalizando programa...")
			break
		}

		analizador.Analizar(comando)
	}
}

