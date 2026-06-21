package analizador

import (
	"bufio"
	"fmt"
	"os"
)
// hace una pausa hasta que se presione enter para continuar

func PAUSE() {
	fmt.Println("Pausa... Presione ENTER para continuar...")
	fmt.Println()

	reader := bufio.NewReader(
		os.Stdin,
	)

	reader.ReadString('\n')
}