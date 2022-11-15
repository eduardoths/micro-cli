package file_test

import (
	"fmt"
	"testing"

	"github.com/eduardoths/micro-cli/generator/file"
)

func TestFile_String(t *testing.T) {
	type testCase struct {
		it   string
		file file.File
		want string
	}

	tc := []testCase{
		{
			it: "should return a file that has only a package",
			file: file.File{
				Package: "teste",
			},
			want: "package teste\n",
		},
		{
			it: "should return a file with one import",
			file: file.File{
				Package: "test",
				Imports: file.Imports{
					file.Import{Path: "strings"},
				},
			},
			want: "package test\n" +
				"\n" + `import "strings"` + "\n",
		},
		{
			it: "should return a file with one import with alias",
			file: file.File{
				Package: "test",
				Imports: file.Imports{
					file.Import{Path: "strings", Name: "str"},
				},
			},
			want: "package test\n\n" +
				`import str "strings"` + "\n",
		},
		{
			it: "should return a file with two imports",
			file: file.File{
				Package: "test",
				Imports: file.Imports{
					file.Import{Path: "strings", Name: "str"},
					file.Import{Path: "errors"},
				},
			},
			want: "package test\n\n" +
				"import (\n" +
				fmt.Sprintf("\t%s\n", `"errors"`) +
				fmt.Sprintf("\tstr %s\n", `"strings"`) +
				")\n",
		},
		{
			it: "should return a file with one function",
			file: file.File{
				Package: "test",
				Imports: file.Imports{
					{Path: "fmt"},
				},
				Funcs: []file.Implementation{
					{
						Func: file.Method{
							Name: "main",
						},
						CodeLines: []string{
							"fmt.Println(\"Hello world!\")",
						},
					},
				},
			},
			want: "package test\n\n" +
				"import \"fmt\"\n\n" +
				"func main() {\n" +
				"\tfmt.Println(\"Hello world!\")\n" +
				"}\n",
		},
		{
			it: "should return a file with method implementation",
			file: file.File{
				Package: "test",
				Imports: file.Imports{
					{Path: "fmt"},
				},
				Funcs: []file.Implementation{
					{
						StructAlias: "h",
						StructName:  "*Hello",
						Func: file.Method{
							Name:    "SayHello",
							Params:  file.Args{{Name: "name", Type: "string"}},
							Results: file.Args{{Type: "string"}},
						},
						CodeLines: []string{
							"fmt.Printf(\"Hello, %s!\", name)",
						},
					},
				},
			},
			want: "package test\n\n" +
				"import \"fmt\"\n\n" +
				"func (h *Hello) SayHello(name string) string {\n" +
				"\tfmt.Printf(\"Hello, %s!\", name)\n" +
				"}\n",
		},
		{
			it: "should return a file with an empty interface",
			file: file.File{
				Package: "test",
				Interfaces: []file.Interface{
					{
						Name: "Xpto",
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {}\n",
		},
		{
			it: "should return a file with an interface containing one method without params or results",
			file: file.File{
				Package: "test",
				Interfaces: []file.Interface{
					{
						Name: "Xpto",
						Methods: []file.Method{
							{Name: "Create"},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto interface {\n" +
				"\tCreate()\n" +
				"}\n",
		},
		{
			it: "should return a file with an interface containing one method with one named param",
			file: file.File{
				Package: "test",
				Interfaces: []file.Interface{
					{
						Name: "Xpto",
						Methods: []file.Method{
							{
								Name: "Create",
								Params: file.Args{
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
				"}\n",
		},
		{
			it: "should return a file with an interface containing one method with one unnamed param",
			file: file.File{
				Package: "test",
				Interfaces: []file.Interface{
					{
						Name: "Xpto",
						Methods: []file.Method{
							{
								Name: "Create",
								Params: file.Args{
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
				"}\n",
		},
		{
			it: "should return a file with an interface containing one method with two params",
			file: file.File{
				Package: "test",
				Interfaces: []file.Interface{
					{
						Name: "Xpto",
						Methods: []file.Method{
							{
								Name: "Create",
								Params: file.Args{
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
				"}\n",
		},
		{
			it: "should return a file with an interface containing one method with one named return",
			file: file.File{
				Package: "test",
				Interfaces: []file.Interface{
					{
						Name: "Xpto",
						Methods: []file.Method{
							{
								Name: "Create",
								Results: file.Args{
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
				"}\n",
		},
		{
			it: "should return a file with an interface containing one method with one unnamed return",
			file: file.File{
				Package: "test",
				Interfaces: []file.Interface{
					{
						Name: "Xpto",
						Methods: []file.Method{
							{
								Name: "Create",
								Results: file.Args{
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
				"}\n",
		},
		{
			it: "should return a file with an interface containing one method with two unnamed returns",
			file: file.File{
				Package: "test",
				Interfaces: []file.Interface{
					{
						Name: "Xpto",
						Methods: []file.Method{
							{
								Name: "Create",
								Results: file.Args{
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
				"}\n",
		},
		{
			it: "should return a file with an interface containing one method with three named returns",
			file: file.File{
				Package: "test",
				Interfaces: []file.Interface{
					{
						Name: "Xpto",
						Methods: []file.Method{
							{
								Name: "Create",
								Results: file.Args{
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
				"}\n",
		},
		{
			it: "should return a file with an interface containing three methods",
			file: file.File{
				Package: "test",
				Interfaces: []file.Interface{
					{
						Name: "Xpto",
						Methods: []file.Method{
							{
								Name: "Create",
								Params: file.Args{
									{Type: "structs.Example", Name: "xpto1"},
									{Type: "string", Name: "name"},
									{Type: "int", Name: "i"},
								},
								Results: file.Args{
									{Type: "bool", Name: "ok"},
									{Type: "int", Name: "n"},
									{Type: "error", Name: "err"},
								},
							},
							{
								Name: "Update",
								Params: file.Args{
									{Type: "structs.Example", Name: "xpto1"},
									{Type: "string", Name: "name"},
									{Type: "int", Name: "i"},
								},
								Results: file.Args{
									{Type: "bool", Name: "ok"},
									{Type: "int", Name: "n"},
									{Type: "error", Name: "err"},
								},
							},
							{
								Name: "Delete",
								Params: file.Args{
									{Type: "structs.Example", Name: "xpto1"},
									{Type: "string", Name: "name"},
									{Type: "int", Name: "i"},
								},
								Results: file.Args{
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
				"}\n",
		},
		{
			it: "should return a file with two interfaces",
			file: file.File{
				Package: "test",
				Interfaces: []file.Interface{
					{Name: "XptoOne"},
					{Name: "XptoTwo"},
				},
			},
			want: "package test\n\n" +
				"type XptoOne interface {}\n\n" +
				"type XptoTwo interface {}\n",
		},
		{
			it: "should return a file with an empty struct",
			file: file.File{
				Package: "test",
				Structs: []file.Struct{
					{Name: "Xpto"},
				},
			},
			want: "package test\n\n" +
				"type Xpto struct {}\n",
		},
		{
			it: "should return a file with one field",
			file: file.File{
				Package: "test",
				Structs: []file.Struct{
					{
						Name: "Xpto",
						Fields: []file.Field{
							{Name: "Field"},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto struct {\n" +
				"\tField\n" +
				"}\n",
		},
		{
			it: "should return a file with one field with type",
			file: file.File{
				Package: "test",
				Structs: []file.Struct{
					{
						Name: "Xpto",
						Fields: []file.Field{
							{Name: "Field", Type: "string"},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto struct {\n" +
				"\tField string\n" +
				"}\n",
		},
		{
			it: "should return a file with one field with type and tag",
			file: file.File{
				Package: "test",
				Structs: []file.Struct{
					{
						Name: "Xpto",
						Fields: []file.Field{
							{Name: "Field", Type: "string", Tag: "`json:\"-\"`"},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto struct {\n" +
				"\tField string `json:\"-\"`\n" +
				"}\n",
		},
		{
			it: "should return a file with three fields",
			file: file.File{
				Package: "structs",
				Structs: []file.Struct{
					{
						Name: "Xpto",
						Fields: []file.Field{
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
				"}\n",
		},
		{
			it: "should return a file with a struct with one implementation",
			file: file.File{
				Package: "structs",
				Structs: []file.Struct{
					{
						Name: "Xpto",
						Implementations: []file.Implementation{
							{
								StructAlias: "x",
								StructName:  "*Xpto",
								Func: file.Method{
									Name:   "SetString",
									Params: file.Args{{Name: "s", Type: "string"}},
								},
								CodeLines: []string{
									"x.ExampleString = s",
									"panic(\"Ovo da panico\")",
								},
							},
						},
					},
				},
			},
			want: "package structs\n\n" +
				"type Xpto struct {}\n\n" +
				"func (x *Xpto) SetString(s string) {\n" +
				"\tx.ExampleString = s\n" +
				"\tpanic(\"Ovo da panico\")\n" +
				"}\n",
		},
		{
			it: "should return a file with a struct with two implementations",
			file: file.File{
				Package: "structs",
				Structs: []file.Struct{
					{
						Name: "Xpto",
						Implementations: []file.Implementation{
							{
								StructAlias: "x",
								StructName:  "Xpto",
								Func: file.Method{
									Name: "Foo",
								},
							},
							{
								StructAlias: "x",
								StructName:  "Xpto",
								Func: file.Method{
									Name: "Bar",
								},
							},
						},
					},
				},
			},
			want: "package structs\n\n" +
				"type Xpto struct {}\n\n" +
				"func (x Xpto) Foo() {\n" +
				"}\n\n" +
				"func (x Xpto) Bar() {\n" +
				"}\n",
		},
		{
			it: "should return a file with two empty structs",
			file: file.File{
				Package: "test",
				Structs: []file.Struct{
					{
						Name: "Xpto",
						Implementations: []file.Implementation{
							{
								StructAlias: "x",
								StructName:  "Xpto",
								Func: file.Method{
									Name: "Foo",
								},
							},
						},
					},
					{
						Name: "XptoAgain",
						Implementations: []file.Implementation{
							{
								StructAlias: "xa",
								StructName:  "XptoAgain",
								Func: file.Method{
									Name: "Bar",
								},
							},
						},
					},
				},
			},
			want: "package test\n\n" +
				"type Xpto struct {}\n\n" +
				"func (x Xpto) Foo() {\n" +
				"}\n\n" +
				"type XptoAgain struct {}\n\n" +
				"func (xa XptoAgain) Bar() {\n" +
				"}\n",
		},
		{
			it: "should return a complete file",
			file: file.File{
				Package: "complextest",
				Imports: []file.Import{
					{Path: "github.com/eduardoths/my_structs/structs"},
					{Name: "mystructs", Path: "github.com/eduardoths/my_structs"},
				},
				Interfaces: []file.Interface{
					{
						Name: "Xpto",
						Methods: []file.Method{
							{
								Name: "Done",
								Params: file.Args{
									{Name: "s", Type: "structs.Struct"},
								},
								Results: file.Args{{Type: "error"}},
							},
							{
								Name:    "String",
								Results: file.Args{{Type: "string"}},
							},
						},
					},
					{
						Name: "Err",
						Methods: []file.Method{
							{
								Name:    "Error",
								Results: file.Args{{Type: "string"}},
							},
						},
					},
					{Name: "empty"},
				},
				Structs: []file.Struct{
					{
						Name: "Foo",
						Fields: []file.Field{
							{
								Name: "mystructs.Foo",
							},
						},
					},
					{
						Name: "Bar",
						Fields: []file.Field{
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
				"\tmystructs \"github.com/eduardoths/my_structs\"\n" +
				"\t\"github.com/eduardoths/my_structs/structs\"\n" +
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
				"type emptyStruct struct {}\n",
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
