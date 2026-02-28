package utils

import (
	"fmt"
	"reflect"
)

func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
func AssignByTag(src, dst interface{}, ignore []string, tag string) error {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()
	dstType := dstVal.Type()

	// 创建一个映射，将源结构体的json标签映射到目标结构体的字段名
	tagToFieldName := make(map[string]string)
	for i := 0; i < dstType.NumField(); i++ {
		field := dstType.Field(i)
		jsonTag := field.Tag.Get(tag)
		if jsonTag != "" {
			fieldName := field.Name
			tagToFieldName[jsonTag] = fieldName
		}
	}
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		srcFieldType := srcField.Type()
		jsonTag := srcVal.Type().Field(i).Tag.Get("json")
		if contains(ignore, jsonTag) {

			continue
		}
		fieldName, exists := tagToFieldName[jsonTag]

		if !exists {
			// 如果目标结构体没有对应的字段，则忽略并继续下一个字段
			continue
		}
		dstField := dstVal.FieldByName(fieldName)
		if !dstField.IsValid() {
			// 如果找不到对应的字段，则返回错误
			return fmt.Errorf("cannot find field %q in destination struct", fieldName)
		}
		dstFieldType := dstField.Type()

		// 确保类型兼容
		if !srcFieldType.AssignableTo(dstFieldType) && !dstFieldType.AssignableTo(srcFieldType) {
			return fmt.Errorf("incompatible types between source and destination fields: %s and %s", srcFieldType, dstFieldType)
		}

		// 如果类型不同但兼容（例如，将int赋值给interface{}），则需要先转换类型
		if srcFieldType != dstFieldType {
			srcField = srcField.Convert(dstFieldType)
		}

		// 设置目标结构体的字段值
		dstField.Set(srcField)
	}

	return nil
}
