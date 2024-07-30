package services

import (
	"bytes"
	"encoding/base64"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePDF(templateBytes []byte, data map[string]interface{}) (string, error) {
	// Parse the HTML template
	t, err := template.New("upload").Parse(string(templateBytes))
	if err != nil {
		return "", err
	}

	// Create a buffer to store the filled template
	var filledTemplate bytes.Buffer

	// Execute the template with the JSON data, storing the result in the buffer
	if err := t.Execute(&filledTemplate, data); err != nil {
		return "", err
	}

	// Initialize a new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return "", err
	}

	// Add a new page to the PDF generator with the filled template content
	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(filledTemplate.Bytes())))
	if err := pdfg.Create(); err != nil {
		return "", err
	}

	// Encode the PDF bytes to base64
	pdfBase64 := base64.StdEncoding.EncodeToString(pdfg.Bytes())
	return pdfBase64, nil
}
