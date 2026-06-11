package utilidades

func StringABytes20(texto string) [20]byte {

	var arreglo [20]byte

	copy(arreglo[:], texto)

	return arreglo
}