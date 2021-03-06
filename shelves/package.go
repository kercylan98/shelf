package shelves

import (
	"fmt"
	"github.com/kercylan98/shelf"
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
