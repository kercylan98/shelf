# Shelf

> 可用于构建灵活代码、文本生成的工具

### 案例

```
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
```

```
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

```

### 输出

```
package shelf

import (
    "github.com/kercylan/shelf/test_line1"
    "github.com/kercylan/shelf/test_line2"
    "github.com/kercylan/shelf/test_line3"
)

func main() (err error) {
    // TODO: ...
}

```