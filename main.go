package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"gen", "g"},
				Usage:   "generate",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					findStructs("examples/example_1.go")

					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("completed task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("new task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type StructModel struct {
	Name       string
	Properties []StructPropertyModel
}
type StructPropertyModel struct {
	Name string
	Type string
}

func findStructs(fileName string) {
	// Open the Go file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	myStructs := []StructModel{}
	// Walk through the AST and find struct types
	ast.Inspect(node, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if ok {
			if structType, isStruct := typeSpec.Type.(*ast.StructType); isStruct {
				fmt.Println("Struct found:", typeSpec.Name.Name)
				structModel := StructModel{Name: typeSpec.Name.Name}

				// Iterate over the fields in the struct
				for _, field := range structType.Fields.List {
					for _, fieldName := range field.Names {
						fmt.Printf("  Field: %s, Type: %s\n", fieldName.Name, field.Type)
						propModel := StructPropertyModel{
							Name: fieldName.Name,
							Type: typeToString(field.Type),
						}
						structModel.Properties = append(structModel.Properties, propModel)
					}
				}
				fmt.Println()
				myStructs = append(myStructs, structModel)
			}
		}
		return true
	})
	fmt.Printf("struct defs: %s", myStructs)
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
		return fmt.Sprintf("%s.%s", typeToString(t.X), t.Sel.Name) // Qualified types like "pkg.Type"
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", typeToString(t.Key), typeToString(t.Value)) // Map types
	// Add more cases as needed for other complex types.
	default:
		return fmt.Sprintf("%T", expr) // Fallback to the type's Go syntax for unknown cases
	}
}
