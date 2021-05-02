package shelf

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

type Shelf interface {
	// 获取ID
	GetID() string
	// 获取虚拟架子
	GetVirtual() *VirtualShelf
	// 是否是根节点
	IsRoot() bool
	// 获取父节点
	GetParent() Shelf
	// 获取所有子节点
	GetAllChildren() []Shelf
	// 移动该节点到特定节点下
	Move(parent Shelf)
	// 添加节点到该节点末尾，并返回支持继续添加到相同父节点的节点
	Add(shelf Shelf) Shelf
	// 添加节点到该节点末尾，并返回子节点
	AddC(shelf Shelf) Shelf
	// 删除特定节点
	Del(shelf Shelf)
	// 设置架子层级
	SetLevel(level int)
	// 设置壳子
	SetParent(shelf Shelf)
	// 渲染该节点
	Render() string
	// 设置输出内容
	SetWrite(func() string) Shelf
}

type VirtualShelf struct {
	_init	  sync.Once		 // 初始化锁
	_id       string         // 架子的id
	_parent   Shelf          // 架子的壳子
	_children []Shelf        // 架子里的架子
	_mapper   map[string]int // 架子里的架子的映射
	_level    int            // 架子的层级
	_write 	  func() string  // 输出函数
}


func (slf *VirtualShelf) init() {
	slf._init.Do(func() {
		slf._id = uuid.NewV4().String()
		slf._children = []Shelf{}
		slf._mapper = map[string]int{}
		if slf._write == nil {
			slf._write = func() string {
				panic("please use VirtualShelf.SetWrite()")
			}
		}
	})
}

func (slf *VirtualShelf) SetWrite(f func() string) Shelf {
	slf._write = f
	return slf
}

func (slf *VirtualShelf) GetVirtual() *VirtualShelf {
	slf.init()
	return slf
}

func (slf *VirtualShelf) SetParent(parent Shelf) {
	slf.init()
	if slf.GetParent() != nil {
		slf.GetParent().Del(slf)
	}
	parentVs := parent.GetVirtual()
	parentVs._mapper[slf.GetID()] = len(parentVs._children)
	parentVs._children = append(parentVs._children, slf)
	slf._parent = parent
	slf._level = parentVs._level + 1
}

func (slf *VirtualShelf) Move(parent Shelf) {
	slf.init()
	if slf._parent != nil {
		slf._parent.Del(slf)
	}
	slf.SetParent(parent)
}

func (slf *VirtualShelf) Add(shelf Shelf) Shelf {
	slf.init()
	shelf.Move(slf)
	return slf
}

func (slf *VirtualShelf) AddC(shelf Shelf) Shelf {
	slf.init()
	slf.Add(shelf)
	return shelf
}

func (slf *VirtualShelf) Del(shelf Shelf) {
	slf.init()
	if index, exist := slf._mapper[shelf.GetID()]; exist {
		delete(slf._mapper, shelf.GetID())
		slf._children = append(slf._children[:index], slf._children[index + 1:]...)
	}
}

func (slf *VirtualShelf) Render() string {
	slf.init()
	var result string
	for i := 0; i < slf._level; i++ {
		result = result + "    "
	}
	result = result + slf._write() + "\n"
	for _, child := range slf._children {
		result += child.Render()
	}
	return result
}


func (slf *VirtualShelf) SetLevel(level int) {
	slf.init()
	slf._level = level
}

func (slf *VirtualShelf) GetID() string {
	slf.init()
	return slf._id
}

func (slf *VirtualShelf) IsRoot() bool {
	slf.init()
	return slf._parent == nil
}

func (slf *VirtualShelf) GetParent() Shelf {
	slf.init()
	return slf._parent
}

func (slf *VirtualShelf) GetAllChildren() []Shelf {
	slf.init()
	return slf._children
}
