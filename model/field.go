package model

type Field struct {
	Name            string
	Type            string
	Filter          bool
	IdField         bool
	CreateValidator string
	FilterValidator string
}
