package utilidades

import "strings"

func StringABytes4(texto string) [4]byte {
	var arreglo [4]byte
	copy(arreglo[:], texto)
	return arreglo
}

func StringABytes16(texto string) [16]byte {
	var arreglo [16]byte
	copy(arreglo[:], texto)
	return arreglo
}

func StringABytes20(texto string) [20]byte {
	var arreglo [20]byte
	copy(arreglo[:], texto)
	return arreglo
}

func BytesAString(datos []byte) string {
	return strings.TrimRight(string(datos), "\x00")
}