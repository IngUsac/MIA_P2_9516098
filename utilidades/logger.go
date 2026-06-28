package utilidades

import (
	"bytes"
	"fmt"
)

/*
Logger central del proyecto.

Toda la salida del sistema deberá pasar por este logger.
Esto permitirá enviarla posteriormente al Frontend.
*/

var Output bytes.Buffer

// Limpia el buffer antes de ejecutar un comando.
func ClearOutput() {
	Output.Reset()
}

// Devuelve toda la salida generada.
func GetOutput() string {
	return Output.String()
}

// Escribe un mensaje tanto en consola como en el buffer.
func Log(a ...interface{}) {

	fmt.Println(a...)

	fmt.Fprintln(&Output, a...)
}

// Igual que fmt.Printf.
func Logf(format string, a ...interface{}) {

	fmt.Printf(format, a...)

	fmt.Fprintf(&Output, format, a...)
}