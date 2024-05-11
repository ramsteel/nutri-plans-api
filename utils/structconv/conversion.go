package structconv

import (
	"fmt"
	"reflect"
)

func ToMap(i any) map[string]any {
	res := make(map[string]any)
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Tag.Get("conv")
		if name != "" {
			res[name] = v.Field(i).Interface()
			fmt.Println(name)
		}
	}

	return res
}
