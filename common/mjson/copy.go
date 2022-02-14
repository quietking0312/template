package mjson

import (
	"reflect"
)

func CopyStatus(src, dst interface{}) error {
	sV := reflect.ValueOf(src)
	dV := reflect.ValueOf(dst).Elem()
	for i := 0; i < sV.NumField(); i++ {
		name := sV.Type().Field(i).Name
		value := sV.FieldByName(name)
		dValue := dV.FieldByName(name)
		if dValue.IsValid() == false {
			continue
		}
		if dValue.Type() == value.Type() {
			dValue.Set(value)
		}
	}
	return nil
}
