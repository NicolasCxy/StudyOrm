package reflect

import (
	"reflect"
)

func IterateFunc(entity any, methodName string, values []any) (map[string]FuncInfo, error) {
	typ := reflect.TypeOf(entity)

	numMethod := typ.NumMethod()
	res := make(map[string]FuncInfo, numMethod)

	for i := 0; i < numMethod; i++ {
		method := typ.Method(i)
		if method.Name != methodName {
			continue
		}

		fn := method.Func

		numIn := fn.Type().NumIn()
		input := make([]reflect.Type, 0, numIn)
		inputValue := make([]reflect.Value, 0, numIn)

		input = append(input, reflect.TypeOf(entity))
		inputValue = append(inputValue, reflect.ValueOf(entity))

		for j := 1; j < numIn; j++ {
			fnInType := fn.Type().In(j)
			input = append(input, fnInType)
			if len(values) < 1 {
				inputValue = append(inputValue, reflect.Zero(fnInType))
			} else {
				value := values[j-1]
				inputValue = append(inputValue, reflect.ValueOf(value))
			}
		}

		numOut := fn.Type().NumOut()
		output := make([]reflect.Type, 0, numOut)
		for j := 0; j < numOut; j++ {
			output = append(output, fn.Type().Out(j))
		}

		resValues := fn.Call(inputValue)
		result := make([]any, 0, len(resValues))
		for _, v := range resValues {
			result = append(result, v.Interface())
		}

		res[method.Name] = FuncInfo{
			Name:       method.Name,
			InputType:  input,
			OutputType: output,
			Result:     result,
		}

	}

	return res, nil

}

type FuncInfo struct {
	Name       string
	InputType  []reflect.Type
	OutputType []reflect.Type
	Result     []any
}
