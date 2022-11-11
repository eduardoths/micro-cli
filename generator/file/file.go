package file

import (
	"fmt"
	"strings"
)

type File struct {
	Package    string
	Imports    Imports
	Funcs      []Implementation
	Interfaces []Interface
	Structs    []Struct
}

func (f File) String() string {
	var sb strings.Builder
	sb.WriteString("package " + f.Package + "\n")
	sb.WriteString(f.Imports.String())
	for i := range f.Funcs {
		sb.WriteString(f.Funcs[i].String())
	}
	for i := range f.Interfaces {
		sb.WriteString(f.Interfaces[i].String())
	}
	for i := range f.Structs {
		sb.WriteString(f.Structs[i].String())
	}
	return sb.String()
}

type Imports []Import

func (imports Imports) String() string {
	if len(imports) == 0 {
		return ""
	}
	if len(imports) == 1 {
		return fmt.Sprintf("\nimport %s", imports[0].String())
	}
	var sb strings.Builder
	sb.WriteString("\nimport (\n")
	for j := range imports {
		sb.WriteString("\t")
		sb.WriteString(imports[j].String())
	}
	sb.WriteString(")\n")
	return sb.String()
}

type Import struct {
	Path string
	Name string
}

func (i Import) String() string {
	var sb strings.Builder
	if i.Name != "" {
		sb.WriteString(i.Name + " ")
	}
	sb.WriteString(`"`)
	sb.WriteString(i.Path)
	sb.WriteString(`"`)
	sb.WriteString("\n")
	return sb.String()
}

type Interface struct {
	Name    string
	Methods []Method
}

func (i Interface) String() string {
	var sb strings.Builder
	sb.WriteString("\ntype " + i.Name + " interface {")
	if len(i.Methods) > 0 {
		sb.WriteString("\n")
	}
	for j := range i.Methods {
		sb.WriteString("\t" + i.Methods[j].String() + "\n")
	}
	sb.WriteString("}\n")
	return sb.String()
}

type Method struct {
	Name    string
	Params  Args
	Results Args
}

func (m Method) String() string {
	var sb strings.Builder
	sb.WriteString(m.Name)
	sb.WriteString("(")
	sb.WriteString(m.Params.String())
	sb.WriteString(")")
	if len(m.Results) > 0 {
		sb.WriteString(" ")
		if m.Results[0].Name != "" || len(m.Results) > 1 {
			sb.WriteString("(")
		}
		sb.WriteString(m.Results.String())
		if m.Results[0].Name != "" || len(m.Results) > 1 {
			sb.WriteString(")")
		}
	}
	return sb.String()
}

type Args []Arg

func (a Args) String() string {
	var sb strings.Builder
	for i, arg := range a {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(arg.String())
	}
	return sb.String()
}

type Arg struct {
	Name string
	Type string
}

func (a Arg) String() string {
	if a.Name == "" {
		return a.Type
	}
	return fmt.Sprintf("%s %s", a.Name, a.Type)
}

type Struct struct {
	Name            string
	Fields          []Field
	Implementations []Implementation
}

func (s Struct) String() string {
	var sb strings.Builder
	sb.WriteString("\ntype ")
	sb.WriteString(s.Name)
	sb.WriteString(" struct {")

	if len(s.Fields) > 0 {
		sb.WriteString("\n")
	}
	for i := range s.Fields {
		sb.WriteString(s.Fields[i].String())
	}

	sb.WriteString("}\n")

	for i := range s.Implementations {
		sb.WriteString(s.Implementations[i].String())
	}
	return sb.String()
}

type Field struct {
	Name string
	Type string
	Tag  string
}

func (f Field) String() string {
	var sb strings.Builder
	sb.WriteString("\t")
	sb.WriteString(f.Name)
	if f.Type != "" {
		sb.WriteString(" ")
		sb.WriteString(f.Type)
	}
	if f.Tag != "" {
		sb.WriteString(" ")
		sb.WriteString(f.Tag)
	}
	sb.WriteString("\n")
	return sb.String()
}

type Implementation struct {
	StructAlias string
	StructName  string
	Func        Method
	CodeLines   []string
}

func (i Implementation) String() string {
	var sb strings.Builder
	sb.WriteString("\nfunc ")
	if i.StructName != "" {
		sb.WriteString("(")
		if i.StructAlias != "" {
			sb.WriteString(i.StructAlias + " ")
		}
		sb.WriteString(i.StructName)
		sb.WriteString(") ")
	}
	sb.WriteString(i.Func.String())
	sb.WriteString(" {\n")
	for j := range i.CodeLines {
		sb.WriteString("\t")
		sb.WriteString(i.CodeLines[j])
		sb.WriteString("\n")
	}
	sb.WriteString("}\n")
	return sb.String()
}
