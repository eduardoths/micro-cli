package entity_test

import (
	"testing"

	"github.com/eduardoths/micro-cli/generator/entity"
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
