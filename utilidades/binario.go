package utilidades

import (
	"encoding/binary"
	"os"
)

// Obtiene el tamaño de una estructura en bytes
func ObtenerTamano(estructura interface{}) int {
	return binary.Size(estructura)
}

// Escribe una estructura en una posición específica del archivo
func EscribirObjeto(archivo *os.File, dato interface{}, posicion int64) error {

	_, err := archivo.Seek(posicion, 0)
	if err != nil {
		return err
	}

	return binary.Write(archivo, binary.LittleEndian, dato)
}

// Lee una estructura desde una posición específica del archivo
func LeerObjeto(archivo *os.File, dato interface{}, posicion int64) error {

	_, err := archivo.Seek(posicion, 0)
	if err != nil {
		return err
	}

	return binary.Read(archivo, binary.LittleEndian, dato)
}