package comandos

import (
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

// GenerarPDFTabla:
// Genera un PDF con formato tabular similar al enunciado.

func GenerarPDFTabla(
	ruta string,
	titulo string,
	datos [][]string,
) error {

	carpeta := filepath.Dir(
		ruta,
	)

	err := os.MkdirAll(
		carpeta,
		0755,
	)

	if err != nil {
		return err
	}

	pdf := gofpdf.New(
		"P",
		"mm",
		"A4",
		"",
	)

	pdf.AddPage()

	// Encabezado

	pdf.SetFont(
		"Arial",
		"B",
		14,
	)

	pdf.SetFillColor(
		0,
		100,
		0,
	)

	pdf.SetTextColor(
		255,
		255,
		255,
	)

	pdf.CellFormat(
		190,
		10,
		titulo,
		"1",
		1,
		"C",
		true,
		0,
		"",
	)

	// Restaurar color

	pdf.SetTextColor(
		0,
		0,
		0,
	)

	pdf.SetFont(
		"Arial",
		"",
		10,
	)

	for i, fila := range datos {

		if i%2 == 0 {

			pdf.SetFillColor(
				235,
				235,
				235,
			)

		} else {

			pdf.SetFillColor(
				50,
				180,
				90,
			)
		}

		pdf.CellFormat(
			95,
			8,
			fila[0],
			"1",
			0,
			"L",
			true,
			0,
			"",
		)

		pdf.CellFormat(
			95,
			8,
			fila[1],
			"1",
			1,
			"L",
			true,
			0,
			"",
		)
	}

	pdf.Ln(5)

	pdf.CellFormat(
		190,
		8,
		titulo,
		"",
		1,
		"C",
		false,
		0,
		"",
	)

	return pdf.OutputFileAndClose(
		ruta,
	)
}

  // GenerarPDFTexto:
// Genera un PDF para reportes de texto plano.

func GenerarPDFTexto(
	ruta string,
	titulo string,
	contenido string,
) error {

	carpeta := filepath.Dir(
		ruta,
	)

	err := os.MkdirAll(
		carpeta,
		0755,
	)

	if err != nil {
		return err
	}

	pdf := gofpdf.New(
		"P",
		"mm",
		"A4",
		"",
	)

	pdf.AddPage()

	pdf.SetFont(
		"Arial",
		"B",
		14,
	)

	pdf.CellFormat(
		190,
		10,
		titulo,
		"",
		1,
		"C",
		false,
		0,
		"",
	)

	pdf.Ln(5)

	pdf.SetFont(
		"Courier",
		"",
		8,
	)

	pdf.MultiCell(
		0,
		4,
		contenido,
		"",
		"",
		false,
	)

	return pdf.OutputFileAndClose(
		ruta,
	)
}