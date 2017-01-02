package permutation

import (
	"reflect"
	"testing"
)

func Test_List_Indizes(t *testing.T) {
	testValues := []interface{}{"a", "b"}

	newService, err := NewService(DefaultServiceConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newList, err := NewList(DefaultListConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newList.SetMaxGrowth(3)
	newList.SetRawValues(testValues)

	// Make sure the initial index is empty.
	newIndizes := newList.Indizes()
	if len(newIndizes) != 0 {
		t.Fatal("expected", 0, "got", newIndizes)
	}

	// Make sure the initial index is obtained even after some permutations. Note
	// that we have 2 values to permutate. We are going to calculate permutations
	// of the base 2 number system. This is the binary number system. The 4th
	// permutation in the binary system is 10.
	err = newService.PermuteBy(newList, 4)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newIndizes = newList.Indizes()
	if !reflect.DeepEqual(newIndizes, []int{0, 1}) {
		t.Fatal("expected", []int{1, 1}, "got", newIndizes)
	}

	// The 12th permutation (current index already is 10) in the binary system is
	// 110.
	err = newService.PermuteBy(newList, 8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newIndizes = newList.Indizes()
	if !reflect.DeepEqual(newIndizes, []int{1, 0, 1}) {
		t.Fatal("expected", []int{1, 0, 1}, "got", newIndizes)
	}
}

func Test_List_RawValues(t *testing.T) {
	testValues := []interface{}{"a", "b"}

	newService, err := NewService(DefaultServiceConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newList, err := NewList(DefaultListConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newList.SetMaxGrowth(3)
	newList.SetRawValues(testValues)

	// Make sure the initial values are still obtained on the fresh Service.
	newValues := newList.RawValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}

	// Make sure the initial values are still obtained even after some
	// permutations.
	err = newService.PermuteBy(newList, 4)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newValues = newList.RawValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}

	err = newService.PermuteBy(newList, 8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newValues = newList.RawValues()
	if !reflect.DeepEqual(testValues, newValues) {
		t.Fatal("expected", newValues, "got", testValues)
	}
}
