package entity

import (
	"github.com/eduardoths/micro-cli/generator/file"
	"github.com/eduardoths/micro-cli/utils"
)

type Repository struct {
	repoName   EntityName
	structName EntityName
	implStruct file.Struct

	Interface file.Interface
	Imports   file.Imports
}

type imethod struct {
	method         file.Method
	imports        file.Imports
	implementation []string
}

func NewRepository(structName EntityName, basePkg string) Repository {
	repo := Repository{
		repoName: NewEntityName(
			structName.PascalCase()+"Repository",
			utils.MergePaths(REPOSITORIES_PATH, structName.SnakeCase()),
			basePkg,
		),
		structName: structName,
	}
	repo.build()

	return repo
}

func (r *Repository) build() {
	r.buildInterface()
	r.buildImports()
	r.buildImplementation()
}

func (r *Repository) File() file.File {
	return file.File{
		Package: r.repoName.ImportName(),
		Imports: r.Imports,
		Structs: []file.Struct{r.implStruct},
	}

}

func (r *Repository) buildInterface() {
	internalMethods := r.internalMethods()
	r.Interface = file.Interface{
		Name:    r.repoName.PascalCase(),
		Methods: []file.Method{},
	}

	for _, imethod := range internalMethods {
		r.Interface.Methods = append(r.Interface.Methods, imethod.method)
	}
}

func (r *Repository) buildImports() {
	internalMethods := r.internalMethods()

	r.Imports = make(file.Imports, 0)
	for _, imethod := range internalMethods {
		r.Imports = append(r.Imports, imethod.imports...)
	}
}

func (r *Repository) buildImplementation() {
	getAll := r.getAllMethod()
	get := r.getMethod()
	r.implStruct = file.Struct{
		Name: r.repoName.PascalCase(),
		Implementations: []file.Implementation{
			{
				StructAlias: r.repoName.Alias(),
				StructName:  r.repoName.PascalCase(),
				Func:        getAll.method,
				CodeLines:   getAll.implementation,
			},
			{
				StructAlias: r.repoName.Alias(),
				StructName:  r.repoName.PascalCase(),
				Func:        get.method,
				CodeLines:   get.implementation,
			},
		},
	}
}

func (r Repository) internalMethods() []imethod {
	return []imethod{
		r.getAllMethod(),
		r.getMethod(),
	}
}

func (r Repository) getAllMethod() imethod {
	return imethod{
		method: file.Method{
			Name: "GetAll",
			Params: file.Args{
				{
					Name: "ctx",
					Type: CONTEXT_TYPE,
				},
			},
			Results: file.Args{
				{
					Name: r.structName.CamelCase(),
					Type: "[]" + r.structName.Type(),
				},
				{
					Name: "err",
					Type: "error",
				},
			},
		},
		imports: file.Imports{
			{Path: CONTEXT_PKG},
			r.structName.FileImport(),
		},
	}
}

func (r Repository) getMethod() imethod {
	return imethod{
		method: file.Method{
			Name: "Get",
			Params: file.Args{
				{
					Name: "ctx",
					Type: CONTEXT_TYPE,
				},
				{
					Name: "id",
					Type: ID_TYPE,
				},
			},
			Results: file.Args{
				{
					Name: r.structName.CamelCase(),
					Type: r.structName.Type(),
				},
				{
					Name: "err",
					Type: "error",
				},
			},
		},
		imports: file.Imports{
			{Path: CONTEXT_PKG},
			{Path: ID_PKG},
			r.structName.FileImport(),
		},
	}
}
