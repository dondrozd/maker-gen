package model

type MakerModel struct {
	PackageName string
	Imports     []ImportModel
	Structs     []StructModel
}
