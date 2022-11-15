package entity

import (
	"strings"
)

type EntityName struct {
	name    string
	dirPath string
	basePkg string
}

func NewEntityName(name string, dirPath string, basePkg string) EntityName {
	return EntityName{
		name:    name,
		dirPath: dirPath,
		basePkg: basePkg,
	}
}

func (en EntityName) PascalCase() string {
	noUnderscore := strings.ReplaceAll(en.name, "_", " ")
	title := strings.Title(noUnderscore)
	return strings.ReplaceAll(title, " ", "")
}
