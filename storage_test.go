package main

import (
	"testing"
)

func TestCheckKey(t *testing.T) {

	if checkKey(23) == nil {
		t.Error("Problem with number.")
	}
	if checkKey(23.1) == nil {
		t.Error("Problem with real number.")
	}

	if checkKey("string") != nil {
		t.Error("Problem with string.")
	}

	arrayInt := []interface{}{1, 2, 3, 4, 5}
	if checkKey(arrayInt) != nil {
		t.Error("Problem with interface array.")
	}

	arrayString := []interface{}{"1", "2", "3", "4", "5"}
	if checkKey(arrayString) != nil {
		t.Error("Problem with interface array.")
	}

	arrayInterface := []interface{}{}
	if checkKey(arrayInterface) != nil {
		t.Error("Problem with interface array.")
	}

	var mapInterface map[interface{}]interface{}
	if checkKey(mapInterface) != nil {
		t.Error("Problem with map interface -> interface.")
	}

	var mapString map[string]interface{}
	if checkKey(mapString) != nil {
		t.Error("Problem with map string -> interface.")
	}
}

func TestCheckPattern(t *testing.T) {
	type testData struct {
		search string
		key    string
		answer bool
	}

}
func TestSet(t *testing.T) {
	/*t.Error("no test")*/

}

func TestGet(t *testing.T) {

}

func TestDelete(t *testing.T) {

}

func TestKeys(t *testing.T) {

}
