package utils

import (
	"reflect"
	"testing"
)

func CheckResult(t *testing.T, name string, result interface{}, expected interface{}) {
	if result == expected {
		t.Logf("\t\tPASSED: %v %v, expected %v", name, result, expected)
	} else {
		t.Fatalf("\t\tFAILED: %v %v, expected %v", name, result, expected)
	}
}

func CheckNil(t *testing.T, name string, resultPointer interface{}) {
	if reflect.ValueOf(resultPointer).IsNil() {
		t.Logf("\t\tPASSED: %v %v, expected nil", name, resultPointer)
	} else {
		t.Fatalf("\t\tFAILED: %v %v, expected nil", name, resultPointer)
	}
}

func CheckNotNil(t *testing.T, name string, resultPointer interface{}) {
	if reflect.ValueOf(resultPointer).IsNil() {
		t.Fatalf("\t\tFAILED: %v is nil!", name)
	}
}
