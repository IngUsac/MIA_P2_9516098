package analizador

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// EXECUTE: archivo tipo bat
// Ejecuta línea por línea un archivo .mia 
// Ejemplo:  >>execute -path="./pruebas/login.mia"
//
func EXECUTE(
	parametros map[string]string,
) {

	path := parametros["path"]

	if path == "" {

		fmt.Println(
			"ERROR: falta parametro path",
		)

		return
	}

	archivo, err := os.Open(
		path,
	)

	if err != nil {

		fmt.Println(
			"ERROR: no se pudo abrir el archivo",
		)

		return
	}

	defer archivo.Close()

	fmt.Println()
	fmt.Println("===== EXECUTE =====")
	fmt.Println("Path:", path)

	scanner := bufio.NewScanner(
		archivo,
	)

	for scanner.Scan() {

		linea := strings.TrimSpace(
			scanner.Text(),
		)

		if linea == "" {
			continue
		}

		// Ignorar comentarios
		if strings.HasPrefix(
			linea,
			"#",
		) {
			continue
		}

		fmt.Println(">>", linea)

		Analizar(
			linea,
		)
	}
}