package comandos

import (
	"fmt"

	"MIA_P1_9516098/estructuras"
)

// LOGOUT: Finaliza la sesión activa.

func LOGOUT() {

	if !estructuras.SesionActual.Activa {

		fmt.Println(
			"ERROR: no existe una sesion activa",
		)

		return
	}

	estructuras.SesionActual = estructuras.Sesion{}

	fmt.Println()
	fmt.Println("===== LOGOUT =====")
	fmt.Println("Sesion cerrada correctamente")
	fmt.Println()
}