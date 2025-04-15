package axo

import (
	"fmt"
	"reflect"
)

func StructToMap(obj any) map[string]any {
	result := make(map[string]any)

	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	fmt.Println(typ)

	for i := range val.NumField() {
		fieldName := typ.Field(i).Name
		fieldValueKind := val.Field(i).Kind()
		var fieldValue any

		if fieldValueKind == reflect.Struct {
			fieldValue = StructToMap(val.Field(i).Interface())
		} else {
			fieldValue = val.Field(i).Interface()
		}

		result[fieldName] = fieldValue
	}

	return result
}
