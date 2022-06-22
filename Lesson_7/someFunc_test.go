package Lesson_7

import (
	"reflect"
	"testing"
)

func TestSomeFunc(t *testing.T) {
	st := struct {
		FieldString string `json:"field_string"`
		FieldInt    int
		Slice       []int
		Object      struct {
			NestedField int
		}
	}{
		FieldString: "someString",
		FieldInt:    007,
		Slice:       []int{1, 2, 3},
		Object:      struct{ NestedField int }{NestedField: 302},
	}

	m := map[string]interface{}{
		"FieldString": "asd",
	}
	err := SomeFunc(&st, m)
	if err != nil || st.FieldString != m["FieldString"] {
		t.Errorf("ChangeStruct(%v, %v) failed.", st, m)
	}

	m = map[string]interface{}{
		"FieldString": "asd",
		"Object": map[string]interface{}{
			"NestedField": 2,
		},
		"steak":    true,
		"FieldInt": 123,
		"Slice":    []int{123, 123, 123},
	}

	err = SomeFunc(&st, m)
	if err != nil ||
		st.FieldString != m["FieldString"] ||
		st.Object.NestedField != m["Object"].(map[string]interface{})["NestedField"] ||
		st.FieldInt != m["FieldInt"] ||
		!reflect.DeepEqual(st.Slice, m["Slice"]) {
		t.Errorf("ChangeStruct(%v, %v) failed.", st, m)
	}
}
