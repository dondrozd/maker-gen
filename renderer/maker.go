package renderer

import (
	"embed"
	"io"
	"log"
	"text/template"

	"github.com/dondrozd/maker-gen/model"
)

//go:embed templates/*
var embeddedFS embed.FS

func RenderMaker(model model.MakerModel, writer io.Writer) error {
	// Parse all templates in the embedded `templates` directory
	tmpl, err := template.ParseFS(embeddedFS, "templates/*.gotmpl")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	// render maker
	err = tmpl.ExecuteTemplate(writer, "maker.gotmpl", model)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	return nil
}
