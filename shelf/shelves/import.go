package shelves

import (
	"fmt"
	"shelf/shelf"
)

type ImportStart struct {
	shelf.VirtualShelf
}

func NewImportStart() shelf.Shelf {
	return new(ImportStart).SetWrite(func() string {
		return "import ("
	})
}


type ImportEnd struct {
	shelf.VirtualShelf
}

func NewImportEnd() shelf.Shelf {
	return new(ImportEnd).SetWrite(func() string {
		return ")"
	})
}

type Import struct {
	shelf.VirtualShelf
	Path 		string
}

func NewImport(path string) shelf.Shelf {
	slf := new(Import)
	slf.Path = path
	return slf.SetWrite(func() string {
		return fmt.Sprintf(`"%s"`, slf.Path)
	})
}