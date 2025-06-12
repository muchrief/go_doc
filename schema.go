package go_doc

import (
	"reflect"
	"time"

	"github.com/muchrief/go_doc/helper"
)

type DataType string

const (
	DataTypeNumber  DataType = "number"
	DataTypeInteger DataType = "integer"
	DataTypeString  DataType = "string"
	DataTypeBoolean DataType = "boolean"
	DataTypeArray   DataType = "array"
	DataTypeObject  DataType = "object"

	DataTypeUnknown DataType = "unknown"
)

func GetDataTypeMapper(dataType reflect.Kind) DataType {
	switch dataType {
	case reflect.String:
		return DataTypeString
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return DataTypeInteger

	case reflect.Float32,
		reflect.Float64:
		return DataTypeNumber

	case reflect.Array,
		reflect.Slice:
		return DataTypeArray

	case reflect.Map,
		reflect.Struct,
		reflect.Interface,
		reflect.Pointer:
		return DataTypeObject

	case reflect.Bool:
		return DataTypeBoolean
	}

	return DataTypeUnknown
}

type Schema struct {
	Type                 DataType           `yaml:"type" json:"type"`
	Properties           map[string]*Schema `yaml:"properties,omitempty" json:"properties,omitempty"`
	Items                *Schema            `yaml:"items,omitempty" json:"items,omitempty"`
	Format               string             `yaml:"format,omitempty" json:"format,omitempty"`
	AdditionalProperties *Schema            `yaml:"additionalProperties,omitempty" json:"additionalProperties,omitempty"`
	Required             []string           `yaml:"required,omitempty" json:"required,omitempty"`
}

func BuildSchema(data interface{}) *Schema {
	schema := &Schema{}
	if data == nil {
		schema.Type = DataTypeUnknown
		return schema
	}

	dataType := reflect.TypeOf(data)
	kind := dataType.Kind()

	schemaType := GetDataTypeMapper(kind)
	if schemaType == DataTypeUnknown {
		return schema
	}
	schema.Type = schemaType

	switch kind {
	case reflect.Struct:
		switch data.(type) {
		case time.Time:
			schema = &Schema{
				Type: DataTypeString,
			}

			return schema
		}

		if schema.Properties == nil {
			schema.Properties = map[string]*Schema{}
		}

		for i := 0; i < dataType.NumField(); i++ {
			field := dataType.Field(i)

			dataModel := reflect.Zero(field.Type).Interface()

			fieldName := helper.GetFieldName(field, "json")
			if field.Anonymous {
				newSchema := BuildSchema(dataModel)
				for key, value := range newSchema.Properties {
					schema.Properties[key] = value
				}

			} else {
				schemaProperties := BuildSchema(dataModel)
				schema.Properties[fieldName] = schemaProperties
			}
		}

	case reflect.Map:
		valueType := dataType.Elem()

		dataModel := reflect.Zero(valueType).Interface()

		additionalProperties := BuildSchema(dataModel)
		schema.AdditionalProperties = additionalProperties

	case reflect.Pointer:
		if schema.Properties == nil {
			schema.Properties = map[string]*Schema{}
		}

		dataModel := reflect.Zero(dataType.Elem()).Interface()
		newSchema := BuildSchema(dataModel)
		for key, value := range newSchema.Properties {
			schema.Properties[key] = value
		}

	case reflect.Array, reflect.Slice:
		valueType := dataType.Elem()

		dataModel := reflect.Zero(valueType).Interface()
		properties := BuildSchema(dataModel)
		schema.Items = properties

	case reflect.Invalid:
		return schema
	}

	return schema
}
