package main
<<<<<<< HEAD
=======

>>>>>>> 37b9f52cca376eb771b410ea59d0dd6ec547f7cb
import (
	"bufio"
	"fmt"
	"os"
	"strings"
<<<<<<< HEAD
	
	"MIA_P1_9516098/analizador"
=======
>>>>>>> 37b9f52cca376eb771b410ea59d0dd6ec547f7cb
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

<<<<<<< HEAD
		analizador.Analizar(comando)
=======
		fmt.Println("Comando recibido:", comando)
>>>>>>> 37b9f52cca376eb771b410ea59d0dd6ec547f7cb
	}
}