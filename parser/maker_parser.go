package parsers

import (
	"bufio"
	"dondrozd/maker-gen/model"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func MakerParse(fileName string) (model.GoFileModel, error) {
	targetFile := model.GoFileModel{
		Name: fileName,
	}
	// Open the Go file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, nil, parser.AllErrors)
	if err != nil {
		return targetFile, fmt.Errorf("error parsing file: %w", err)
	}
	// add package name from the file
	targetFile.PackageName = node.Name.Name
	// add module path
	targetFile.ModulePath, err = getModulePath()
	if err != nil {
		return targetFile, fmt.Errorf("error getting module path: %w", err)
	}
	// Discover and print imports
	fmt.Println("Imports:")
	for _, imp := range node.Imports {
		importPath := imp.Path.Value // Import path is quoted, e.g., `"fmt"`
		if imp.Name != nil {
			fmt.Printf("  %s %s\n", imp.Name.Name, importPath) // Alias imports
			targetFile.Imports = append(targetFile.Imports, model.ImportModel{
				Alias:      imp.Name.Name,
				ImportPath: importPath,
			})
		} else {
			fmt.Printf("  %s\n", importPath) // Regular imports
			targetFile.Imports = append(targetFile.Imports, model.ImportModel{
				ImportPath: importPath,
			})
		}
	}
	// Walk through the AST and find struct types
	ast.Inspect(node, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if ok {
			if structType, isStruct := typeSpec.Type.(*ast.StructType); isStruct {
				fmt.Println("Struct found:", typeSpec.Name.Name)
				structModel := model.StructModel{Name: typeSpec.Name.Name}

				// Iterate over the fields in the struct
				for _, field := range structType.Fields.List {
					for _, fieldName := range field.Names {
						fmt.Printf("  Field: %s, Type: %s\n", fieldName.Name, field.Type)
						propModel := model.StructPropertyModel{
							Name: fieldName.Name,
							Type: typeToString(field.Type),
						}
						structModel.Properties = append(structModel.Properties, propModel)
					}
				}
				fmt.Println()
				targetFile.Structs = append(targetFile.Structs, structModel)
			}
		}
		return true
	})

	return targetFile, nil
}

func typeToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name // Basic types like int, string, etc.
	case *ast.StarExpr:
		return "*" + typeToString(t.X) // Pointer types
	case *ast.ArrayType:
		return "[]" + typeToString(t.Elt) // Array types
	case *ast.SelectorExpr:
		fmt.Println("SelectorExpr")
		return fmt.Sprintf("%s.%s", typeToString(t.X), t.Sel.Name) // Qualified types like "pkg.Type"
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", typeToString(t.Key), typeToString(t.Value)) // Map types
		// Add more cases as needed for other complex types.
	default:
		return fmt.Sprintf("%T", expr) // Fallback to the type's Go syntax for unknown cases
	}
}

func getModulePath() (string, error) {
	modulePath, err := findGoMod()
	if err != nil {
		return "", err
	}

	file, err := os.Open(modulePath)
	if err != nil {
		return "", fmt.Errorf("could not open go.mod: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	return "", fmt.Errorf("module path not found in go.mod")
}

func findGoMod() (string, error) {
	// Start from the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get current directory: %w", err)
	}

	for {
		// Construct the path to go.mod in the current directory
		goModPath := filepath.Join(dir, "go.mod")

		// Check if go.mod exists in this directory
		if _, err := os.Stat(goModPath); err == nil {
			return goModPath, nil // go.mod found
		}

		// Move up to the parent directory
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// If we've reached the root directory, stop searching
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("go.mod not found in any parent directory")
}
