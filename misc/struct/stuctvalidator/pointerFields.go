package structvalidator

import (
	"reflect"
)

//CheckPointerFieldsExist return false if any field is nil ad slice of fileds that are nil
func CheckPointerFieldsExist(structPtr interface{}, requiredFields []string) (isExist bool, nilSlice []string) {
	s := reflect.ValueOf(structPtr).Elem()
	typeOfT := s.Type()
	isExist = true

	for _, rf := range requiredFields {

		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			if typeOfT.Field(i).Name == rf && reflect.ValueOf(f.Interface()).IsNil() {
				nilSlice = append(nilSlice, typeOfT.Field(i).Name)
				isExist = false
			}
		}
	}
	return
}
