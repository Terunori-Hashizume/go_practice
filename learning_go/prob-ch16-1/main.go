package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ValidateStringLength(v any) error {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return fmt.Errorf("input must be a struct")
	}
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return fmt.Errorf("input must be a struct")
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("input must be a struct")
	}

	rt := rv.Type()
	var invalid []string

	for i := 0; i < rv.NumField(); i++ {
		fieldType := rt.Field(i)
		minStr := fieldType.Tag.Get("minStrlen")
		if minStr == "" {
			continue
		}
		if fieldType.Type.Kind() != reflect.String {
			continue
		}

		min, err := strconv.Atoi(minStr)
		if err != nil {
			invalid = append(invalid, fmt.Sprintf("%s(tag minStrlen=%q is invalid)", fieldType.Name, minStr))
			continue
		}

		value := rv.Field(i).String()
		if len(value) < min {
			invalid = append(invalid, fmt.Sprintf("%s(len=%d, min=%d)", fieldType.Name, len(value), min))
		}
	}

	if len(invalid) > 0 {
		return fmt.Errorf("string length validation failed: %s", strings.Join(invalid, ", "))
	}
	return nil
}

func main() {
	fmt.Println("Hello, World!")
}
