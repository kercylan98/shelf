package shelves

import (
	"fmt"
	"github.com/kercylan98/shelf"
)

type FuncParameter struct {
	Name string
	Type string
}

type FuncStart struct {
	shelf.VirtualShelf
	Name    string
	Inputs  []FuncParameter
	Outputs []FuncParameter
}

func NewFuncStart(name string, inputs []FuncParameter, outputs []FuncParameter) shelf.Shelf {
	slf := &FuncStart{
		Name:    name,
		Inputs:  inputs,
		Outputs: outputs,
	}
	return slf.SetWrite(func() string {
		var input, output string
		if slf.Inputs != nil && len(slf.Inputs) > 0 {
			for _, parameter := range slf.Inputs {
				input = fmt.Sprintf("%s%s %s, ", input, parameter.Name, parameter.Type)
			}
			input = input[:len(input)-2]
		}
		if slf.Outputs != nil && len(slf.Outputs) > 0 {
			for _, parameter := range slf.Outputs {
				output = fmt.Sprintf("%s%s %s, ", output, parameter.Name, parameter.Type)
			}
			output = output[:len(output)-2]
		}
		if slf.Outputs != nil && len(slf.Outputs) > 0 {
			return fmt.Sprintf("func %s(%s) (%s) {", slf.Name, input, output)
		}
		return fmt.Sprintf("func %s(%s) {", slf.Name, input)
	})
}

type FuncEnd struct {
	shelf.VirtualShelf
}

func NewFuncEnd() shelf.Shelf {
	return new(FuncEnd).SetWrite(func() string {
		return "}"
	})
}
