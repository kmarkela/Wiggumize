package utils

import (
	"reflect"
)

func GetStructFieldNames(s interface{}) []string {

	result := []string{}
	// Get the type of the struct
	structType := reflect.TypeOf(s)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		result = append(result, field.Name)
	}

	return result
}
