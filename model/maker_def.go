package model

type GenerateParams struct {
	SrcFile     string
	DestFile    string
	PackageName string
	StructName  string
	WithPrefix  string
}

type MakerModel struct {
	PackageName string
	Imports     []ImportModel
	Structs     []MakerStructModel
}

type MakerStructModel struct {
	Name       string
	WithPrefix string
	Properties []StructPropertyModel
}
