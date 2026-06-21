package comandos

import (
	"fmt"
	"strings"
	
	"MIA_P1_9516098/estructuras"
)


// REP: Genera reportes del sistema de archivos.
// Parámetros:
// name -> nombre del reporte
// path -> ruta destino del reporte
// id   -> id de partición montada

func REP(
	parametros map[string]string,
) {

	fmt.Println(" REP, parametros", parametros)
	fmt.Println()

	name, ok := parametros["name"]

	if !ok || strings.TrimSpace(name) == "" {

		fmt.Println(
			"ERROR: falta parametro name",
		)

		return
	}

	path, ok := parametros["path"]

	if !ok || strings.TrimSpace(path) == "" {

		fmt.Println(
			"ERROR: falta parametro path",
		)

		return
	}

	id, ok := parametros["id"]

	if !ok || strings.TrimSpace(id) == "" {

		fmt.Println(
			"ERROR: falta parametro id",
		)

		return
	}
//**--

fmt.Println()
fmt.Println("ID recibido:", id)

fmt.Println("PARTICIONES MONTADAS:")

for _, p := range estructuras.ParticionesMontadas {

	fmt.Println(
		"ID:", p.ID,
		"Nombre:", p.Name,
	)
}

fmt.Println()
//**--
	particion, existe :=
		BuscarParticionMontadaPorID(id)

	if !existe {

		fmt.Println(
			"ERROR: particion no montada",
		)

		return
	}

	switch strings.ToLower(name) {

	case "sb":

		ReporteSB(
			particion,
			path,
		)

	case "bm_inode":
		ReporteBMInode(
			particion,
			path,
		)

	case "bm_block":
		ReporteBMBlock(
			particion,
			path,
		)

	case "inode":
		ReporteINODE(
			particion,
			path,
		)


	case "block":
		ReporteBLOCK(
			particion,
			path,
		)
			
	case "mbr":
		ReporteMBR(
			particion,
			path,
		)

	case "disk":
		ReporteDISK(
			particion,
			path,
		)

	case "tree":
		ReporteTREE(
			particion,
			path,
		)
			


	default:

		fmt.Println(
			"ERROR: reporte no soportado",
		)
	}
}