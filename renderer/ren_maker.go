package renderer

import (
	"embed"
	"fmt"
	"io"
	"text/template"

	"github.com/dondrozd/maker-gen/model"
)

//go:embed templates/*
var embeddedFS embed.FS

func RenderMaker(model model.MakerModel, writer io.Writer) error {
	// Parse all templates in the embedded `templates` directory
	tmpl, err := template.ParseFS(embeddedFS, "templates/*.gotmpl")
	if err != nil {
		return fmt.Errorf("Error parsing templates: %w", err)
	}
	// render maker
	err = tmpl.ExecuteTemplate(writer, "maker.gotmpl", model)
	if err != nil {
		return fmt.Errorf("Error executing template: %w", err)
	}

	return nil
}
