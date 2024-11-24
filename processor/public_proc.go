package processor

import (
	"fmt"
	"unicode"

	"github.com/dondrozd/maker-gen/model"
)

func PublicProc(fileModel model.GoFileModel, structName string) (model.MakerModel, error) {
	imports := buildImports(fileModel)
	// only keep the targeted struct
	structs, err := filterStructByName(fileModel.Structs, structName)
	if err != nil {
		return model.MakerModel{}, err
	}

	return model.MakerModel{
		PackageName: fileModel.PackageName,
		Imports:     imports,
		Structs:     structs,
	}, nil
}

func buildImports(fileModel model.GoFileModel) []model.ImportModel {
	imports := make([]model.ImportModel, len(fileModel.Imports)+1)
	// add existing imports from the file
	copy(imports, fileModel.Imports[:])
	// add import for the subject
	imports[len(imports)-1] = model.ImportModel{
		ImportPath: fmt.Sprintf("\"%s/%s\"", fileModel.ModulePath, fileModel.PackageName),
	}
	return imports
}

func filterStructByName(structModels []model.StructModel, structName string) ([]model.StructModel, error) {
	for _, structModel := range structModels {
		if structModel.Name == structName {
			return []model.StructModel{
				mapStruct(structModel),
			}, nil
		}
	}

	return nil, fmt.Errorf("could not find struct with name: %s", structName)
}

func mapStruct(sructModel model.StructModel) model.StructModel {
	return model.StructModel{
		Name:       sructModel.Name,
		Properties: mapPublicProperties(sructModel.Properties),
	}
}

func mapPublicProperties(structPropertyModel []model.StructPropertyModel) []model.StructPropertyModel {
	props := []model.StructPropertyModel{}
	for _, prop := range structPropertyModel {
		if isPublic(prop.Name) {
			props = append(props, prop)
		}
	}

	return props
}

func isPublic(s string) bool {
	firstChar := rune(s[0]) // Convert the first character to a rune
	return unicode.IsUpper(firstChar)
}
