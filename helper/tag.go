package helper

import (
	"reflect"
	"strings"
)

func GetFieldName(field reflect.StructField, tag string) string {
	value := field.Tag.Get(tag)
	if value == "" {
		value = field.Name
	} else {
		values := strings.Split(value, ",")
		value = values[0]
	}

	return value
}
