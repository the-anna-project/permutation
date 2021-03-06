package permutation

import (
	"reflect"
	"testing"
)

// Test_Service_PermuteBy_AbsoluteDelta tests permutations by providing deltas
// always to a new Service. That we we need to provide absolute deltas.
func Test_Service_PermuteBy_AbsoluteDelta(t *testing.T) {
	testCases := []struct {
		Input        int
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        0,
			Expected:     []interface{}{},
			ErrorMatcher: nil,
		},
		{
			Input:        1,
			Expected:     []interface{}{"a"},
			ErrorMatcher: nil,
		},
		{
			Input:        2,
			Expected:     []interface{}{"b"},
			ErrorMatcher: nil,
		},
		{
			Input:        3,
			Expected:     []interface{}{"a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        4,
			Expected:     []interface{}{"a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        5,
			Expected:     []interface{}{"b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        6,
			Expected:     []interface{}{"b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        7,
			Expected:     []interface{}{"a", "a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        8,
			Expected:     []interface{}{"a", "a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        9,
			Expected:     []interface{}{"a", "b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        10,
			Expected:     []interface{}{"a", "b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        11,
			Expected:     []interface{}{"b", "a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        12,
			Expected:     []interface{}{"b", "a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        13,
			Expected:     []interface{}{"b", "b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        14,
			Expected:     []interface{}{"b", "b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        15,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
		{
			Input:        23,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
		{
			Input:        583,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
	}

	newService, err := NewService(DefaultServiceConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		// Note that we use a completely new list for each test case. That way we
		// test the absolute delta incrementation.
		newList, err := NewList(DefaultListConfig())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
		newList.SetMaxGrowth(3)
		newList.SetRawValues([]interface{}{"a", "b"})

		err = newService.PermuteBy(newList, testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		output := newList.PermutedValues()

		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

// Test_Service_PermuteBy_Increment tests if increments by 1 always work.
func Test_Service_PermuteBy_Increment(t *testing.T) {
	testCases := []struct {
		Input        int
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        1,
			Expected:     []interface{}{"a"},
			ErrorMatcher: nil,
		},
		{
			Input:        1,
			Expected:     []interface{}{"b"},
			ErrorMatcher: nil,
		},
		{
			Input:        1,
			Expected:     []interface{}{"a", "a"},
			ErrorMatcher: nil,
		},
	}

	newService, err := NewService(DefaultServiceConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	// Note that we use the same list for all test cases. That way we test the
	// incremental delta incrementation.
	newList, err := NewList(DefaultListConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newList.SetMaxGrowth(3)
	newList.SetRawValues([]interface{}{"a", "b"})

	for i, testCase := range testCases {
		err := newService.PermuteBy(newList, testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		output := newList.PermutedValues()

		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

// Test_Service_PermuteBy_RelativeDelta tests permutations by providing deltas
// always to an already existing Service. That we we need to provide relative
// deltas.
func Test_Service_PermuteBy_RelativeDelta(t *testing.T) {
	testCases := []struct {
		Input        int
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        3,
			Expected:     []interface{}{"a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        4,
			Expected:     []interface{}{"a", "a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        5,
			Expected:     []interface{}{"b", "a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        2,
			Expected:     []interface{}{"b", "b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        1,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
		{
			Input:        32,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
		{
			Input:        772,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
	}

	// Note we use the same service for all test cases.
	newService, err := NewService(DefaultServiceConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newList, err := NewList(DefaultListConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newList.SetMaxGrowth(3)
	newList.SetRawValues([]interface{}{"a", "b"})

	for i, testCase := range testCases {
		err := newService.PermuteBy(newList, testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		output := newList.PermutedValues()

		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}

// Test_Service_PermuteBy_MinGrowth tests permutations where the permuted result
// has a constant lentgh.
func Test_Service_PermuteBy_MinGrowth(t *testing.T) {
	testCases := []struct {
		Input        int
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        0,
			Expected:     []interface{}{"a", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        1,
			Expected:     []interface{}{"a", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        2,
			Expected:     []interface{}{"b", "a"},
			ErrorMatcher: nil,
		},
		{
			Input:        3,
			Expected:     []interface{}{"b", "b"},
			ErrorMatcher: nil,
		},
		{
			Input:        4,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
		{
			Input:        101,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
		{
			Input:        239,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
		{
			Input:        752,
			Expected:     nil,
			ErrorMatcher: IsMaxGrowthReached,
		},
	}

	// Note we use the same service for all test cases.
	newService, err := NewService(DefaultServiceConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newList, err := NewList(DefaultListConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newList.SetMaxGrowth(2)
	newList.SetMinGrowth(2)
	newList.SetRawValues([]interface{}{"a", "b"})

	for i, testCase := range testCases {
		err := newService.PermuteBy(newList, testCase.Input)
		if (err != nil && testCase.ErrorMatcher == nil) || (testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err)) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}

		output := newList.PermutedValues()

		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}
