package main

import "testing"

func TestSimpleVariadicToSlice(t *testing.T) {

	if val := simpleVariadicToSlice(); val != nil {
		t.Error("value should be nil", nil)
	} else {
		t.Log("simpleVariadicToSLice() => nil")
	}

	vals := simpleVariadicToSlice(1, 2, 3)
	expected := []int{1, 2, 3}
	isErr := false

	for i := 0; i < 3; i++ {
		if vals[i] != expected[i] {
			isErr = true
			break
		}
	}
	if isErr {
		t.Error("value should be []int {1, 2, 3}", vals)
	} else {
		t.Log("simpleVariadicToSlice(1, 2, 3) => [] int{1, 2, 3}")
	}

	vals = simpleVariadicToSlice(expected...)
	isErr = false
	for i := 0; i < 3; i++ {
		if vals[i] != expected[i] {
			isErr = true
			break
		}
	}
	if isErr {
		t.Error("value should be []int {1, 2, 3}", vals)
	} else {
		t.Log("simpleVariadicToSlice(1, 2, 3) => [] int{1, 2, 3}")
	}

}

func TestMixedVariadicToSlice(t *testing.T) {
	name, number := mixedVariadicToSlice("bob")
	if name == "bob" && number == nil {
		t.Log("Recieved as Expected: bob, <nil slice>")
	} else {
		t.Errorf("Received unexpected values: %s, %d", name, number)
	}
}
