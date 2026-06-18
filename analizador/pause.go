package analizador

import (
	"bufio"
	"fmt"
	"os"
)
// hace una pausa hasta que se presione enter para continuar

func PAUSE() {

	fmt.Println()
	fmt.Println(
		"Presione ENTER para continuar...",
	)

	reader := bufio.NewReader(
		os.Stdin,
	)

	reader.ReadString('\n')
}