package utils_test

import (
	"testing"

	"github.com/eduardoths/micro-cli/utils"
)

func TestToSnakeCase(t *testing.T) {
	type testCase struct {
		it   string
		in   string
		want string
	}

	tc := []testCase{
		{
			it:   "should return the input if it's already snake case",
			in:   "test_snake_case.go",
			want: "test_snake_case.go",
		},
		{
			it:   "should convert camelCase to sake case",
			in:   "myTestSnakeCase.go",
			want: "my_test_snake_case.go",
		},
	}

	for _, c := range tc {
		t.Run(c.it, func(t *testing.T) {
			actual := utils.ToSnakeCase(c.in)
			if c.want != actual {
				t.Errorf("TestToSnakeCase failed.\nGot:\t\t%s\nwant:\t%s", actual, c.want)
				t.Logf("Case: %s", c.it)
			}
		})
	}
}
