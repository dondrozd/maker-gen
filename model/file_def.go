package model

type GoFileModel struct {
	Name        string
	PackageName string
	ModulePath  string
	Imports     []ImportModel
	Structs     []StructModel
}

type ImportModel struct {
	Alias      string
	ImportPath string
}

type StructModel struct {
	Name       string
	Properties []StructPropertyModel
}

type StructPropertyModel struct {
	Name string
	Type string
}
