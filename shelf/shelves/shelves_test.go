package shelves

import (
	"fmt"
	"shelf/shelf"
	"testing"
)

func Test(t *testing.T) {
	var draw []shelf.Shelf

	draw = append(draw,
		NewPackage("shelf"),
		NewLine(1),
		NewImportStart().
			Add(NewImport("github.com/kercylan/shelf/test_line1")).
			Add(NewImport("github.com/kercylan/shelf/test_line2")).
			Add(NewImport("github.com/kercylan/shelf/test_line3")),
		NewImportEnd(),
		NewLine(1),
		NewFuncStart("main", []FuncParameter{}, []FuncParameter{
				{Name: "err", Type: "error"},
		}).
			Add(NewAnnotation("TODO: ...")),
		NewFuncEnd(),
		NewLine(1),
	)

	for _, s := range draw {
		fmt.Print(s.Render())
	}

}
