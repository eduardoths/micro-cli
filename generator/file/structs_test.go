package file_test

import (
	"reflect"
	"testing"

	"github.com/eduardoths/micro-cli/generator/file"
)

func TestNewStructFile(t *testing.T) {
	type testCase struct {
		it   string
		in   file.Struct
		want file.File
	}

	tc := []testCase{
		{
			it: "should generate a struct",
			in: file.Struct{
				Name: "example",
			},
			want: file.File{
				Package: file.STRUCTS_PKG,
				Structs: []file.Struct{
					{
						Name: "example",
					},
				},
			},
		},
	}

	for _, c := range tc {
		t.Run(c.it, func(t *testing.T) {
			actual := file.NewStructFile(c.in)
			if !reflect.DeepEqual(c.want, actual) {
				t.Errorf("Test: %s failed\nexpected:\t%v\ngot:\t\t%v",
					t.Name(),
					c.want,
					actual,
				)
			}
		})
	}
}
