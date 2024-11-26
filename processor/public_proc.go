package processor

import (
	"fmt"
	"log/slog"
	"unicode"

	"github.com/dondrozd/maker-gen/model"
)

func PublicProc(fileModel model.GoFileModel, cmd model.GenerateParams) (model.MakerModel, error) {
	slog.Info("PublicProc called")
	imports := buildImports(fileModel)
	// only keep the targeted struct
	structs, err := filterStructByName(fileModel.Structs, cmd)
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

func filterStructByName(structModels []model.StructModel, cmd model.GenerateParams) ([]model.MakerStructModel, error) {
	slog.Info("filterStructByName:" + cmd.StructName)
	for _, structModel := range structModels {
		// support * in the future
		if structModel.Name == cmd.StructName {
			return []model.MakerStructModel{
				mapStruct(structModel, cmd),
			}, nil
		}
	}

	return nil, fmt.Errorf("could not find struct with name: %s", cmd.StructName)
}

func mapStruct(sructModel model.StructModel, cmd model.GenerateParams) model.MakerStructModel {
	return model.MakerStructModel{
		Name:       sructModel.Name,
		WithPrefix: cmd.WithPrefix,
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
	return unicode.IsUpper(rune(s[0]))
}
