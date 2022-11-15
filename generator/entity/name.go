package entity

import (
	"strings"
	"unicode"

	"github.com/eduardoths/micro-cli/generator/file"
	"github.com/eduardoths/micro-cli/utils"
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

func (en EntityName) Type() string {
	return en.importName() + "." + en.PascalCase()
}

func (en EntityName) importName() string {
	fullPkg := utils.MergePaths(en.basePkg, en.dirPath)
	pkgDirs := strings.Split(strings.TrimRight(fullPkg, "/"), "/")
	lastDir := pkgDirs[len(pkgDirs)-1]
	return strings.ReplaceAll(lastDir, "_", "")
}

func (en EntityName) FileImport() file.Import {
	path := strings.TrimRight(utils.MergePaths(en.basePkg, en.dirPath), "/")

	importName := en.importName()
	if strings.HasSuffix(path, importName) {
		importName = ""
	}

	return file.Import{
		Path: path,
		Name: importName,
	}
}
