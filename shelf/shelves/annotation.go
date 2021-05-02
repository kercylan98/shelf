package shelves

import (
	"fmt"
	"shelf/shelf"
)

type Annotation struct {
	shelf.VirtualShelf
	Content string
}

func NewAnnotation(content string) shelf.Shelf {
	slf := new(Annotation)
	slf.Content = content
	return slf.SetWrite(func() string {
		return fmt.Sprintf("// %s", slf.Content)
	})
}