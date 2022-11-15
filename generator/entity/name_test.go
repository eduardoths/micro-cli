package entity_test

import (
	"testing"

	"github.com/eduardoths/micro-cli/generator/entity"
	"github.com/eduardoths/micro-cli/generator/file"
	"github.com/eduardoths/micro-cli/tests/utils"
)

func TestEntityName_PascalCase(t *testing.T) {
	type testCase struct {
		it   string
		in   entity.EntityName
		want string
	}

	tc := []testCase{
		{
			it:   "should return PascalCase names as pascal case",
			in:   entity.NewEntityName("PascalCase", "", ""),
			want: "PascalCase",
		},
		{
			it:   "should return camelCase name as PascalCase",
			in:   entity.NewEntityName("camelCase", "", ""),
			want: "CamelCase",
		},
		{
			it:   "should return snake_case name as PascalCase",
			in:   entity.NewEntityName("snake_case", "", ""),
			want: "SnakeCase",
		},
	}

	for _, c := range tc {
		t.Run(c.it, func(t *testing.T) {
			actual := c.in.PascalCase()
			if c.want != actual {
				utils.Error(t, c.want, actual)
			}
		})
	}
}

func TestEntityName_CamelCase(t *testing.T) {
	type testCase struct {
		it   string
		in   entity.EntityName
		want string
	}

	tc := []testCase{
		{
			it:   "should return PascalCase names as camelCase",
			in:   entity.NewEntityName("PascalCase", "", ""),
			want: "pascalCase",
		},
		{
			it:   "should return camelCase name as camelCase",
			in:   entity.NewEntityName("camelCase", "", ""),
			want: "camelCase",
		},
		{
			it:   "should return snake_case name as camelCase",
			in:   entity.NewEntityName("snake_case", "", ""),
			want: "snakeCase",
		},
	}

	for _, c := range tc {
		t.Run(c.it, func(t *testing.T) {
			actual := c.in.CamelCase()
			if c.want != actual {
				utils.Error(t, c.want, actual)
			}
		})
	}
}

func TestEntityName_Type(t *testing.T) {
	type testCase struct {
		it   string
		in   entity.EntityName
		want string
	}

	tc := []testCase{
		{
			it:   "should return structs.XptoStruct",
			in:   entity.NewEntityName("xptoStruct", "structs", ""),
			want: "structs.XptoStruct",
		},
		{
			it:   "should remove trailing '/' from pkg",
			in:   entity.NewEntityName("xptoStruct", "xpto/structs", ""),
			want: "structs.XptoStruct",
		},
		{
			it:   "should use base pkg if dirPath is empty",
			in:   entity.NewEntityName("Struct", "", "github.com/eduardoths/xpto"),
			want: "xpto.Struct",
		},
		{
			it:   "should convert entity name from snake case to pascal case",
			in:   entity.NewEntityName("xpto_struct", "", "github.com/eduardoths/xpto"),
			want: "xpto.XptoStruct",
		},
		{
			it:   "should remove underscores before pkg name",
			in:   entity.NewEntityName("Struct", "", "github.com/eduardoths/xpto_pkg"),
			want: "xptopkg.Struct",
		},
	}

	for _, c := range tc {
		t.Run(c.it, func(t *testing.T) {
			actual := c.in.Type()
			if c.want != actual {
				utils.Error(t, c.want, actual)
			}
		})
	}
}

func TestEntityName_FileImport(t *testing.T) {
	type testCase struct {
		it   string
		in   entity.EntityName
		want file.Import
	}

	tc := []testCase{
		{
			it: "should return a file import for structs pkg",
			in: entity.NewEntityName("Xpto", "src/structs", "github.com/eduardoths/microservice"),
			want: file.Import{
				Path: "github.com/eduardoths/microservice/src/structs",
				Name: "",
			},
		},
		{
			it: "should return a file import for structs pkg and remove trailing '/'",
			in: entity.NewEntityName("Xpto", "src/structs/", "github.com/eduardoths/microservice"),
			want: file.Import{
				Path: "github.com/eduardoths/microservice/src/structs",
				Name: "",
			},
		},
		{
			it: "should have an import name if it ends with snake_case path",
			in: entity.NewEntityName("XptoStructRepository", "src/repositories/xpto_struct", "github.com/eduardoths/microservice"),
			want: file.Import{
				Path: "github.com/eduardoths/microservice/src/repositories/xpto_struct",
				Name: "xptostruct",
			},
		},
	}

	for _, c := range tc {
		t.Run(c.it, func(t *testing.T) {
			actual := c.in.FileImport()
			if c.want != actual {
				utils.Error(t, c.want, actual)
			}
		})
	}
}

func TestEntityName_FilePath(t *testing.T) {
	type testCase struct {
		it   string
		in   entity.EntityName
		want string
	}

	tc := []testCase{
		{
			it:   "should return file path as pascal case",
			in:   entity.NewEntityName("PascalCase", "", ""),
			want: "./pascal_case.go",
		},
		{
			it:   "should ignore base pkg",
			in:   entity.NewEntityName("PascalCase", "", "github.com/eduardoths/"),
			want: "./pascal_case.go",
		},
		{
			it:   "should have the correct file path",
			in:   entity.NewEntityName("PascalCase", "src/structs/", "github.com/eduardoths/"),
			want: "src/structs/pascal_case.go",
		},
	}

	for _, c := range tc {
		t.Run(c.it, func(t *testing.T) {
			actual := c.in.FilePath()
			if c.want != actual {
				utils.Error(t, c.want, actual)
			}
		})
	}
}

func TestEntityName_Alias(t *testing.T) {
	type testCase struct {
		it   string
		in   entity.EntityName
		want string
	}

	tc := []testCase{
		{
			it:   "should return a simples alias",
			in:   entity.NewEntityName("Struct", "", ""),
			want: "s",
		},
		{
			it:   "should return correct alias",
			in:   entity.NewEntityName("an_example_struct_repository", "", ""),
			want: "aesr",
		},
	}

	for _, c := range tc {
		t.Run(c.it, func(t *testing.T) {
			actual := c.in.Alias()
			if c.want != actual {
				utils.Error(t, c.want, actual)
			}
		})
	}
}
