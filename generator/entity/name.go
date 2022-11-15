package entity

import (
	"strings"
	"unicode"
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

func (en EntityName) CamelCase() string {
	pascal := en.PascalCase()
	pascalRunes := []rune(pascal)
	camelRunes := make([]rune, 0, len(pascal))
	camelRunes = append(camelRunes, unicode.ToLower(pascalRunes[0]))
	camelRunes = append(camelRunes, pascalRunes[1:]...)
	return string(camelRunes)
}
