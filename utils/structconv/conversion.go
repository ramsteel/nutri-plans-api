package structconv

import "reflect"

func ToMap(i any) map[string]any {
	res := make(map[string]any)
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Tag.Get("json")
		res[name] = v.Field(i).Interface()
	}

	return res
}
