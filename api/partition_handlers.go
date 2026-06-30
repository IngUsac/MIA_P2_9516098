package api

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"MIA_P1_9516098/comandos"
)

/*
PartitionInfo

Información que será enviada al Frontend para representar
una partición de manera amigable.
*/
type PartitionInfo struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	TypeName string `json:"typeName"`

	Fit     string `json:"fit"`
	FitName string `json:"fitName"`

	Start int32 `json:"start"`
	Size  int32 `json:"size"`

	Status  string `json:"status"`
	Mounted bool   `json:"mounted"`

	ID string `json:"id"`
}

/*
GetPartitionsHandler

Obtiene todas las particiones existentes dentro de un disco
y las devuelve en formato JSON para ser consumidas por
el Frontend.
*/
func GetPartitionsHandler(w http.ResponseWriter, r *http.Request) {

	disk := r.URL.Query().Get("disk")

	if disk == "" {

		http.Error(
			w,
			"Falta parámetro disk",
			http.StatusBadRequest,
		)

		return
	}

	path := "./SALIDAS/discos/" + disk

	archivo, err := os.Open(path)

	if err != nil {

		http.Error(
			w,
			"No fue posible abrir el disco",
			http.StatusInternalServerError,
		)

		return
	}

	defer archivo.Close()

	mbr, err := comandos.LeerMBR(archivo)

	if err != nil {

		http.Error(
			w,
			"No fue posible leer el MBR",
			http.StatusInternalServerError,
		)

		return
	}

	var partitions []PartitionInfo

	for _, part := range mbr.MbrPartitions {

		// Ignorar entradas vacías.
		if part.PartSize == 0 {

			continue
		}

		//---------------------------------------
		// Tipo de partición
		//---------------------------------------

		typeName := "Primaria"

		switch part.PartType {

		case 'E':
			typeName = "Extendida"

		case 'L':
			typeName = "Lógica"
		}

		//---------------------------------------
		// Tipo de ajuste
		//---------------------------------------

		fitName := ""

		switch part.PartFit {

		case 'B':
			fitName = "Best Fit"

		case 'F':
			fitName = "First Fit"

		case 'W':
			fitName = "Worst Fit"
		}

		//---------------------------------------
		// Estado de montaje
		//---------------------------------------

		mounted := part.PartStatus == 'M'

		partitions = append(partitions, PartitionInfo{

			Name: strings.TrimRight(
				string(part.PartName[:]),
				"\x00",
			),

			Type: string(part.PartType),

			TypeName: typeName,

			Fit: string(part.PartFit),

			FitName: fitName,

			Start: part.PartStart,

			Size: part.PartSize,

			Status: string(part.PartStatus),

			Mounted: mounted,

			ID: strings.TrimRight(
				string(part.PartID[:]),
				"\x00",
			),
		})
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(partitions)
}