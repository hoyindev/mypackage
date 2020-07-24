package structvalidator

import (
	structvalidator "myUtils/structValidator"
	"reflect"
	"testing"
)

func TestCheckPointerFieldsExist(t *testing.T) {
	type S struct {
		Name  *string
		ID    *string
		No    *int
		Price *int
	}

	var (
		price = -1
		ptr1  = &price
		str   = ""
		ptr2  = &str
		s     = S{
			Price: ptr1,
			Name:  ptr2,
			// ID:    ptr2,
			No: ptr1,
		}
	)

	isExist, nilSlice := structvalidator.CheckPointerFieldsExist(&s, []string{"Name", "Price", "ID", "No"})

	wantIsExist := false
	wantNilSlice := []string{"ID"}
	if !reflect.DeepEqual(isExist, wantIsExist) {
		t.Errorf("CheckPointerFieldsExist() gotNilSlice = %v, want %v", isExist, wantIsExist)
	}
	if !reflect.DeepEqual(nilSlice, wantNilSlice) {
		t.Errorf("CheckPointerFieldsExist() gotNilSlice = %v, want %v", nilSlice, wantNilSlice)
	}

}
