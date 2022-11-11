package file

func NewStructFile(s Struct) File {
	return File{
		Package: STRUCTS_PKG,
		Structs: []Struct{s},
	}
}
