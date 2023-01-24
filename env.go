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
		tag := typeField.Tag.Get("env")
		if tag != "" {
			envValue := os.Getenv(tag)
			if envValue != "" {
				switch valueField.Kind() {
				case reflect.String:
					valueField.SetString(envValue)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					intValue, _ := strconv.ParseInt(envValue, 10, 64)
					valueField.SetInt(intValue)
				case reflect.Float32, reflect.Float64:
					floatValue, _ := strconv.ParseFloat(envValue, 64)
					valueField.SetFloat(floatValue)
				case reflect.Bool:
					boolValue, _ := strconv.ParseBool(envValue)
					valueField.SetBool(boolValue)
				}
			}
		}
	}
}
