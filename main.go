package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
	"MIA_P1_9516098/analizador"
)

func main() {

	fmt.Println("====================================")
	fmt.Println("  PROYECTO 1 - MIA")
	fmt.Println("  SISTEMA DE ARCHIVOS EXT2")
	fmt.Println("====================================")

	lector := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">> ")

		if !lector.Scan() {
			break
		}

		comando := strings.TrimSpace(lector.Text())

		if strings.ToLower(comando) == "exit" {
			fmt.Println("Finalizando programa...")
			break
		}

		analizador.Analizar(comando)
	}
}