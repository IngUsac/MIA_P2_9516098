package analizador

import (
	"strings"
)

func ObtenerParametros(comando string) map[string]string {

	parametros := make(map[string]string)

	tokens := dividirTokens(comando)

	for _, token := range tokens {

		if !strings.HasPrefix(token, "-") {
			continue
		}

		token = strings.TrimPrefix(token, "-")

		if strings.Contains(token, "=") {

			partes := strings.SplitN(
				token,
				"=",
				2,
			)

			clave := strings.ToLower(partes[0])

			valor := strings.Trim(
				partes[1],
				"\"",
			)

			parametros[clave] = valor

		} else {

			parametros[strings.ToLower(token)] = "true"
		}
	}

	return parametros
}

func dividirTokens(comando string) []string {

	var tokens []string
	var actual strings.Builder

	enComillas := false

	for _, c := range comando {

		switch c {

		case '"':

			enComillas = !enComillas
			actual.WriteRune(c)

		case ' ', '\t':

			if enComillas {

				actual.WriteRune(c)

			} else {

				if actual.Len() > 0 {

					tokens = append(
						tokens,
						actual.String(),
					)

					actual.Reset()
				}
			}

		default:

			actual.WriteRune(c)
		}
	}

	if actual.Len() > 0 {

		tokens = append(
			tokens,
			actual.String(),
		)
	}

	return tokens
}

/*
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

*/