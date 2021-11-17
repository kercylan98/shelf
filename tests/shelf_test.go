package tests

import (
	"fmt"
	"github.com/kercylan98/shelf/shelves"
	"testing"
)

// 测试架子绑定
func TestShelf_Bind(t *testing.T) {
	shelf1 := shelves.NewPackage("shelf1")
	shelf2 := shelves.NewPackage("shelf2")
	shelf3 := shelves.NewPackage("shelf3")

	shelf1.Bind(shelf2, shelf3)

	fmt.Println(shelf1.Render(0))
	fmt.Println(shelf3.Render(0))
}

// 测试架子绑定后转移到自己
func TestShelf_Bind_Transfer_Self(t *testing.T) {
	shelf1 := shelves.NewPackage("shelf1")
	shelf2 := shelves.NewPackage("shelf2")
	shelf3 := shelves.NewPackage("shelf3")

	shelf1.Bind(shelf2, shelf3)
	fmt.Println(shelf1.Render(0))

	shelf1.Bind(shelf2)
	fmt.Println(shelf1.Render(0))
}

// 测试架子绑定后转移到其他架子
func TestShelf_Bind_Transfer_Other(t *testing.T) {
	shelf1 := shelves.NewPackage("shelf1")
	shelf2 := shelves.NewPackage("shelf2")
	shelf3 := shelves.NewPackage("shelf3")

	shelf1.Bind(shelf2, shelf3)
	fmt.Println(shelf1.Render(0))

	// 如果2是3的父节点的情况下，将2挂载为3的子节点，那么两者应该交换
	shelf2.Bind(shelf3)
	fmt.Println(shelf1.Render(0))
}

// 测试架子绑定到自己最近到壳子上
func TestShelf_Bind_Self_Box(t *testing.T) {
	shelf1 := shelves.NewPackage("shelf1")
	shelf2 := shelves.NewPackage("shelf2")
	shelf3 := shelves.NewPackage("shelf3")
	shelf4 := shelves.NewPackage("shelf4")

	shelf1.Bind(shelf2.Bind(shelf3), shelf4)
	fmt.Println(shelf1.Render(0))

	shelf3.Bind(shelf2)
	fmt.Println(shelf1.Render(0))
}

// 测试架子绑定到自己更上层壳子上
func TestShelf_Bind_Self_Box_Box(t *testing.T) {
	shelf1 := shelves.NewPackage("shelf1")
	shelf2 := shelves.NewPackage("shelf2")
	shelf3 := shelves.NewPackage("shelf3")
	shelf4 := shelves.NewPackage("shelf4")

	shelf1.Bind(shelf2.Bind(shelf3.Bind(shelf4)))
	fmt.Println(shelf1.Render(0))

	shelf3.Bind(shelf1)
	fmt.Println(shelf3.Render(0))
}

// 测试架子绑定到根壳子上
func TestShelf_Bind_Re_Bind_Self(t *testing.T) {
	shelf1 := shelves.NewPackage("shelf1")
	shelf2 := shelves.NewPackage("shelf2")
	shelf3 := shelves.NewPackage("shelf3")
	shelf4 := shelves.NewPackage("shelf4")

	shelf1.Bind(shelf2, shelf3, shelf4)
	fmt.Println(shelf1.Render(0))

	shelf2.Bind(shelf1)
	fmt.Println(shelf2.Render(0))
}

// 测试只有一个架子时移动架子
func TestShelf_Move_One(t *testing.T) {
	root := shelves.NewPackage("pkg")
	item1 := shelves.NewPackage("child1")

	root.Bind(item1)

	fmt.Println(root.Render(0))
	root.Move(item1, 0)
	fmt.Println(root.Render(0))
}

// 测试有两个架子时移动架子
func TestShelf_Move_Two(t *testing.T) {
	root := shelves.NewPackage("pkg")
	item1 := shelves.NewPackage("child1")
	item2 := shelves.NewPackage("child2")

	root.Bind(item1, item2)

	fmt.Println(root.Render(0))
	root.Move(item1, 1)
	fmt.Println(root.Render(0))
}

// 测试有三个架子时移动架子
func TestShelf_Move_Three(t *testing.T) {
	root := shelves.NewPackage("pkg")
	item1 := shelves.NewPackage("child1")
	item2 := shelves.NewPackage("child2")
	item3 := shelves.NewPackage("child3")

	root.Bind(item1, item2, item3)

	fmt.Println(root.Render(0))
	root.Move(item1, 0)
	fmt.Println(root.Render(0))
	root.Move(item1, 1)
	fmt.Println(root.Render(0))
	root.Move(item1, 2)
	fmt.Println(root.Render(0))
	root.Move(item1, 0)
	fmt.Println(root.Render(0))

}
