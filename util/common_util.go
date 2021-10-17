package util

import (
	"reflect"
)

func HasLength(i interface{}) (result bool) {

	iType := reflect.TypeOf(i).String()
	result = false

	switch iType {
	case "int":

	case "string":
		iString := i.(string)
		if i == nil || iString == "" {
			return
		}
		result = true
		return result
	default:
		return
	}
	return
}
