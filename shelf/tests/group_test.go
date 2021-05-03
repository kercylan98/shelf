package tests

import (
	"fmt"
	"shelf/shelf"
	"shelf/shelf/shelves"
	"testing"
)

// 测试组绑定架子
func TestGroup_Bind(t *testing.T) {
	group := new(shelf.Group)

	item1 := shelves.NewPackage("pkg1")
	item2 := shelves.NewPackage("pkg2")
	item3 := shelves.NewPackage("pkg3")

	group.Bind(item1, item2, item3)

	fmt.Println(group.Render())
}

// 测试组内只有一个架子时移动架子
func TestGroup_Move_One(t *testing.T) {
	group := new(shelf.Group)

	item1 := shelves.NewPackage("pkg1")

	group.Bind(item1)

	fmt.Println(group.Render())
	group.Move(item1, 2)
	fmt.Println(group.Render())

}

// 测试组内有两个架子时移动架子
func TestGroup_Move_Two(t *testing.T) {
	group := new(shelf.Group)

	item1 := shelves.NewPackage("pkg1")
	item2 := shelves.NewPackage("pkg2")

	group.Bind(item1, item2)

	fmt.Println(group.Render())
	group.Move(item1, 1)
	fmt.Println(group.Render())

}

// 测试组内有三个架子时移动架子
func TestGroup_Move_Three(t *testing.T) {
	group := new(shelf.Group)

	item1 := shelves.NewPackage("pkg1")
	item2 := shelves.NewPackage("pkg2")
	item3 := shelves.NewPackage("pkg3")

	group.Bind(item1, item2, item3)

	fmt.Println(group.Render())
	group.Move(item1, 0)
	fmt.Println(group.Render())
	group.Move(item1, 1)
	fmt.Println(group.Render())
	group.Move(item1, 2)
	fmt.Println(group.Render())
}

// 测试组绑定架子后向自己转移
func TestGroup_Bind_Transfer_Self(t *testing.T) {
	group := new(shelf.Group)

	item1 := shelves.NewPackage("pkg1")
	item2 := shelves.NewPackage("pkg2")
	item3 := shelves.NewPackage("pkg3")

	group.Bind(item1, item2, item3)
	fmt.Println("before ===========> ", group.Render())
	group.Bind(item2)
	fmt.Println("after  ===========> ", group.Render())
}

// 测试组绑定架子后向其他组转移
func TestGroup_Bind_Transfer_Other(t *testing.T) {
	group1 := new(shelf.Group)
	group2 := new(shelf.Group)

	item1 := shelves.NewPackage("pkg1")
	item2 := shelves.NewPackage("pkg2")
	item3 := shelves.NewPackage("pkg3")

	group1.Bind(item1, item3)
	group2.Bind(item2)
	fmt.Println(group1.Render())
	fmt.Println(group2.Render())

	group1.Bind(item2, item3)
	group2.Bind(item1)
	fmt.Println(group1.Render())
	fmt.Println(group2.Render())
}

// 测试组删除
func TestGroup_Del(t *testing.T) {
	group := new(shelf.Group)

	item1 := shelves.NewPackage("pkg1")
	item2 := shelves.NewPackage("pkg2")
	item3 := shelves.NewPackage("pkg3")

	group.Bind(item1, item2, item3)
	fmt.Println(group.Render())

	group.Del(item2)
	fmt.Println(group.Render())
}

// 测试组删除后添加
func TestGroup_Del_After_Add(t *testing.T) {
	group := new(shelf.Group)

	item1 := shelves.NewPackage("pkg1")
	item2 := shelves.NewPackage("pkg2")
	item3 := shelves.NewPackage("pkg3")

	group.Bind(item1, item2, item3)
	group.Del(item2)
	fmt.Println(group.Render())
	group.Bind(item2)
	fmt.Println(group.Render())
}
