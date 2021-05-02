package shelves

import (
	"fmt"
	"shelf/shelf"
)

type Package struct {
	shelf.VirtualShelf
	Name string
}

func NewPackage(name string) shelf.Shelf {
	slf := &Package{Name: name}
	return slf.SetWrite(func() string {
		return fmt.Sprintf("package %s", slf.Name)
	})
}