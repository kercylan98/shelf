package shelf

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type Shelf interface {
	// 获取ID
	GetID() string
	// 获取虚拟架子
	GetVirtual() *VirtualShelf
	// 获取壳子
	GetParent() Shelf
	// 获取所在的组
	GetGroup() *Group
	// 获取所有子架子
	GetAllChildren() []Shelf
	// 是否是根架子
	IsRoot() bool
	// 设置输出内容
	SetWrite(func() string) Shelf
	// 检查子架子是否存在某架子并返回其壳子
	Contains(shelf Shelf) (Shelf, bool)

	// 绑定某架子到该架子到末尾，并返回支持继续添加到相同架子的架子
	Bind(shelf ...Shelf) Shelf
	// 绑定架子到该架子末尾，并返回子架子
	BindC(shelf Shelf) Shelf
	// 删除特定架子
	Del(shelf Shelf)
	// 移动子架子位置
	Move(shelf Shelf, index int) error

	// 设置壳子
	SetParent(shelf Shelf)
	// 设置所在的组
	SetGroup(group *Group)

	// 渲染该架子
	Render(level int) string
}

type VirtualShelf struct {
	_init     sync.Once      // 初始化锁
	_id       string         // 架子的id
	_parent   Shelf          // 架子的壳子
	_children []Shelf        // 架子里的架子
	_mapper   map[string]int // 架子里的架子的映射
	_write    func() string  // 输出函数
	_group    *Group         // 架子所在的组
}

func (slf *VirtualShelf) Move(shelf Shelf, index int) error {
	// 检查子成员是否存在
	if i, exist := slf._mapper[shelf.GetID()]; exist {
		slf._children = append(slf._children[:i], slf._children[i+1:]...)
	} else {
		return errors.New("the shelf does not contain the child shelf")
	}

	// 索引范围控制
	if index < 0 {
		index = 0
	}

	if index > len(slf._children)-1 {
		index = len(slf._children) - 1
	}

	// 移动
	switch index {
	case 0:
		slf._children = append([]Shelf{shelf}, slf._children...)
	case 1:
		slf._children = append(append([]Shelf{slf._children[0]}, shelf), slf._children[1:]...)
	default:
		left := slf._children[:index]
		right := slf._children[index:]
		slf._children = append(append(left, shelf), right...)
	}

	// 重建索引
	for i, s := range slf._children {
		slf._mapper[s.GetID()] = i
	}
	return nil
}

func (slf *VirtualShelf) Contains(shelf Shelf) (Shelf, bool) {
	if _, exist := slf._mapper[shelf.GetID()]; exist {
		return slf, true
	}
	for _, child := range slf._children {
		if parent, exist := child.Contains(shelf); exist {
			return parent, true
		}
	}
	return nil, false
}

func (slf *VirtualShelf) SetGroup(group *Group) {
	if group == nil {
		if slf._group != nil {
			slf._group.Del(slf)
		}
		slf._group = group
		return
	}

	if _, exist := group.mapper[slf._id]; exist {
		slf._group = group
		return
	}

	if slf._parent != nil {
		slf._parent.Del(slf)
	}
	if slf._group != nil {
		slf._group.Del(slf)
	}

	group.Bind(slf)
}

func (slf *VirtualShelf) GetGroup() *Group {
	return slf._group
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
	slf.init()
	slf._write = f
	return slf
}

func (slf *VirtualShelf) GetVirtual() *VirtualShelf {
	return slf
}

func (slf *VirtualShelf) SetParent(parent Shelf) {
	if slf.GetParent() != nil {
		slf.GetParent().Del(slf)
	}

	if parent != nil {
		parentVs := parent.GetVirtual()
		parentVs._mapper[slf.GetID()] = len(parentVs._children)
		parentVs._children = append(parentVs._children, slf)
	}

	slf._parent = parent
}

func (slf *VirtualShelf) Bind(shelf ...Shelf) Shelf {
	for _, s := range shelf {
		// 跳过自己与自己绑定
		if s.GetID() == slf._id {
			continue
		}

		if slf._parent != nil {
			if slf._parent.GetID() == s.GetID() {
				if slf._parent.GetParent() != nil {
					slf._parent.GetParent().Bind(slf)
				} else {
					if slf._parent.GetGroup() != nil {
						slf._parent.GetGroup().Bind(slf)
					} else {
						slf._parent.Del(slf)
						slf.Bind(s)
					}
				}
				slf.Bind(s)
				continue
			}
			if p, exist := s.Contains(slf); exist {
				// 当壳子的壳子已经是最外层架子的时候，应该将两个架子当作新建的架子来重新挂载，否则升级架子后再成为壳子
				p.Del(slf)
				slf.Bind(s)
				continue
			}
		}

		if s.GetParent() != nil {
			s.GetParent().Del(s)
		}
		s.SetParent(slf)
	}
	return slf
}

func (slf *VirtualShelf) BindC(shelf Shelf) Shelf {
	slf.Bind(shelf)
	return shelf
}

func (slf *VirtualShelf) Del(shelf Shelf) {
	if index, exist := slf._mapper[shelf.GetID()]; exist {
		delete(slf._mapper, shelf.GetID())
		if len(slf._children) == 1 {
			slf._children = []Shelf{}
		} else {
			slf._children = append(slf._children[:index], slf._children[index+1:]...)
			for i, s := range slf._children {
				slf._mapper[s.GetID()] = i
			}
		}
		shelf.SetParent(nil)
	}
}

func (slf *VirtualShelf) Render(level int) string {
	var result string
	for i := 0; i < level; i++ {
		result = result + "    "
	}
	result = result + slf._write() + "\n"
	for _, child := range slf._children {
		result += child.Render(level + 1)
	}
	return result
}

func (slf *VirtualShelf) GetID() string {
	return slf._id
}

func (slf *VirtualShelf) IsRoot() bool {
	return slf._parent == nil
}

func (slf *VirtualShelf) GetParent() Shelf {
	return slf._parent
}

func (slf *VirtualShelf) GetAllChildren() []Shelf {
	return slf._children
}
