package renderer

import (
	"dondrozd/maker-gen/model"
	"embed"
	"fmt"
	"io"
	"log"
	"text/template"
)

//go:embed templates/*
var embeddedFS embed.FS

type MakerModel struct {
	PackageName string
	Imports     []model.ImportModel
	Structs     []model.StructModel
}

func RenderMaker(model model.GoFileModel, writer io.Writer) error {
	makerModel, err := mapToMaker(model)
	if err != nil {
		return err
	}
	// Parse all templates in the embedded `templates` directory
	tmpl, err := template.ParseFS(embeddedFS, "templates/*.gotmpl")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	// render maker
	err = tmpl.ExecuteTemplate(writer, "maker.gotmpl", makerModel)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	return nil
}

func mapToMaker(fileModel model.GoFileModel) (MakerModel, error) {
	imports := buildImports(fileModel)
	// only keep the targeted struct
	structs, err := filterStructByName(fileModel.Structs, "")
	if err != nil {
		return MakerModel{}, err
	}

	return MakerModel{
		PackageName: fileModel.PackageName,
		Imports:     imports,
		Structs:     structs,
	}, nil
}

func buildImports(fileModel model.GoFileModel) []model.ImportModel {
	imports := make([]model.ImportModel, len(fileModel.Imports)+1)
	// add existing imports from the file
	for importIdx, importModel := range fileModel.Imports {
		imports[importIdx] = importModel
	}
	// add import for the subject
	imports[len(imports)-1] = model.ImportModel{
		ImportPath: fileModel.ModulePath + "/" + fileModel.PackageName,
	}
	return imports
}

func filterStructByName(structModels []model.StructModel, structName string) ([]model.StructModel, error) {
	for _, structModel := range structModels {
		if structModel.Name == structName {
			return []model.StructModel{structModel}, nil
		}
	}

	return nil, fmt.Errorf("couldnt find struct with name: %s", structName)
}
