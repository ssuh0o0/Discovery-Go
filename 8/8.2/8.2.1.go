// 8.2.1 유닛테스트
package main

import (
	"reflect"
	"testing"
)

// 자료형을 한정시킨 assertEqual
func assertEqualString(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("%s != %s", expected, actual)
	}
}

// 범용 적인 assertEqual
func assertEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}
}

// 테이블 기반 테스트를 만들수 있음
func Test(t *testing.T) {
	examples := []struct {
		desc, expected, input string
	}{{
		desc:     "empty case",
		expected: "",
		input:    "",
	}}
	for _, ex := range examples {
		actual := ex.input
		if ex.expected != actual {
			t.Errorf("%s: %s != %s", ex.desc, ex.expected, ex.input)
		}
	}
}
