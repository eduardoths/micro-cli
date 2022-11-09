package generate_test

import (
	"fmt"
	"testing"

	"github.com/eduardoths/micro-cli/services/generate"
)

func TestFile_String(t *testing.T) {
	type testCase struct {
		it   string
		file generate.File
		want string
	}

	tc := []testCase{
		{
			it: "should return a file that has only a package",
			file: generate.File{
				Package: "teste",
			},
			want: "package teste\n\n",
		},
		{
			it: "should return a file with one import",
			file: generate.File{
				Package: "test",
				Imports: generate.Imports{
					generate.Import{Path: "strings"},
				},
			},
			want: "package test\n\n" +
				`import "strings"` + "\n\n",
		},
		{
			it: "should return a file with one import with alias",
			file: generate.File{
				Package: "test",
				Imports: generate.Imports{
					generate.Import{Path: "strings", Name: "str"},
				},
			},
			want: "package test\n\n" +
				`import str "strings"` + "\n\n",
		},
		{
			it: "should return a file with two imports",
			file: generate.File{
				Package: "test",
				Imports: generate.Imports{
					generate.Import{Path: "strings", Name: "str"},
					generate.Import{Path: "errors"},
				},
			},
			want: "package test\n\n" +
				"import (\n" +
				fmt.Sprintf("\tstr %s\n", `"strings"`) +
				fmt.Sprintf("\t%s\n", `"errors"`) +
				")\n\n",
		},
		{
			it: "should return a file with an empty interface",
			file: generate.File{
				Package: "test",
				Interfaces: []generate.Interface{
					{
						Name: "Xpto",
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {}\n\n",
		},
		{
			it: "should return a file with an interface containing one method without params or results",
			file: generate.File{
				Package: "test",
				Interfaces: []generate.Interface{
					{
						Name: "Xpto",
						Methods: []generate.Method{
							{Name: "Create"},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {\n" +
				"\tCreate()\n" +
				"}\n\n",
		},
		{
			it: "should return a file with an interface containing one method with one named param",
			file: generate.File{
				Package: "test",
				Interfaces: []generate.Interface{
					{
						Name: "Xpto",
						Methods: []generate.Method{
							{
								Name: "Create",
								Params: generate.Args{
									{
										Name: "s",
										Type: "string",
									},
								},
							},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {\n" +
				"\tCreate(s string)\n" +
				"}\n\n",
		},
		{
			it: "should return a file with an interface containing one method with one unnamed param",
			file: generate.File{
				Package: "test",
				Interfaces: []generate.Interface{
					{
						Name: "Xpto",
						Methods: []generate.Method{
							{
								Name: "Create",
								Params: generate.Args{
									{Type: "string"},
								},
							},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {\n" +
				"\tCreate(string)\n" +
				"}\n\n",
		},
		{
			it: "should return a file with an interface containing one method with two params",
			file: generate.File{
				Package: "test",
				Interfaces: []generate.Interface{
					{
						Name: "Xpto",
						Methods: []generate.Method{
							{
								Name: "Create",
								Params: generate.Args{
									{Name: "s", Type: "string"},
									{Name: "i", Type: "int"},
								},
							},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {\n" +
				"\tCreate(s string, i int)\n" +
				"}\n\n",
		},
		{
			it: "should return a file with an interface containing one method with one named return",
			file: generate.File{
				Package: "test",
				Interfaces: []generate.Interface{
					{
						Name: "Xpto",
						Methods: []generate.Method{
							{
								Name: "Create",
								Results: generate.Args{
									{Name: "err", Type: "error"},
								},
							},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {\n" +
				"\tCreate() (err error)\n" +
				"}\n\n",
		},
		{
			it: "should return a file with an interface containing one method with one unnamed return",
			file: generate.File{
				Package: "test",
				Interfaces: []generate.Interface{
					{
						Name: "Xpto",
						Methods: []generate.Method{
							{
								Name: "Create",
								Results: generate.Args{
									{Type: "error"},
								},
							},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {\n" +
				"\tCreate() error\n" +
				"}\n\n",
		},
		{
			it: "should return a file with an interface containing one method with two unnamed returns",
			file: generate.File{
				Package: "test",
				Interfaces: []generate.Interface{
					{
						Name: "Xpto",
						Methods: []generate.Method{
							{
								Name: "Create",
								Results: generate.Args{
									{Type: "bool"},
									{Type: "error"},
								},
							},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {\n" +
				"\tCreate() (bool, error)\n" +
				"}\n\n",
		},
		{
			it: "should return a file with an interface containing one method with three named returns",
			file: generate.File{
				Package: "test",
				Interfaces: []generate.Interface{
					{
						Name: "Xpto",
						Methods: []generate.Method{
							{
								Name: "Create",
								Results: generate.Args{
									{Type: "bool", Name: "ok"},
									{Type: "int", Name: "n"},
									{Type: "error", Name: "err"},
								},
							},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {\n" +
				"\tCreate() (ok bool, n int, err error)\n" +
				"}\n\n",
		},
		{
			it: "should return a file with an interface containing three methods",
			file: generate.File{
				Package: "test",
				Interfaces: []generate.Interface{
					{
						Name: "Xpto",
						Methods: []generate.Method{
							{
								Name: "Create",
								Params: generate.Args{
									{Type: "structs.Example", Name: "xpto1"},
									{Type: "string", Name: "name"},
									{Type: "int", Name: "i"},
								},
								Results: generate.Args{
									{Type: "bool", Name: "ok"},
									{Type: "int", Name: "n"},
									{Type: "error", Name: "err"},
								},
							},
							{
								Name: "Update",
								Params: generate.Args{
									{Type: "structs.Example", Name: "xpto1"},
									{Type: "string", Name: "name"},
									{Type: "int", Name: "i"},
								},
								Results: generate.Args{
									{Type: "bool", Name: "ok"},
									{Type: "int", Name: "n"},
									{Type: "error", Name: "err"},
								},
							},
							{
								Name: "Delete",
								Params: generate.Args{
									{Type: "structs.Example", Name: "xpto1"},
									{Type: "string", Name: "name"},
									{Type: "int", Name: "i"},
								},
								Results: generate.Args{
									{Type: "bool", Name: "ok"},
									{Type: "int", Name: "n"},
									{Type: "error", Name: "err"},
								},
							},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {\n" +
				"\tCreate(xpto1 structs.Example, name string, i int) (ok bool, n int, err error)\n" +
				"\tUpdate(xpto1 structs.Example, name string, i int) (ok bool, n int, err error)\n" +
				"\tDelete(xpto1 structs.Example, name string, i int) (ok bool, n int, err error)\n" +
				"}\n\n",
		},
		{
			it: "should return a file with two interfaces",
			file: generate.File{
				Package: "test",
				Interfaces: []generate.Interface{
					{Name: "XptoOne"},
					{Name: "XptoTwo"},
				},
			},
			want: "package test\n\n" +
				"type XptoOne interface {}\n\n" +
				"type XptoTwo interface {}\n\n",
		},
		{
			it: "should return a file with an empty struct",
			file: generate.File{
				Package: "test",
				Structs: []generate.Struct{
					{Name: "Xpto"},
				},
			},
			want: "package test\n\n" +
				"type Xpto struct {}\n\n",
		},
		{
			it: "should return a file with one field",
			file: generate.File{
				Package: "test",
				Structs: []generate.Struct{
					{
						Name: "Xpto",
						Fields: []generate.Field{
							{Name: "Field"},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto struct {\n" +
				"\tField\n" +
				"}\n\n",
		},
		{
			it: "should return a file with one field with type",
			file: generate.File{
				Package: "test",
				Structs: []generate.Struct{
					{
						Name: "Xpto",
						Fields: []generate.Field{
							{Name: "Field", Type: "string"},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto struct {\n" +
				"\tField string\n" +
				"}\n\n",
		},
		{
			it: "should return a file with one field with type and tag",
			file: generate.File{
				Package: "test",
				Structs: []generate.Struct{
					{
						Name: "Xpto",
						Fields: []generate.Field{
							{Name: "Field", Type: "string", Tag: "`json:\"-\"`"},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto struct {\n" +
				"\tField string `json:\"-\"`\n" +
				"}\n\n",
		},
		{
			it: "should return a file with three fields",
			file: generate.File{
				Package: "structs",
				Structs: []generate.Struct{
					{
						Name: "Xpto",
						Fields: []generate.Field{
							{Name: "Str", Type: "string", Tag: "`json:\"-\"`"},
							{Name: "Int", Type: "int"},
							{Name: "pkg.Field"},
						},
					},
				},
			},
			want: "package structs\n\n" +
				"type Xpto struct {\n" +
				"\tStr string `json:\"-\"`\n" +
				"\tInt int\n" +
				"\tpkg.Field\n" +
				"}\n\n",
		},
		{
			it: "should return a file with two empty structs",
			file: generate.File{
				Package: "test",
				Structs: []generate.Struct{
					{Name: "Xpto"},
					{Name: "XptoAgain"},
				},
			},
			want: "package test\n\n" +
				"type Xpto struct {}\n\n" +
				"type XptoAgain struct {}\n\n",
		},
		{
			it: "should return a complete file",
			file: generate.File{
				Package: "complextest",
				Imports: []generate.Import{
					{Path: "github.com/eduardoths/my_structs/structs"},
					{Name: "mystructs", Path: "github.com/eduardoths/my_structs"},
				},
				Interfaces: []generate.Interface{
					{
						Name: "Xpto",
						Methods: []generate.Method{
							{
								Name: "Done",
								Params: generate.Args{
									{Name: "s", Type: "structs.Struct"},
								},
								Results: generate.Args{{Type: "error"}},
							},
							{
								Name:    "String",
								Results: generate.Args{{Type: "string"}},
							},
						},
					},
					{
						Name: "Err",
						Methods: []generate.Method{
							{
								Name:    "Error",
								Results: generate.Args{{Type: "string"}},
							},
						},
					},
					{Name: "empty"},
				},
				Structs: []generate.Struct{
					{
						Name: "Foo",
						Fields: []generate.Field{
							{
								Name: "mystructs.Foo",
							},
						},
					},
					{
						Name: "Bar",
						Fields: []generate.Field{
							{
								Name: "Ok",
								Type: "bool",
							},
						},
					},
					{Name: "emptyStruct"},
				},
			},
			want: "package complextest\n\n" +
				"import (\n" +
				"\t\"github.com/eduardoths/my_structs/structs\"\n" +
				"\tmystructs \"github.com/eduardoths/my_structs\"\n" +
				")\n\n" +
				"type Xpto interface {\n" +
				"\tDone(s structs.Struct) error\n" +
				"\tString() string\n" +
				"}\n\n" +
				"type Err interface {\n" +
				"\tError() string\n" +
				"}\n\n" +
				"type empty interface {}\n\n" +
				"type Foo struct {\n" +
				"\tmystructs.Foo\n" +
				"}\n\n" +
				"type Bar struct {\n" +
				"\tOk bool\n" +
				"}\n\n" +
				"type emptyStruct struct {}\n\n",
		},
	}

	for _, c := range tc {
		t.Run(c.it, func(t *testing.T) {
			actual := c.file.String()
			if c.want != actual {
				t.Errorf("File.String() failed, want: \n%s\ngot\n%s", c.want, actual)
				t.Logf("\nTest:%s", c.it)
				t.FailNow()
			}
		})
	}
}
