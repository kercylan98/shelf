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
	fmt.Println("before group1: ", group1.Render())
	fmt.Println("before group2: ", group2.Render())

	group1.Bind(item2, item3)
	group2.Bind(item1)
	fmt.Println("after  group1: ", group1.Render())
	fmt.Println("after  group2: ", group2.Render())
}

// 测试组删除
func TestGroup_Del(t *testing.T) {
	group := new(shelf.Group)

	item1 := shelves.NewPackage("pkg1")
	item2 := shelves.NewPackage("pkg2")
	item3 := shelves.NewPackage("pkg3")

	group.Bind(item1, item2, item3)
	fmt.Println("before ===========> ", group.Render())
	group.Del(item2)
	fmt.Println("after  ===========> ", group.Render())
}

// 测试组删除后添加
func TestGroup_Del_After_Add(t *testing.T) {
	group := new(shelf.Group)

	item1 := shelves.NewPackage("pkg1")
	item2 := shelves.NewPackage("pkg2")
	item3 := shelves.NewPackage("pkg3")

	group.Bind(item1, item2, item3)
	group.Del(item2)
	fmt.Println("before ===========> ", group.Render())
	group.Bind(item2)
	fmt.Println("after  ===========> ", group.Render())
}
