package shelf

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

type Group struct {
	_init   sync.Once      // 初始化
	id      string         // 组id
	shelves []Shelf        // 组里面的架子
	mapper  map[string]int // 组里面的架子映射
}

// 初始化组
func (slf *Group) init() {
	slf._init.Do(func() {
		slf.id = uuid.NewV4().String()
		slf.shelves = []Shelf{}
		slf.mapper = map[string]int{}
	})
}

// 设置该组内特定架子的位置
func (slf *Group) Move(shelf Shelf, index int) {
	if len(slf.shelves) <= 1 {
		return
	}
	if index < 0 {
		index = 0
	}
	if index > len(slf.shelves)-1 {
		index = len(slf.shelves) - 1
	}
	if i, exist := slf.mapper[shelf.GetID()]; exist {
		slf.shelves = append(slf.shelves[:i], slf.shelves[i+1:]...)
	} else {
		return
	}

	switch index {
	case 0:
		slf.shelves = append([]Shelf{shelf}, slf.shelves...)
	case 1:
		slf.shelves = append(append([]Shelf{slf.shelves[0]}, shelf), slf.shelves[1:]...)
	default:
		left := slf.shelves[:index]
		right := slf.shelves[index:]
		slf.shelves = append(append(left, shelf), right...)
	}

	for i, s := range slf.shelves {
		slf.mapper[s.GetID()] = i
	}
}

// 从该组内删除某架子
func (slf *Group) Del(shelf Shelf) {
	slf.init()
	if index, exist := slf.mapper[shelf.GetID()]; exist {
		delete(slf.mapper, shelf.GetID())
		if len(slf.shelves) == 1 {
			slf.shelves = []Shelf{}
		} else {
			slf.shelves = append(slf.shelves[:index], slf.shelves[index+1:]...)
			for i, s := range slf.shelves {
				slf.mapper[s.GetID()] = i
			}
		}
		if shelf.GetGroup() != nil {
			shelf.SetGroup(nil)
		}
	}
}

// 绑定架子到该组，如果架子已存在组则会将对应绑定转移到该组
func (slf *Group) Bind(shelf ...Shelf) *Group {
	slf.init()
	for _, s := range shelf {
		if s.GetGroup() != nil && s.GetGroup().id == slf.id {
			if index, exist := slf.mapper[s.GetID()]; exist {
				delete(slf.mapper, s.GetID())
				if len(slf.shelves) == 1 {
					slf.shelves = []Shelf{}
				} else {
					for _, s := range slf.shelves[index+1:] {
						slf.mapper[s.GetID()] = slf.mapper[s.GetID()] - 1
					}
					slf.shelves = append(slf.shelves[:index], slf.shelves[index+1:]...)
				}
				slf.mapper[s.GetID()] = len(slf.shelves)
				slf.shelves = append(slf.shelves, s)
			}
			continue
		}
		if s.GetParent() == nil {
			if s.GetGroup() != nil {
				s.GetGroup().Del(s)
			}
			slf.mapper[s.GetID()] = len(slf.shelves)
			slf.shelves = append(slf.shelves, s)
			s.SetGroup(slf)
			continue
		}

		if s.GetParent() != nil {
			s.GetParent().Del(s)
			s.SetParent(nil)
		} else {
			if s.GetGroup() != nil {
				s.GetGroup().Del(s)
			}
		}
		s.SetGroup(slf)
	}
	return slf
}

// 渲染该组
func (slf *Group) Render() string {
	slf.init()
	var result string
	for _, shelf := range slf.shelves {
		result += shelf.Render(0)
	}
	return result
}
