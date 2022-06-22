package Lesson_7

import (
	"errors"
	"reflect"
)

func SomeFunc(in interface{}, values map[string]interface{}) error {
	if in == nil || values == nil {
		return nil
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return errors.New("'Input' value is not a struct")
	}
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		valuesMap := values[typeField.Name]
		if valuesMap == nil {
			continue
		}
		if typeField.Type.Kind() == reflect.Struct {
			v, ok := valuesMap.(map[string]interface{})
			_ = v
			if !ok {
				continue
			}
			err := SomeFunc(val.Field(i).Addr().Interface(), v)
			if err != nil {
				return err
			}
			continue
		}
		if reflect.TypeOf(valuesMap).Kind() != typeField.Type.Kind() {
			return errors.New("value of type " + reflect.TypeOf(valuesMap).Kind().String() + " is not assignable to type " + typeField.Type.Kind().String())
		}
		val.Field(i).Set(reflect.ValueOf(valuesMap))
	}
	return nil
}
