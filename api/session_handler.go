package api


import (
	"encoding/json"
	"net/http"

	"MIA_P1_9516098/estructuras"
)

type SessionResponse struct {
	Logged   bool   `json:"logged"`
	User     string `json:"user"`
	ID       string `json:"id"`
	Partition string `json:"partition"`
}

func GetSessionHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	sesion := estructuras.SesionActual

	resp := SessionResponse{

		Logged: sesion.Activa,

		User: sesion.User,

		ID: sesion.ID,

		Partition: "",
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(resp)
}