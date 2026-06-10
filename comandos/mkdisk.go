package comandos

import "fmt"

func EjecutarMKDISK(parametros map[string]string) {

	fmt.Println("===== MKDISK =====")

	for clave, valor := range parametros {
		fmt.Printf("%s = %s\n", clave, valor)
	}
}