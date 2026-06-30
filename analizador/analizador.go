

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
	parametros := ObtenerParametros(comando)

	switch strings.ToLower(partes[0]) {

	case "mkdisk": 	comandos.EjecutarMKDISK(parametros)

	case "fdisk":	comandos.EjecutarFDISK(parametros)

	case "mount":	comandos.MOUNT(parametros,)	

	case "mkfs":	comandos.MKFS(parametros,)	

	case "login":	comandos.LOGIN(parametros)

	case "logout":	comandos.LOGOUT()	

	case "mkgrp":	comandos.MKGRP(parametros)		
	
	case "rmgrp":	comandos.RMGRP(parametros)

	case "mkusr":	comandos.MKUSR(parametros)

	case "execute":	EXECUTE(parametros)

	case "pause":	PAUSE()

	case "rmusr":	comandos.RMUSR(parametros)

	case "chgrp":	comandos.CHGRP(parametros)
		
	case "cat":		comandos.CAT(parametros)

	case "mkdir":	comandos.MKDIR(parametros)

	case "mkfile":	comandos.MKFILE(parametros)

	case "rep":		comandos.REP(parametros)	

	case "rename":	comandos.EjecutarRename(parametros)

	case "remove":	comandos.EjecutarRemove(parametros)

	case "edit":	comandos.EjecutarEdit(parametros)

			

	default:
		println("ERROR: Comando no reconocido: ",comando)
	}
}
