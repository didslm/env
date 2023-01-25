package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func PopulateWithEnv(s any) error {
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		_, ok := typeField.Tag.Lookup("obj")
		if ok && valueField.Kind() == reflect.Struct {
			err := PopulateWithEnv(valueField.Addr().Interface())
			if err != nil {
				return err
			}

			continue
		}

		tag, ok := typeField.Tag.Lookup("env")
		if ok {
			err := assignValue(&valueField, os.Getenv(tag))
			if err != nil {
				return fmt.Errorf("Error assigning value, '%s'", err)
			}
		}
	}

	return nil
}

func assignValue(field *reflect.Value, value string) error {
	fieldName := field.Type().Name()

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("error parsing int value for field %s: %v", fieldName, err)
		}
		field.SetInt(intValue)
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("error parsing float value for field %s: %v", fieldName, err)
		}
		field.SetFloat(floatValue)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("error parsing bool value for field %s: %v", fieldName, err)
		}
		field.SetBool(boolValue)
	default:
		return fmt.Errorf("unsupported type for field %s: %s", fieldName, field.Kind())
	}
	return nil
}
