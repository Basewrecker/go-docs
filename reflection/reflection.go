package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i = i + 1 {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}
