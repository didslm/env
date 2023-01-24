package env

import (
	"os"
	"reflect"
	"strconv"
)

func PopulateWithEnv(s interface{}) {
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		obj := typeField.Tag.Get("obj")
		if obj != "" && valueField.Kind() == reflect.Struct {
			PopulateWithEnv(valueField.Addr().Interface())
			continue
		}

		tag := typeField.Tag.Get("env")
		if tag != "" {
			envValue := os.Getenv(tag)
			if envValue != "" {
				assignValue(&valueField, envValue)
			}
		}
	}
}

func assignValue(field *reflect.Value, value string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, _ := strconv.ParseInt(value, 10, 64)
		field.SetInt(intValue)
	case reflect.Float32, reflect.Float64:
		floatValue, _ := strconv.ParseFloat(value, 64)
		field.SetFloat(floatValue)
	case reflect.Bool:
		boolValue, _ := strconv.ParseBool(value)
		field.SetBool(boolValue)
	}
}
