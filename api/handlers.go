package api

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"

    "MIA_P1_9516098/analizador"
)

// APIResponse: Estructura estándar para todas las respuestas del Backend.

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

//CommandRequest: Estructura esperada desde el Frontend.
/*
Ejemplo:
	{
		"command":"mkdisk -size=10 -unit=M -fit=WF -path=\"./Disco1.dsk\""
	}
*/

type CommandRequest struct {
	Command string `json:"command"`
}

// StatusHandler: GET /api/status Permite verificar que el Backend está funcionando.

func StatusHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "Backend funcionando correctamente",
	})
}

// NotImplementedHandler: Handler temporal para endpoints aún no implementados.

func NotImplementedHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)

	json.NewEncoder(w).Encode(APIResponse{
		Success: false,
		Error:   "Endpoint aún no implementado",
	})
}



//Captura toda la salida enviada a stdout durante la ejecución de una función y la devuelve como string.

func captureOutput(f func()) string {

	old := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	f()

	w.Close()

	os.Stdout = old

	var buf bytes.Buffer

	io.Copy(&buf, r)

	return buf.String()
}


//GenericCommandHandler: Recibe cualquier comando mediante POST.
/*
Ejemplo:  POST /mkdisk
			{
				"command":"mkdisk -size=20 -unit=M ..."
			}
La ejecución real continúa realizándose mediante el analizador del Proyecto 1.
*/

func GenericCommandHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(
			w,
			"Metodo no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {

		http.Error(
			w,
			"No fue posible leer la petición",
			http.StatusBadRequest,
		)

		return
	}

	var request CommandRequest

	err = json.Unmarshal(body, &request)

	if err != nil {

		http.Error(
			w,
			"JSON inválido",
			http.StatusBadRequest,
		)

		return
	}

	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" COMANDO RECIBIDO DESDE API")
	fmt.Println(" ")
	fmt.Println(request.Command)
	fmt.Println(" ")

	// Ejecuta el comando utilizando el analizador existente.
	output := captureOutput(func() {
		analizador.Analizar(request.Command)
	})

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "Comando ejecutado correctamente",
		Data: map[string]string{
			"output": output,
		},
	})
	/*
	analizador.Analizar(request.Command)

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "Comando ejecutado correctamente",
	})*/
}