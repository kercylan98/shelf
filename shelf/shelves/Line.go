package shelves

import "github.com/kercylan98/shelf/shelf"

type Line struct {
	shelf.VirtualShelf
	Count 	int
}

func NewLine(count int) shelf.Shelf {
	slf := &Line{Count: count}
	return slf.SetWrite(func() string {
		if slf.Count <= 0 {
			slf.Count = 1
		}
		var result string
		for i := 0; i < slf.Count - 1; i++ {
			result += "\n"
		}
		return result
	})
}
