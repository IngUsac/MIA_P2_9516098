package analizador

import (
	"strings"

	"MIA_P1_9516098/comandos"
)

func Analizar(comando string) {

	comando = strings.TrimSpace(comando)

	if comando == "" {
		return
	}

	partes := strings.Fields(comando)

	switch strings.ToLower(partes[0]) {

	case "mkdisk":

		parametros := ObtenerParametros(comando)

		comandos.EjecutarMKDISK(parametros)

	case "fdisk":

		parametros := ObtenerParametros(comando)

		comandos.EjecutarFDISK(parametros)

	case "mount":

		parametros := ObtenerParametros(comando)

		comandos.Mount(parametros,)	

	default:
		println("Comando no reconocido")
	}
}