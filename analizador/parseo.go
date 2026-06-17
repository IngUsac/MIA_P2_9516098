package analizador

import (
	"regexp"
	"strings"
)

func ObtenerParametros(comando string) map[string]string {

	parametros := make(map[string]string)

	// Captura:
	// -clave=valor
	// -clave="valor con espacios"
	re := regexp.MustCompile(`-(\w+)=("[^"]*"|\S+)`)

	coincidencias := re.FindAllStringSubmatch(comando, -1)

	for _, match := range coincidencias {

		clave := strings.ToLower(match[1])

		valor := strings.Trim(match[2], "\"")

		parametros[clave] = valor
	}

	return parametros
}