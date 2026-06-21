package analizador

import (
	
	"strings"
)
func ObtenerParametros(comando string) map[string]string {

	parametros := make(map[string]string)

	tokens := strings.Fields(comando)

	for _, token := range tokens {

		if !strings.HasPrefix(token, "-") {
			continue
		}

		token = strings.TrimPrefix(
			token,
			"-",
		)

		if strings.Contains(
			token,
			"=",
		) {

			partes := strings.SplitN(
				token,
				"=",
				2,
			)

			clave := strings.ToLower(
				partes[0],
			)

			valor := strings.Trim(
				partes[1],
				"\"",
			)

			parametros[clave] = valor

		} else {

			clave := strings.ToLower(
				token,
			)

			parametros[clave] = "true"
		}
	}

	return parametros
}