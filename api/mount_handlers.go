package api

import (
	"encoding/json"
	"net/http"

	"MIA_P1_9516098/estructuras"
)

type MountInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetMountsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	mounts := []MountInfo{}

	for _, m := range estructuras.ParticionesMontadas {

		mounts = append(
			mounts,
			MountInfo{
				ID:   m.ID,
				Name: m.Name,
			},
		)

	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		mounts,
	)

}